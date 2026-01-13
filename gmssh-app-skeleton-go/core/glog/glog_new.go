package glog

import (
	"fmt"
	"os"
	"path/filepath"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// NewLogger creates and configures a new zap.Logger instance based on the provided LogConfig
// It handles:
// - Directory creation
// - Log file rotation
// - Output configuration (file/stdout)
// - Log level setup
// - Log formatting and coloring
//
// Parameters:
//   - config: Pointer to LogConfig containing all logging configuration
//
// Returns:
//   - *zap.Logger: Configured logger instance
//   - error: Any error that occurred during setup
func NewLogger(config *LogConfig) (*zap.Logger, error) {
	// Ensure log directory exists with proper permissions (0755: owner read/write/execute, group/others read/execute)
	if err := os.MkdirAll(config.Path, 0755); err != nil {
		return nil, fmt.Errorf("create log directory failed: %v", err)
	}

	// Process date placeholders in filename (e.g., replace {Y-m-d} with current date)
	filename := replaceDatePlaceholders(config.File)
	fullPath := filepath.Join(config.Path, filename)

	// Parse and set log level from config (defaults to InfoLevel if invalid)
	var level zapcore.Level
	if err := level.UnmarshalText([]byte(config.Level)); err != nil {
		level = zapcore.InfoLevel // Default level if parsing fails
	}

	// Configure log rotation using lumberjack:
	// - MaxSize: 10MB per log file
	// - MaxBackups: Number of old log files to retain (from config)
	// - MaxAge: Days to retain logs (converted from hours in config)
	// - Compress: Whether to compress rotated logs (based on config)
	lumberJackLogger := &lumberjack.Logger{
		Filename:   fullPath,
		MaxSize:    10, // Megabytes
		MaxBackups: config.RotateBackupLimit,
		MaxAge:     parseRotateExpire(config.RotateExpire) / 24, // Convert hours to days
		Compress:   config.RotateBackupCompress > 0,             // Enable compression if value > 0
	}

	// Configure output destination(s)
	var ws zapcore.WriteSyncer
	if config.Stdout {
		// Write to both stdout and log file if Stdout enabled
		ws = zapcore.NewMultiWriteSyncer(
			zapcore.AddSync(os.Stdout),
			zapcore.AddSync(lumberJackLogger),
		)
	} else {
		// Write only to log file
		ws = zapcore.AddSync(lumberJackLogger)
	}

	// Configure log message encoding
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder // Use ISO8601 timestamp format

	// Enable colored level output if configured
	if config.WriterColorEnable {
		encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	} else {
		encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	}

	// Create the core logger with:
	// - Console encoder with our configuration
	// - Configured output destination(s)
	// - Specified log level
	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderConfig),
		ws,
		level,
	)

	// Create the final logger with caller information enabled
	logger := zap.New(core, zap.AddCaller())

	// Replace the global logger instance
	zap.ReplaceGlobals(logger)
	return logger, nil
}
