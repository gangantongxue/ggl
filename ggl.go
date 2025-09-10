// Copyright (c) 2025 gangantongxue. All rights reserved.
// Licensed under the MIT License.
package ggl

import (
	"go.uber.org/zap"
	"time"
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

func Err(err error) zap.Field {
	return zap.Error(err)
}

func Errs(key string, errs []error) zap.Field {
	return zap.Errors(key, errs)
}

func Str(key, val string) zap.Field {
	return zap.String(key, val)
}

func Strs(key string, vals []string) zap.Field {
	return zap.Strings(key, vals)
}

func ByteStr(key string, val []byte) zap.Field {
	return zap.ByteString(key, val)
}

func ByteStrs(key string, vals [][]byte) zap.Field {
	return zap.ByteStrings(key, vals)
}

func Bool(key string, val bool) zap.Field {
	return zap.Bool(key, val)
}

func Bools(key string, vals []bool) zap.Field {
	return zap.Bools(key, vals)
}

func Int(key string, val int) zap.Field {
	return zap.Int(key, val)
}

func Ints(key string, vals []int) zap.Field {
	return zap.Ints(key, vals)
}

func Int8(key string, val int8) zap.Field {
	return zap.Int8(key, val)
}

func Int8s(key string, vals []int8) zap.Field {
	return zap.Int8s(key, vals)
}

func Int16(key string, val int16) zap.Field {
	return zap.Int16(key, val)
}

func Int16s(key string, vals []int16) zap.Field {
	return zap.Int16s(key, vals)
}

func Int32(key string, val int32) zap.Field {
	return zap.Int32(key, val)
}

func Int32s(key string, vals []int32) zap.Field {
	return zap.Int32s(key, vals)
}

func Int64(key string, val int64) zap.Field {
	return zap.Int64(key, val)
}

func Int64s(key string, vals []int64) zap.Field {
	return zap.Int64s(key, vals)
}

func Uint(key string, val uint) zap.Field {
	return zap.Uint(key, val)
}

func Uints(key string, vals []uint) zap.Field {
	return zap.Uints(key, vals)
}

func Uint8(key string, val uint8) zap.Field {
	return zap.Uint8(key, val)
}

func Uint8s(key string, vals []uint8) zap.Field {
	return zap.Uint8s(key, vals)
}

func Uint16(key string, val uint16) zap.Field {
	return zap.Uint16(key, val)
}

func Uint16s(key string, vals []uint16) zap.Field {
	return zap.Uint16s(key, vals)
}

func Uint32(key string, val uint32) zap.Field {
	return zap.Uint32(key, val)
}

func Uint32s(key string, vals []uint32) zap.Field {
	return zap.Uint32s(key, vals)
}

func Uint64(key string, val uint64) zap.Field {
	return zap.Uint64(key, val)
}

func Uint64s(key string, vals []uint64) zap.Field {
	return zap.Uint64s(key, vals)
}

func Float32(key string, val float32) zap.Field {
	return zap.Float32(key, val)
}

func Float32s(key string, vals []float32) zap.Field {
	return zap.Float32s(key, vals)
}

func Float64(key string, val float64) zap.Field {
	return zap.Float64(key, val)
}

func Float64s(key string, vals []float64) zap.Field {
	return zap.Float64s(key, vals)
}

func Time(key string, val time.Time) zap.Field {
	return zap.Time(key, val)
}

func Times(key string, vals []time.Time) zap.Field {
	return zap.Times(key, vals)
}

func Duration(key string, val time.Duration) zap.Field {
	return zap.Duration(key, val)
}

func Durations(key string, vals []time.Duration) zap.Field {
	return zap.Durations(key, vals)
}

func Any(key string, val interface{}) zap.Field {
	return zap.Any(key, val)
}

func Namespace(key string) zap.Field {
	return zap.Namespace(key)
}

func Binary(key string, val []byte) zap.Field {
	return zap.Binary(key, val)
}

func Skip() zap.Field {
	return zap.Skip()
}
