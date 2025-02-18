package logger

import (
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"

	"ABTest/pkgs/utils"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

const (
	LowercaseLevelEncoder      = "Lowercase"
	LowercaseColorLevelEncoder = "LowercaseColor"
	CapitalLevelEncoder        = "Capital"
	CapitalColorLevelEncoder   = "CapitalColor"
	EncodeLevelCustomer        = "Customer"
)

const (
	FileDebug  = "/debug.log"
	FileInfo   = "/info.log"
	FileWarn   = "/warn.log"
	FileError  = "/error.log"
	FilePanic  = "/panic.log"
	FileStderr = "/stderr.log"
)

const (
	DefaultConfig    = "./configs/logger.yaml"
	DefaultDirectory = "ABTest"
)

const (
	CallerDepth = 8
)

const (
	CallerTypeNormal   = 0
	CallerTypeSeparate = 1
	CallerTypeFull     = 2
)

type Zap struct {
	Encoder       string  `yaml:"encoder"`
	Path          string  `yaml:"path"`
	Directory     string  `yaml:"directory"`
	ShowLine      bool    `yaml:"show_line"`
	EncodeLevel   string  `yaml:"encode_level"`
	StacktraceKey string  `yaml:"stacktrace_key"`
	LogStdout     bool    `yaml:"log_stdout"`
	Segment       Segment `yaml:"segment"`
	CallerType    int     `yaml:"caller_type"`
}

type Segment struct {
	MaxSize    int  `yaml:"max_size"`
	MaxAge     int  `yaml:"max_age"`
	MaxBackups int  `yaml:"max_backups"`
	Compress   bool `yaml:"compress"`
}

var (
	logger *zap.SugaredLogger
)

func MakeLogger(cfgPath string, directory string) {
	var (
		logCfg = new(Zap)
	)
	err := utils.ReadYamlFile(cfgPath, logCfg)
	if err != nil {
		panic(err)
	}
	logCfg.Directory = directory
	logger = newLogger(logCfg)
}

func newLogger(cfg *Zap) (sl *zap.SugaredLogger) {
	// zap.LevelEnablerFunc(func(lev zapcore.Level) bool 用来划分不同级别的输出
	// 根据不同的级别输出到不同的日志文件

	// 调试级别
	debugLevel := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level == zap.DebugLevel
	})
	// 日志级别
	infoLevel := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level == zap.InfoLevel
	})
	// 警告级别
	warnLevel := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level == zap.WarnLevel
	})
	// 错误级别
	errorLevel := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level == zap.ErrorLevel
	})
	// panic级别
	panicLevel := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level >= zap.DPanicLevel
	})

	path := cfg.Path + cfg.Directory
	if !isDir(path) {
		mkdir(path)
	}

	cores := [...]zapcore.Core{
		getEncoderCore(path+FileDebug, debugLevel, cfg),
		getEncoderCore(path+FileInfo, infoLevel, cfg),
		getEncoderCore(path+FileWarn, warnLevel, cfg),
		getEncoderCore(path+FileError, errorLevel, cfg),
		getEncoderCore(path+FilePanic, panicLevel, cfg),
	}

	//zapcore.NewTee(cores ...zapcore.Core) zapcore.Core
	//NewTee创建一个Core，将日志条目复制到两个或更多的底层Core中

	logger := zap.New(zapcore.NewTee(cores[:]...))
	//用文件名、行号和zap调用者的函数名注释每条消息
	if cfg.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}
	sl = logger.Sugar()
	sl.Sync()
	return
}

func isDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}
func mkdir(path string) (err error) {
	err = os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return
	}
	err = os.Chmod(path, os.ModePerm)
	return
}

func getEncoderCore(filename string, level zapcore.LevelEnabler, cfg *Zap) (core zapcore.Core) {
	// 使用lumberjack进行日志分割
	writer := getWriteSyncer(filename, cfg)
	return zapcore.NewCore(getEncoder(cfg), writer, level)
}

func getWriteSyncer(filename string, cfg *Zap) zapcore.WriteSyncer {
	hook := &lumberjack.Logger{
		Filename:   filename,               // 日志文件的位置
		MaxSize:    cfg.Segment.MaxSize,    // 在进行切割之前，日志文件的最大大小（以MB为单位）
		MaxBackups: cfg.Segment.MaxBackups, // 保留旧文件的最大个数
		MaxAge:     cfg.Segment.MaxAge,     // 保留旧文件的最大天数
		Compress:   cfg.Segment.Compress,   // 是否压缩/归档旧文件
	}
	if cfg.LogStdout {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(hook))
	}
	return zapcore.AddSync(hook)
}

