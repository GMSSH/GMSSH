package main

import (
	"fmt"
	"time"

	"github.com/DemonZack/simplejrpc-go/core/glog"
	"go.uber.org/zap"
)

func main() {
	m := map[string]any{
		"path":                 "logs/",
		"file":                 "{Y-m-d}.log",
		"level":                "error",
		"stdout":               false,
		"StStatus":             0,
		"rotateBackupLimit":    7,
		"writerColorEnable":    true,
		"RotateBackupCompress": 9,
		"rotateExpire":         "1d",
		"Flag":                 44,
	}

	// config, err := LoadConfig("./testdata/config.json")
	config, err := glog.LoadConfig(m)
	if err != nil {
		panic(fmt.Sprintf("load config failed: %v", err))
	}

	// 初始化日志
	logger, err := glog.NewLogger(config)
	if err != nil {
		panic(fmt.Sprintf("init logger failed: %v", err))
	}
	defer logger.Sync() // 刷新缓冲区的日志

	// 使用示例
	logger.Info("Logger initialized successfully",
		zap.String("path", config.Path),
		zap.String("file", config.File),
		zap.Int("backupLimit", config.RotateBackupLimit),
	)

	logger.Warn("This is a warning message")
	logger.Error("This is an error message")

	// 结构化日志
	logger.Info("User logged in",
		zap.String("username", "john"),
		zap.Int("attempt", 3),
		zap.Duration("duration", time.Second*5),
	)

	zap.L().Info("ddddddddddddddddddddddddd")
}
