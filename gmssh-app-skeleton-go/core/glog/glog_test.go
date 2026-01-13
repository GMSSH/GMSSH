package glog

import (
	"fmt"
	"testing"
	"time"

	// "gmbagent/core"
	"go.uber.org/zap"
)

func TestGlog(t *testing.T) {
	// 加载配置
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
	config, err := LoadConfig(m)
	if err != nil {
		panic(fmt.Sprintf("load config failed: %v", err))
	}

	// 初始化日志
	logger, err := NewLogger(config)
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

// TestGLoggerErrorWithStack 测试GLogger的Error方法是否能正确记录堆栈信息
func TestGLoggerErrorWithStack(t *testing.T) {
	// 加载配置
	m := map[string]any{
		"path":                 "logs/",
		"file":                 "{Y-m-d}.log",
		"level":                "debug",
		"stdout":               false,
		"StStatus":             0,
		"rotateBackupLimit":    7,
		"writerColorEnable":    true,
		"RotateBackupCompress": 9,
		"rotateExpire":         "1d",
		"Flag":                 44,
	}
	// config, err := LoadConfig("./testdata/config.json")
	config, err := LoadConfig(m)
	if err != nil {
		t.Fatalf("load config failed: %v", err)
	}

	// 初始化日志
	logger, err := NewLogger(config)
	if err != nil {
		t.Fatalf("init logger failed: %v", err)
	}
	defer logger.Sync()

	// 创建GLogger实例
	gLogger := NewGLogger(logger)

	// 测试Error方法 - 这会触发堆栈信息记录
	gLogger.Error("这是一个测试错误消息，应该包含堆栈信息")

	// 测试ErrorWithStack方法 - 这会记录所有goroutines的堆栈
	gLogger.ErrorWithStack("测试完整堆栈信息记录")

	t.Log("GLogger堆栈测试完成，请检查日志文件中是否包含堆栈信息")
}
