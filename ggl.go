// Copyright (c) 2025 gangantongxue. All rights reserved.
// Licensed under the MIT License.
package ggl

import (
	"go.uber.org/zap"
)

func Debug(msg string, fields ...zap.Field) {
	zap.L().Debug(msg, fields...)
}

func Info(msg string, fields ...zap.Field) {
	zap.L().Info(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	zap.L().Warn(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	zap.L().Error(msg, fields...)
}

func Panic(msg string, fields ...zap.Field) {
	zap.L().Panic(msg, fields...)
}

func Fatal(msg string, fields ...zap.Field) {
	zap.L().Fatal(msg, fields...)
}
