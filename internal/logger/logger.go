package logger

import (
	"fmt"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func New(level, format string) (*zap.Logger, error) {
	var cfg zap.Config
	if strings.EqualFold(format, "json") {
		cfg = zap.NewProductionConfig()
	} else {
		cfg = zap.NewDevelopmentConfig()
	}
	parsed, err := zapcore.ParseLevel(strings.ToLower(level))
	if err != nil {
		return nil, fmt.Errorf("invalid log level: %w", err)
	}
	cfg.Level = zap.NewAtomicLevelAt(parsed)
	return cfg.Build()
}
