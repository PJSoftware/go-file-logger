package flogger

import (
	"fmt"
	"os"
	"time"
)

const root = "C:/LogFiles"

func LogWriter(app string) *os.File {
	now := time.Now()
	folder := fmt.Sprintf("%s/%4d/%02d/%02d", root, now.Year(), now.Month(), now.Day())
	fileName := fmt.Sprintf("%4d%02d%02d-%s.log", now.Year(), now.Month(), now.Day(), app)
	os.MkdirAll(folder, 0777)

	file, err := os.OpenFile(folder+"/"+fileName, os.O_RDWR | os.O_CREATE | os.O_APPEND, 0644)
	if err != nil {
		return nil
	}

	return file
}
