package main

import (
	"fmt"
	"time"

	"study.go.com/logmodules/logmod"
)

func main() {
	rotateFlag := true
	log := logmod.NewConsoleLogger("info")
	err := log.FileInit(true, "F:\\GoProject\\src\\study.go.com\\go_stu\\logmodules", "detail.log")
	if err != nil {
		fmt.Println("init log file failed, error:", err)
		return
	}
	for {
		log.Debug("This is a debug message", rotateFlag)
		log.Info("This is a info message", !rotateFlag)
		time.Sleep(time.Second)
	}
}
