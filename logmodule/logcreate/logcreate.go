package logcreate

import "study.go.com/logmodule/loglevel"

func NewLog(levelString string) loglevel.LevelStruct {
	level, err := loglevel.ParseLevelString(levelString)
	if err != nil {
		panic(err)
	}
	return loglevel.LevelStruct{LevelName: level}
}

// import (
// 	"study.go.com/logodule/logcreate"
// )

// log := logcreate.NwLog(error")
// log.Debug("debug)
// log.Info("info")
// log.Warnning("warnig")
// log.Error("error")
