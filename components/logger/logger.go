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

var logger *logrus.Logger

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

func Register() error {
	loggerLevel := viper.GetString("component.log.level")
	loggerDir := viper.GetString("component.log.dir")
	loggerConsole := viper.GetBool("component.log.console")
	if loggerLevel == "" || loggerDir == "" {
		return fmt.Errorf("日志配置异常")
	}
	if !file.Exists(loggerDir) {
		err := file.Mkdir(loggerDir)
		if err != nil {
			return fmt.Errorf("创建日志目录失败:%s", err.Error())
		}
	}
	logger = logrus.New()
	//日志级别设置
	var level logrus.Level
	switch loggerLevel {
	case "debug":
		level = logrus.DebugLevel
	case "info":
		level = logrus.InfoLevel
	case "warn":
		level = logrus.WarnLevel
	default:
		level = logrus.ErrorLevel
	}
	//控制台输出设置
	if !loggerConsole {
		src, err := os.OpenFile(os.DevNull, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		if err != nil {
			return fmt.Errorf("日志组件异常")
		}
		logger.Out = bufio.NewWriter(src)
	}
	logger.SetLevel(level)
	logger.SetFormatter(&ConsoleFormatter{})
	pathMap := lfshook.PathMap{
		logrus.InfoLevel:  loggerDir + "info.log",
		logrus.DebugLevel: loggerDir + "debug.log",
		logrus.ErrorLevel: loggerDir + "error.log",
	}
	hook := lfshook.NewHook(pathMap, &FileFormatter{})
	logger.AddHook(hook)
	return nil
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
