package logging

import (
	"log"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func GetLogger() *zap.SugaredLogger {
	loggerConfig := zap.NewProductionConfig()
	loggerConfig.EncoderConfig.TimeKey = "timestamp"
	loggerConfig.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.RFC3339)

	logger, err := loggerConfig.Build()
	if err != nil {
		log.Fatal(err)
	}

	return logger.Sugar()
}