func getEncoder(cfg *Zap) zapcore.Encoder {
	switch cfg.Encoder {
	case "json":
		return zapcore.NewJSONEncoder(getEncoderConfig(cfg))
	case "console":
		return zapcore.NewConsoleEncoder(getEncoderConfig(cfg))
	}
	return zapcore.NewConsoleEncoder(getEncoderConfig(cfg))
}

func getEncoderConfig(cfg *Zap) (config zapcore.EncoderConfig) {
	config = zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "caller",
		StacktraceKey:  cfg.StacktraceKey,              // 栈名
		LineEnding:     zapcore.DefaultLineEnding,      // 默认的结尾\n
		EncodeTime:     customTimeEncoder,              // 时间格式 zapcore.ISO8601TimeEncoder
		EncodeDuration: zapcore.SecondsDurationEncoder, // 编码间隔
		//EncodeCaller:   shortCallerEncoder,           // 绝对路径:zapcore.FullCallerEncoder,相对路径:zapcore.ShortCallerEncoder
		EncodeName: zapcore.FullNameEncoder,
	}
	switch cfg.CallerType {
	case CallerTypeNormal:
		config.EncodeCaller = shortCallerEncoder
	case CallerTypeSeparate:
		config.EncodeCaller = zapcore.ShortCallerEncoder // 绝对路径:zapcore.FullCallerEncoder,相对路径:zapcore.ShortCallerEncoder
	case CallerTypeFull:
		config.EncodeCaller = customEncodeCaller
	}

	switch cfg.EncodeLevel {
	case LowercaseLevelEncoder:
		// 小写编码器(默认)
		config.EncodeLevel = zapcore.LowercaseLevelEncoder
	case LowercaseColorLevelEncoder:
		// 小写编码器带颜色
		config.EncodeLevel = zapcore.LowercaseColorLevelEncoder
	case CapitalLevelEncoder:
		// 大写编码器
		config.EncodeLevel = zapcore.CapitalLevelEncoder
	case CapitalColorLevelEncoder:
		// 大写编码器带颜色
		config.EncodeLevel = zapcore.CapitalColorLevelEncoder
	case EncodeLevelCustomer:
		config.EncodeLevel = customEncodeLevel
	default:
		config.EncodeLevel = zapcore.LowercaseLevelEncoder
	}
	return config
}

// 自定义日志输出时间格式
func customTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString("[" + t.Format("2006-01-02 15:04:05.000") + "]")
}

// 自定义日志级别显示
func customEncodeLevel(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString("[" + level.CapitalString() + "]")
}

// 自定义行号显示
func customEncodeCaller(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString("[" + caller.TrimmedPath() + "]")
}

func shortCallerEncoder(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(getCaller(CallerDepth))
}

func getCaller(callDepth int) string {
	_, file, line, ok := runtime.Caller(callDepth)
	if !ok {
		return ""
	}
	return prettyCaller(file, line)
}

func prettyCaller(file string, line int) string {
	idx := strings.LastIndexByte(file, '/')
	if idx < 0 {
		return fmt.Sprintf("%s:%d", file, line)
	}

	idx = strings.LastIndexByte(file[:idx], '/')
	if idx < 0 {
		return fmt.Sprintf("%s:%d", file, line)
	}
	return fmt.Sprintf("%s:%d", file[idx+1:], line)
}

// 对外接口
func Logger() *zap.SugaredLogger {
	if logger == nil {
		MakeLogger(DefaultConfig, DefaultDirectory)
	}
	return logger
}

func Debug(args ...interface{}) {
	Logger().WithOptions(zap.AddCallerSkip(1)).Debug(args...)
}

func Debugf(template string, args ...interface{}) {
	Logger().WithOptions(zap.AddCallerSkip(1)).Debugf(template, args...)
}

func Info(args ...interface{}) {
	Logger().WithOptions(zap.AddCallerSkip(1)).Info(args...)
}

func Infof(template string, args ...interface{}) {
	Logger().WithOptions(zap.AddCallerSkip(1)).Infof(template, args...)
}

func Warn(args ...interface{}) {
	Logger().WithOptions(zap.AddCallerSkip(1)).Warn(args...)
}

func Warnf(template string, args ...interface{}) {
	Logger().WithOptions(zap.AddCallerSkip(1)).Warnf(template, args...)
}

func Error(args ...interface{}) {
	Logger().WithOptions(zap.AddCallerSkip(1)).Error(args...)
}

func Errorf(template string, args ...interface{}) {
	Logger().WithOptions(zap.AddCallerSkip(1)).Errorf(template, args...)
}

func Panic(args ...interface{}) {
	Logger().WithOptions(zap.AddCallerSkip(1)).Panic(args...)
}
