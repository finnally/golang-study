package logmod

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"strings"
	"time"
)

type ConsoleLogger struct {
	level    LevelStruct
	isFile   bool
	filePath string
	fileName string
	fileObj  *os.File
}

func getInfo(skip int) (funcName, fileName string, lineNo int) {
	pc, file, lineNo, ok := runtime.Caller(skip)
	if !ok {
		fmt.Print("runtime.Caller() failed")
		return
	}
	funcName = runtime.FuncForPC(pc).Name()
	fileName = path.Base(file)
	funcName = strings.Split(funcName, ".")[1]
	return
}

func NewConsoleLogger(level string) *ConsoleLogger {
	level = strings.ToUpper(level)
	Level, err := ParseLevelString(level)
	if err != nil {
		fmt.Println("parse log level error:", err)
	}
	return &ConsoleLogger{
		level:  LevelStruct{Level},
		isFile: false,
	}
}

func (c *ConsoleLogger) isTrue(l LevelStruct, n LevelNumber) bool {
	return l.LevelName >= n
}

func (c *ConsoleLogger) messageFormat(ln LevelNumber, message string, rotateFlag bool) {
	now := time.Now()
	funcName, fileName, lineNo := getInfo(3)
	levelString, _ := ParseLevelInt(ln)
	msg := fmt.Sprintf(
		"[%s] [%v] [%s:%s:%d] %s\n",
		now.Format("2006-01-02 15:04:05"),
		levelString,
		fileName,
		funcName,
		lineNo,
		message,
	)
	if c.isTrue(c.level, ln) {
		fmt.Printf("%s", msg)
		if c.isFile {
			if c.checkTime() && rotateFlag {
				c.MoveFile()
			}
			_, err := c.fileObj.WriteString(msg)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}

func (c *ConsoleLogger) Debug(message string, rotateFlag bool) {
	c.messageFormat(DEBUG, message, rotateFlag)
}

func (c *ConsoleLogger) Info(message string, rotateFlag bool) {
	c.messageFormat(INFO, message, rotateFlag)
}
