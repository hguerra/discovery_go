package logging

import "go.uber.org/zap"

var logger *zap.SugaredLogger

func init() {
	logger = GetLogger()
}

func Catch(err error) {
	if err != nil {
		logger.Panic(err)
	}
}
