package g

import (
	"fmt"
	"go.uber.org/zap"
)

////日志对象初始化
type ZapLogger struct {
	*zap.Logger
}

/*
重写日志
*/
func (logger *ZapLogger) Infof(format string, args ...interface{}) {
	logstr := fmt.Sprintf(format, args...)
	logger.Logger.Info(logstr)
}

func (logger *ZapLogger) Errorf(format string, args ...interface{}) {
	logstr := fmt.Sprintf(format, args...)
	logger.Logger.Error(logstr)
}

func (logger *ZapLogger) Warnf(format string, args ...interface{}) {
	logstr := fmt.Sprintf(format, args...)
	logger.Logger.Warn(logstr)
}
