// Copyright (c) 2025 gangantongxue. All rights reserved.
// Licensed under the MIT License.
package ggl

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"path/filepath"
	"sync"
	"time"
)

// Config 配置
type Config struct {
	LogFileDir    string // 日志文件存放目录
	LogMaxSize    int    // 日志文件最大大小，单位MB
	LogMaxBackups int    // 日志文件最大备份数量
	LogMaxAge     int    // 日志文件最大保留时间，单位天
	LogCompress   bool   // 是否压缩日志文件
	ToConsole     bool   // 是否同时输出到控制台
}

// DefaultConfig DefeatConfig 默认配置
func DefaultConfig() *Config {
	return &Config{
		LogFileDir:    "./log",
		LogMaxSize:    100,
		LogMaxBackups: 10,
		LogMaxAge:     7,
		LogCompress:   true,
		ToConsole:     false,
	}
}

// DailyLogger 日志管理器
type DailyLogger struct {
	mu     sync.Mutex
	logger *zap.Logger
	quit   chan struct{}
}

// getLogFileName 获取日志文件名
func getLogFileName(logFileDir string) string {
	return filepath.Join(logFileDir, fmt.Sprintf("landlady_log_%s.log", time.Now().Format("2006-01-02")))
}

// newZapLogger 初始化zap日志
func newZapLogger(cfg *Config) *zap.Logger {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   getLogFileName(cfg.LogFileDir),
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

// NewDailyLogger 创建DailyLogger实例
func NewDailyLogger(cfg *Config) *DailyLogger {
	d := &DailyLogger{
		quit: make(chan struct{}),
	}
	d.rotateLogger(cfg)
	go d.scheduleRotate(cfg)
	return d
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
		next := now.Add(24 * time.Hour).Truncate(24 * time.Hour)
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
