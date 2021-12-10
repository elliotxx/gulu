package log

import (
	"fmt"
	"os"
	"strings"
)

type Level uint8

var log Logger

const (
	FATAL Level = iota
	ERROR
	WARN
	INFO
	DEBUG
)

func init() {
	config := NewDefaultConfiguration()
	SetLogger(config)
}

// 根据日志配置初始化 Logger
func SetLogger(config *Configuration) {
	logger, err := newZapLogger(config)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to set logger by configuration %v as %s\n", config, err)
		os.Exit(1)
	}
	log = logger
}

// 日志配置
type Configuration struct {
	Name  string `xml:"name,omitempty" json:"name,omitempty"`   // name: 日志名称，用于生成日志文件名等，比如 name = "app"，则生成的日志文件为 app.log, app_error.log, app_debug.log
	Level Level  `xml:"level,omitempty" json:"level,omitempty"` // level: 日志级别， 可选：DEBUG INFO WARN ERROR
	Path  string `xml:"path,omitempty" json:"path,omitempty"`   // path： 日志路径，比如 ./logs
	Count uint16 `xml:"count,omitempty" json:"count,omitempty"` // count: 日志文件数量
	Size  uint32 `xml:"size,omitempty" json:"size,omitempty"`   // size: 日志文件大小， 单位KB
}

// 生成默认日志配置
func NewDefaultConfiguration() *Configuration {
	return &Configuration{
		Name:  "app",
		Level: INFO,
		Path:  "./logs/",
		Count: 5,
		Size:  1024,
	}
}

// Logger 接口
type Logger interface {
	Debugf(format string, args ...interface{})
	Debug(args ...interface{})
	Infof(format string, args ...interface{})
	Info(args ...interface{})
	Warnf(format string, args ...interface{})
	Warn(args ...interface{})
	Errorf(format string, args ...interface{})
	Error(args ...interface{})
	Panicf(format string, args ...interface{})
	Panic(args ...interface{})
	Fatalf(format string, args ...interface{})
	Fatal(args ...interface{})
	SetLevel(level Level)
	GetLogFile() LogFile
}

func Debugf(format string, args ...interface{}) {
	log.Debugf(format, args...)
}

func Debug(args ...interface{}) {
	log.Debug(args...)
}

func Infof(format string, args ...interface{}) {
	log.Infof(format, args...)
}

func Info(args ...interface{}) {
	log.Info(args...)
}

func Warnf(format string, args ...interface{}) {
	log.Warnf(format, args...)
}

func Warn(args ...interface{}) {
	log.Warn(args...)
}

func Errorf(format string, args ...interface{}) {
	log.Errorf(format, args...)
}

func Error(args ...interface{}) {
	log.Error(args...)
}

func Panicf(format string, args ...interface{}) {
	log.Panicf(format, args...)
}

func Panic(args ...interface{}) {
	log.Panic(args...)
}

func Fatalf(format string, args ...interface{}) {
	log.Fatalf(format, args...)
}

func Fatal(args ...interface{}) {
	log.Fatal(args...)
}

func SetLevel(level Level) {
	log.SetLevel(level)
}

func GetLevelFromStr(level string) Level {
	switch strings.ToUpper(level) {
	case "INFO":
		return INFO
	case "FATAL":
		return FATAL
	case "ERROR":
		return ERROR
	case "WARN":
		return WARN
	case "DEBUG":
		return DEBUG
	default:
		log.Info("user set log level is invalid, using default info level")
		return INFO
	}
}

func GetLogFile() LogFile {
	return log.GetLogFile()
}
