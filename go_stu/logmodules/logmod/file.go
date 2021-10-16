package logmod

import (
	"fmt"
	"os"
	"path"
	"time"
)

func (c *ConsoleLogger) FileInit(isFile bool, filePath string, fileName string) error {
	c.isFile = isFile
	c.filePath = filePath
	c.fileName = fileName
	logPath := path.Join(filePath, fileName)
	fileObj, err := os.OpenFile(logPath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("open log failed, error:", err)
		return err
	}
	c.fileObj = fileObj
	return nil
}

func (c *ConsoleLogger) checkTime() bool {
	return time.Now().Format("05") == "00"
}

func (c *ConsoleLogger) MoveFile() {
	nowStr := time.Now().Add(-time.Minute * 1).Format("2006-01-02-1504")
	logPath := path.Join(c.filePath, c.fileName)
	newLogPath := logPath + nowStr
	c.fileObj.Close()
	err := os.Rename(logPath, newLogPath)
	if err != nil {
		fmt.Println(err)
	}

	fileObj, err := os.OpenFile(logPath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("open new logfile error: ", err)
		return
	}
	c.fileObj = fileObj
	time.Sleep(time.Second)
}
