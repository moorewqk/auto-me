package admin

import (
	"fmt"
	"gitee.com/moorewqk/antcom/server/cores/g"
	"gitee.com/moorewqk/antcom/server/utils"
	rotatelogfile "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"path"
	"time"
)

//
var (
	level zapcore.Level
	//coreLog *ZapLogger
)

//
////日志对象初始化
//type ZapLogger struct {
//	zap.Logger
//}

func NewZapLogger() (logger *zap.Logger) {
	//var (
	//	logger *zap.Logger
	//)

	if ok, _ := utils.PathExists(g.GV_SERVER.Zap.Director); !ok { // 判断是否有Director文件夹
		utils.PfFail("create %v directory", g.GV_SERVER.Zap.Director)
		_ = os.Mkdir(g.GV_SERVER.Zap.Director, os.ModePerm)
	}

	switch g.GV_SERVER.Zap.Level { // 初始化配置文件的Level
	case "debug":
		level = zap.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "warn":
		level = zap.WarnLevel
	case "error":
		level = zap.ErrorLevel
	case "dpanic":
		level = zap.DPanicLevel
	case "panic":
		level = zap.PanicLevel
	case "fatal":
		level = zap.FatalLevel
	default:
		level = zap.InfoLevel
	}

	if level == zap.DebugLevel || level == zap.ErrorLevel {
		logger = zap.New(getEncoderCore(), zap.AddStacktrace(level))
	} else {
		logger = zap.New(getEncoderCore())
	}
	if g.GV_SERVER.Zap.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}
	return logger
}

//
//// getEncoderConfig 获取zapcore.EncoderConfig
func getEncoderConfig() (config zapcore.EncoderConfig) {
	config = zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "caller",
		StacktraceKey:  g.GV_SERVER.Zap.StacktraceKey,
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     CustomTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
	}
	switch {
	case g.GV_SERVER.Zap.EncodeLevel == "LowercaseLevelEncoder": // 小写编码器(默认)
		config.EncodeLevel = zapcore.LowercaseLevelEncoder
	case g.GV_SERVER.Zap.EncodeLevel == "LowercaseColorLevelEncoder": // 小写编码器带颜色
		config.EncodeLevel = zapcore.LowercaseColorLevelEncoder
	case g.GV_SERVER.Zap.EncodeLevel == "CapitalLevelEncoder": // 大写编码器
		config.EncodeLevel = zapcore.CapitalLevelEncoder
	case g.GV_SERVER.Zap.EncodeLevel == "CapitalColorLevelEncoder": // 大写编码器带颜色
		config.EncodeLevel = zapcore.CapitalColorLevelEncoder
	default:
		config.EncodeLevel = zapcore.LowercaseLevelEncoder
	}
	return config
}

//
//// getEncoder 获取zapcore.Encoder
func getEncoder() zapcore.Encoder {
	if g.GV_SERVER.Zap.Format == "json" {
		return zapcore.NewJSONEncoder(getEncoderConfig())
	}
	return zapcore.NewConsoleEncoder(getEncoderConfig())
}

//异步日志分割
func WriteSyncer() (zapcore.WriteSyncer, error) {
	fileWriter, err := rotatelogfile.New(
		path.Join(g.GV_SERVER.Zap.Director, "%Y-%m-%d.log"),
		rotatelogfile.WithLinkName(g.GV_SERVER.Zap.LinkName),
		rotatelogfile.WithMaxAge(7*24*time.Hour),
		rotatelogfile.WithRotationTime(24*time.Hour),
	)
	if g.GV_SERVER.Zap.LogInConsole {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(fileWriter)), err
	}
	return zapcore.AddSync(fileWriter), err
}

// getEncoderCore 获取Encoder的zapcore.Core
func getEncoderCore() (core zapcore.Core) {
	writer, err := WriteSyncer() // 使用file-rotatelogs进行日志分割
	if err != nil {
		fmt.Printf("Get Write Syncer Failed err:%v", err.Error())
		return
	}
	return zapcore.NewCore(getEncoder(), writer, level)
}

//
//// 自定义日志输出时间格式
func CustomTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format(g.GV_SERVER.Zap.Prefix + "2006/01/02 - 15:04:05.000"))
}
