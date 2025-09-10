# ggl
一个基于 zap 和 lumberjack 封装的高性能 Go 日志库，提供了日志按天轮转、多级别日志记录等功能。

## 功能特点
- 高性能：基于 zap 实现，保持了 zap 的高性能特性
- 日志按天轮转：每天自动生成新的日志文件
- 支持多种日志级别：Debug、Info、Warn、Err、Panic、Fatal
- 可配置：支持自定义日志文件目录、大小、备份数量、保留时间等
- 多目标输出：可以同时输出到文件和控制台
- 全局替换：可以替换 zap 的全局日志器

## 安装
使用 Go Module 安装：

```bash
go get github.com/gangantongxue/ggl
```

## 快速开始

### 基本使用

```go
package main

import (
	"errors"
	"github.com/gangantongxue/ggl"
)

func main() {
	// 使用默认配置初始化日志库
	logger := ggl.NewDailyLogger(ggl.DefaultConfig())
	defer logger.Stop()

	// 使用各种日志级别
	ggl.Debug("这是一条调试日志", ggl.Str("key", "value"))
	ggl.Info("这是一条信息日志", ggl.Str("key", "value"))
	ggl.Warn("这是一条警告日志", ggl.Str("key", "value"))
	ggl.Error("这是一条错误日志", ggl.Err(errors.New("这是一个错误")))
	ggl.Panic("这是一条严重错误日志", ggl.Str("key", "value"))
	ggl.Fatal("这是一条致命错误日志", ggl.Str("key", "value"))
}
```

### 自定义配置

```go
package main

import (
    "github.com/gangantongxue/ggl"
)

func main() {
    // 自定义配置
    config := &ggl.Config{
        LogFileDir:    "./logs",       // 日志文件目录
        LogMaxSize:    50,             // 单个日志文件最大大小(MB)
        LogMaxBackups: 30,             // 最大备份数量
        LogMaxAge:     15,             // 最大保留天数
        LogCompress:   true,           // 是否压缩备份文件
        ToConsole:     true,           // 同时输出到控制台
    }
    
    logger := ggl.NewDailyLogger(config)
    defer logger.Stop()
    
    ggl.Info("使用自定义配置的日志")
}
```

## 配置详解

`Config` 结构体包含以下字段：

| 字段名 | 类型 | 说明 | 默认值 |
|-------|------|------|-------|
| LogFileDir | string | 日志文件存放目录 | "./log" |
| LogMaxSize | int | 单个日志文件最大大小，单位MB | 100 |
| LogMaxBackups | int | 日志文件最大备份数量 | 10 |
| LogMaxAge | int | 日志文件最大保留时间，单位天 | 7 |
| LogCompress | bool | 是否压缩日志文件 | true |
| ToConsole | bool | 是否同时输出到控制台 | false |

## 日志方法

ggl 提供了以下日志级别方法：

- `Debug(msg string, fields ...zap.Field)`: 记录调试信息
- `Info(msg string, fields ...zap.Field)`: 记录一般信息
- `Warn(msg string, fields ...zap.Field)`: 记录警告信息
- `Err(msg string, fields ...zap.Field)`: 记录错误信息
- `Panic(msg string, fields ...zap.Field)`: 记录严重错误并触发 panic
- `Fatal(msg string, fields ...zap.Field)`: 记录致命错误并退出程序

每个方法都支持传入额外的字段信息，使用 zap 的 Field 类型。

## 高级特性

### 日志按天轮转

日志库会自动在每天 0 点创建新的日志文件，文件名格式为 `landlady_log_YYYY-MM-DD.log`。

### 全局替换

通过 `zap.ReplaceGlobals(d.logger)` 方法，ggl 会替换 zap 的全局日志器，这样你也可以直接使用 zap 的全局方法记录日志：

```go
import "go.uber.org/zap"

// 初始化后可以直接使用 zap 的全局方法
zap.L().Info("这也会使用 ggl 配置的日志器")
```

## 注意事项

1. 确保日志目录存在，否则可能无法创建日志文件
2. 在程序退出前调用 `logger.Stop()` 方法，确保所有日志都被写入文件
3. Panic 和 Fatal 级别的日志会导致程序退出，请谨慎使用
4. 日志文件命名目前硬编码为 `landlady_log_`，如果需要修改，请修改源码中的 `getLogFileName` 函数

## 依赖

- [go.uber.org/zap](https://github.com/uber-go/zap): 高性能日志库
- [gopkg.in/natefinch/lumberjack.v2](https://github.com/natefinch/lumberjack): 日志文件轮转管理

## License

[MIT](LICENSE)
