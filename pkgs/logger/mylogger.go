package logger

import "go.uber.org/zap"

type MyLogger struct {
	logger     *zap.SugaredLogger
	cfgPath    string
	logDirName string
}

// 构造函数
// cfgPath: 配置文件路径
// logDirName:  日志目录名
func NewMyLogger(cfgPath string, logDirName string) *MyLogger {
	logger := &MyLogger{
		cfgPath:    cfgPath,
		logDirName: logDirName,
	}
	logger.Logger()
	return logger
}

// 对外接口
func (mylog *MyLogger) Logger() *zap.SugaredLogger {
	if mylog.logger == nil {
		mylog.logger = MakeLogger(mylog.cfgPath, mylog.logDirName)
	}
	return mylog.logger
}
func (mylog *MyLogger) Debug(args ...interface{}) {
	mylog.Logger().WithOptions(zap.AddCallerSkip(1)).Debug(args...)
}

func (mylog *MyLogger) Debugf(template string, args ...interface{}) {
	mylog.Logger().WithOptions(zap.AddCallerSkip(1)).Debugf(template, args...)
}

func (mylog *MyLogger) Info(args ...interface{}) {
	mylog.Logger().WithOptions(zap.AddCallerSkip(1)).Info(args...)
}

func (mylog *MyLogger) Infof(template string, args ...interface{}) {
	mylog.Logger().WithOptions(zap.AddCallerSkip(1)).Infof(template, args...)
}

func (mylog *MyLogger) Warn(args ...interface{}) {
	mylog.Logger().WithOptions(zap.AddCallerSkip(1)).Warn(args...)
}

func (mylog *MyLogger) Warnf(template string, args ...interface{}) {
	mylog.Logger().WithOptions(zap.AddCallerSkip(1)).Warnf(template, args...)
}

func (mylog *MyLogger) Error(args ...interface{}) {
	mylog.Logger().WithOptions(zap.AddCallerSkip(1)).Error(args...)
}

func (mylog *MyLogger) Errorf(template string, args ...interface{}) {
	mylog.Logger().WithOptions(zap.AddCallerSkip(1)).Errorf(template, args...)
}

func (mylog *MyLogger) Panic(args ...interface{}) {
	mylog.Logger().WithOptions(zap.AddCallerSkip(1)).Panic(args...)
}

func (mylog *MyLogger) Panicf(template string, args ...interface{}) {
	mylog.Logger().WithOptions(zap.AddCallerSkip(1)).Panicf(template, args...)
}
