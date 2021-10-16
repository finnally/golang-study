package loglevel

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

// 指定类型别名
type LevelNumber = uint16

// 定义日志级别常量
const (
	UNKNOWN LevelNumber = iota
	DEBUG
	INFO
	WARNNING
	ERROR
)

// 定义日志级别结构体
type LevelStruct struct {
	LevelName LevelNumber
}

func ParseLevelInt(levelNum LevelNumber) (string, error) {
	switch levelNum {
	case DEBUG:
		return "debug", nil
	case INFO:
		return "info", nil
	case WARNNING:
		return "warnning", nil
	case ERROR:
		return "error", nil
	default:
		return "unknown", errors.New("unknown log level")
	}
}

// 字符串转换为uint16类型：接收字符串，返回LevelNum数据类型
func ParseLevelString(levelString string) (LevelNumber, error) {
	levelString = strings.ToLower(levelString)
	switch levelString {
	case "debug":
		return DEBUG, nil
	case "info":
		return INFO, nil
	case "warnning":
		return WARNNING, nil
	case "error":
		return ERROR, nil
	default:
		return UNKNOWN, errors.New("unknown log level")
	}
}

// 通过uint16类型比较日志级别
func Enable(l LevelStruct, n LevelNumber) bool {
	return l.LevelName <= n
}

// 定义日志级别方法
func (l LevelStruct) Debug(msg string) {
	if Enable(l, DEBUG) {
		now := time.Now().Format("2016-01-02 15:04:05")
		fmt.Printf("[%s] %s\n", now, msg)
	}
}

func (l LevelStruct) Info(msg string) {
	if Enable(l, INFO) {
		now := time.Now().Format("2016-01-02 15:04:05")
		fmt.Printf("[%s] %s\n", now, msg)
	}
}

func (l LevelStruct) Warnning(msg string) {
	if Enable(l, WARNNING) {
		now := time.Now().Format("2016-01-02 15:04:05")
		fmt.Printf("[%s] %s\n", now, msg)
	}
}

func (l LevelStruct) Error(msg string) {
	if Enable(l, ERROR) {
		now := time.Now().Format("2016-01-02 15:04:05")
		fmt.Printf("[%s] %s\n", now, msg)
	}
}
