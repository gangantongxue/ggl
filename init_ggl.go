// Package ggl Copyright (c) 2025 gangantongxue. All rights reserved.
// Licensed under the MIT License.
package ggl

import (
	"os"
	"path/filepath"
	"sync"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// Config 配置
type Config struct {
	LogFileName    string        // 日志文件名
	LogFileDir     string        // 日志文件存放目录
	LogMaxSize     int           // 日志文件最大大小，单位MB
	LogMaxBackups  int           // 日志文件最大备份数量
	LogMaxAge      int           // 日志文件最大保留时间，单位天
	LogCompress    bool          // 是否压缩日志文件
	ToConsole      bool          // 是否同时输出到控制台
	RotateInterval time.Duration // 日志轮转时间间隔，默认1天
}

// DefaultConfig DefeatConfig 默认配置
func DefaultConfig() *Config {
	return &Config{
		LogFileName:    "ggl_log_2006-01-02.log",
		LogFileDir:     "./log",
		LogMaxSize:     100,
		LogMaxBackups:  10,
		LogMaxAge:      7,
		LogCompress:    true,
		ToConsole:      false,
		RotateInterval: time.Hour * 24, // 默认1天轮转一次
	}
}

// DailyLogger 日志管理器
type DailyLogger struct {
	mu     sync.Mutex
	logger *zap.Logger
	quit   chan struct{}
}

// New 创建ggl实例
func New(cfg *Config) *DailyLogger {
	d := &DailyLogger{
		quit: make(chan struct{}),
	}
	d.rotateLogger(cfg)
	go d.scheduleRotate(cfg)
	return d
}

// getLogFileName 获取日志文件名
func getLogFileName(logFileDir string, logFileName string) string {
	return filepath.Join(logFileDir, time.Now().Format(logFileName))
}

// newZapLogger 初始化zap日志
func newZapLogger(cfg *Config) *zap.Logger {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   getLogFileName(cfg.LogFileDir, cfg.LogFileName),
		MaxSize:    cfg.LogMaxSize,
		MaxBackups: cfg.LogMaxBackups,
		MaxAge:     cfg.LogMaxAge,
		Compress:   cfg.LogCompress,
	}

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	encoder := zapcore.NewConsoleEncoder(encoderConfig)

	fileCore := zapcore.NewCore(
		encoder,
		zapcore.AddSync(lumberJackLogger),
		zap.NewAtomicLevelAt(zapcore.InfoLevel),
	)

	var core zapcore.Core
	if cfg.ToConsole {
		consoleCore := zapcore.NewCore(
			encoder,
			zapcore.AddSync(os.Stdout),
			zap.NewAtomicLevelAt(zapcore.InfoLevel),
		)
		core = zapcore.NewTee(fileCore, consoleCore)
	} else {
		core = fileCore
	}

	return zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.WarnLevel))
}

// rotateLogger 切换到新的日志文件
func (d *DailyLogger) rotateLogger(cfg *Config) {
	d.mu.Lock()
	defer d.mu.Unlock()

	if d.logger != nil {
		_ = d.logger.Sync()
	}

	d.logger = newZapLogger(cfg)
	zap.ReplaceGlobals(d.logger) // 替换全局 logger
}

// scheduleRotate 每天 0 点切换一次日志
func (d *DailyLogger) scheduleRotate(cfg *Config) {
	for {
		now := time.Now()
		next := now.Add(cfg.RotateInterval).Truncate(cfg.RotateInterval)
		timer := time.NewTimer(next.Sub(now))

		select {
		case <-timer.C:
			d.rotateLogger(cfg)
		case <-d.quit:
			timer.Stop()
			return
		}
	}
}

// Stop 停止日志管理器
func (d *DailyLogger) Stop() {
	close(d.quit)
	d.mu.Lock()
	defer d.mu.Unlock()
	if d.logger != nil {
		_ = d.logger.Sync()
	}
}
