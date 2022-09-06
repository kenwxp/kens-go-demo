package enty_logger

import (
	"bytes"
	"fmt"
	"kens/demo/util"
	"os"
	"runtime"
	"strings"
)

type loggerLevel int

const (
	LogInfo  = loggerLevel(1)
	LogDebug = loggerLevel(2)
	LogTrace = loggerLevel(3)
)

var setLevel int

func Info(a ...interface{}) {
	print(LogInfo, "INFO", a)
}
func Debug(a ...interface{}) {
	print(LogDebug, "DEBUG", a)
}
func Trace(a ...interface{}) {
	print(LogTrace, "TRACE", a)
}

func print(level loggerLevel, levelStr string, a []interface{}) {
	if checkLoggerLevel(level) {
		arr := append(make([]interface{}, 0), "["+util.TimeNowFormatString()+"]")
		arr = append(arr, levelStr)
		arr = append(arr, "["+GetGid()+"]")
		fmt.Println(append(arr, a...)...)
	}
}
func Init() {
	setLoggerLevel := os.Getenv("LOG_LEVEL")
	if strings.ToUpper(setLoggerLevel) == "INFO" {
		setLevel = 1
	} else if strings.ToUpper(setLoggerLevel) == "DEBUG" {
		setLevel = 2
	} else if strings.ToUpper(setLoggerLevel) == "TRACE" {
		setLevel = 3
	} else {
		setLevel = 1
	}
}

func checkLoggerLevel(level loggerLevel) bool {
	return int(level) <= setLevel
}

func GetGid() string {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	return string(b)
}
