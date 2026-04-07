package logger

import (
	"os"
	"strings"

	"github.com/DenisMekh/mini-transfer-system/account-svc/internal/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ZapLogger struct {
	logger *zap.Logger
}

var zapLogLevelMapping = map[string]zapcore.Level{
	"debug": zapcore.DebugLevel,
	"info":  zapcore.InfoLevel,
	"warn":  zapcore.WarnLevel,
	"error": zapcore.ErrorLevel,
	"fatal": zapcore.FatalLevel,
}

// getLogLevel function for getting logger level
func getLogLevel(cfg *config.LogConfig) zapcore.Level {
	level, exists := zapLogLevelMapping[cfg.Level]
	if !exists {
		return zapcore.InfoLevel
	}
	return level
}

// getEncoder function for getting encoder Config
func getEncoder(cfg *config.LogConfig) zapcore.Encoder {
	var encoderCfg zapcore.EncoderConfig
	if strings.ToLower(cfg.Mode) == "development" {
		encoderCfg = zap.NewDevelopmentEncoderConfig()
	} else {
		encoderCfg = zap.NewProductionEncoderConfig()
	}

	var encoder zapcore.Encoder
	encoderCfg.LevelKey = "LEVEL"
	encoderCfg.CallerKey = "CALLER"
	encoderCfg.TimeKey = "TIME"
	encoderCfg.NameKey = "NAME"
	encoderCfg.MessageKey = "MESSAGE"

	if strings.ToLower(cfg.Encoding) == "console" {
		encoder = zapcore.NewConsoleEncoder(encoderCfg)
	} else {
		encoder = zapcore.NewJSONEncoder(encoderCfg)
	}
	return encoder
}

// New function for creating new example of zapLogger
func New(cfg *config.Config) *ZapLogger {
	level := getLogLevel(&cfg.Log)
	encoder := getEncoder(&cfg.Log)
	core := zapcore.NewCore(encoder, zapcore.Lock(os.Stdout), level)
	logger := zap.New(core, zap.AddCaller())
	return &ZapLogger{
		logger: logger,
	}
}

// Sync method for syncing logger
func (l *ZapLogger) Sync() error {
	return l.logger.Sync()
}

func (l *ZapLogger) Debug(msg string, fields ...zap.Field) {
	l.logger.Debug(msg, fields...)
}

func (l *ZapLogger) Info(msg string, fields ...zap.Field) {
	l.logger.Info(msg, fields...)
}

func (l *ZapLogger) Warn(msg string, fields ...zap.Field) {
	l.logger.Warn(msg, fields...)
}

func (l *ZapLogger) Error(msg string, fields ...zap.Field) {
	l.logger.Error(msg, fields...)
}

func (l *ZapLogger) Fatal(msg string, fields ...zap.Field) {
	l.logger.Fatal(msg, fields...)
}
