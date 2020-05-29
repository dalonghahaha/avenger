package logger

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"os"

	"github.com/logrusorgru/aurora"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/dalonghahaha/avenger/tools/file"
)

var (
	inited  = false
	level   = logrus.InfoLevel
	dir     = "logs"
	console = true
	logger  *logrus.Logger
	loggers map[string]*logrus.Logger
)

type FileFormatter struct {
}

func (f *FileFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	data := make(logrus.Fields)
	data["time"] = entry.Time
	data["message"] = entry.Message
	for key, val := range entry.Data {
		data[key] = val
	}
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}
	encoder := json.NewEncoder(b)
	if err := encoder.Encode(data); err != nil {
		return nil, fmt.Errorf("failed to marshal fields to JSON, %v", err)
	}
	return b.Bytes(), nil
}

type ConsoleFormatter struct {
}

func (f *ConsoleFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	data := fmt.Sprintf("%s - [%s]:%s\r\n", entry.Time.Format("2006-01-02 15:04:05.000000"), entry.Level, entry.Message)
	if entry.Level == logrus.InfoLevel {
		return []byte(aurora.Cyan(data).String()), nil
	}
	if entry.Level == logrus.DebugLevel {
		return []byte(aurora.Green(data).String()), nil
	}
	if entry.Level == logrus.ErrorLevel {
		return []byte(aurora.Red(data).String()), nil
	}
	return []byte(data), nil
}

func config() {
	if inited {
		return
	}
	if viper.GetString("component.log.level") != "" {
		_level := viper.GetString("component.log.level")
		switch _level {
		case "debug":
			level = logrus.DebugLevel
		case "info":
			level = logrus.InfoLevel
		case "error":
			level = logrus.ErrorLevel
		case "warn":
			level = logrus.WarnLevel
		}
	}
	if viper.GetString("component.log.dir") != "" {
		dir = viper.GetString("component.log.dir")
	}
	if !viper.GetBool("component.log.console") {
		console = viper.GetBool("component.log.console")
	}
}

func Register() error {
	config()
	if !file.Exists(dir) {
		err := file.Mkdir(dir)
		if err != nil {
			return fmt.Errorf("创建日志目录失败:%s", err.Error())
		}
	}
	logger = logrus.New()
	//控制台输出设置
	if !console {
		src, err := os.OpenFile(os.DevNull, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		if err != nil {
			return fmt.Errorf("日志组件异常")
		}
		logger.Out = bufio.NewWriter(src)
	}
	logger.SetLevel(level)
	logger.SetFormatter(&ConsoleFormatter{})
	pathMap := lfshook.PathMap{
		logrus.InfoLevel:  fmt.Sprintf("%s/info.log", dir),
		logrus.DebugLevel: fmt.Sprintf("%s/debug.log", dir),
		logrus.ErrorLevel: fmt.Sprintf("%s/error.log", dir),
		logrus.WarnLevel:  fmt.Sprintf("%s/warn.log", dir),
	}
	hook := lfshook.NewHook(pathMap, &FileFormatter{})
	logger.AddHook(hook)
	return nil
}

func GetLogger(name string) (*logrus.Logger, error) {
	config()
	if _logger, ok := loggers[name]; ok {
		return _logger, nil
	}
	if !file.Exists(dir) {
		err := file.Mkdir(dir)
		if err != nil {
			return nil, fmt.Errorf("创建日志目录失败:%s", err.Error())
		}
	}
	logger = logrus.New()
	if !console {
		src, err := os.OpenFile(os.DevNull, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		if err != nil {
			return nil, fmt.Errorf("日志组件异常")
		}
		logger.Out = bufio.NewWriter(src)
	}
	logger.SetLevel(level)
	logger.SetFormatter(&ConsoleFormatter{})
	pathMap := lfshook.PathMap{
		logrus.InfoLevel:  fmt.Sprintf("%s/%s.log", dir, name),
		logrus.DebugLevel: fmt.Sprintf("%s/%s.log", dir, name),
		logrus.ErrorLevel: fmt.Sprintf("%s/%s.log", dir, name),
		logrus.WarnLevel:  fmt.Sprintf("%s/%s.log", dir, name),
	}
	hook := lfshook.NewHook(pathMap, &FileFormatter{})
	logger.AddHook(hook)
	loggers[name] = logger
	return logger, nil
}

func Info(v ...interface{}) {
	logger.Info(v...)
}

func Debug(v ...interface{}) {
	logger.Debug(v...)
}

func Error(v ...interface{}) {
	logger.Error(v...)
}

func Warn(v ...interface{}) {
	logger.Warn(v...)
}

func Infof(format string, v ...interface{}) {
	logger.Infof(format, v...)
}

func Debugf(format string, v ...interface{}) {
	logger.Debugf(format, v...)
}

func Errorf(format string, v ...interface{}) {
	logger.Errorf(format, v...)
}

func Warnf(format string, v ...interface{}) {
	logger.Warnf(format, v...)
}

func InfoData(message string, data map[string]interface{}) {
	fields := logrus.Fields{}
	for k, v := range data {
		fields[k] = v
	}
	logger.WithFields(fields).Info(message)
}

func DebugData(message string, data map[string]interface{}) {
	fields := logrus.Fields{}
	for k, v := range data {
		fields[k] = v
	}
	logger.WithFields(fields).Debug(message)
}

func ErrorData(message string, data map[string]interface{}) {
	fields := logrus.Fields{}
	for k, v := range data {
		fields[k] = v
	}
	logger.WithFields(fields).Error(message)
}

func WarnData(message string, data map[string]interface{}) {
	fields := logrus.Fields{}
	for k, v := range data {
		fields[k] = v
	}
	logger.WithFields(fields).Warn(message)
}
