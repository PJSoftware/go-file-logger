package flog

import (
	"fmt"
	"log"
	"log/slog"
	"os"
	"time"
)

const defaultFolder = "C:/LogFiles"

func rootFolder() string {
	folder := os.Getenv("FILE_LOGGER_FOLDER")
	if folder != "" {
		return folder
	}

	return defaultFolder
}

func LogWriter(app string) *os.File {
	now := time.Now()
	folder := fmt.Sprintf("%s/%4d/%02d/%02d", rootFolder(), now.Year(), now.Month(), now.Day())
	fileName := fmt.Sprintf("%4d%02d%02d-%s.log", now.Year(), now.Month(), now.Day(), app)
	os.MkdirAll(folder, 0777)

	file, err := os.OpenFile(folder+"/"+fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return nil
	}

	return file
}

func UseLog(logFile *os.File) {
	defaultLogger := log.Default()
	defaultLogger.SetOutput(logFile)
}

func UseSLogText(logFile *os.File, opts *slog.HandlerOptions) {
	l := slog.New(slog.NewTextHandler(logFile, opts))
	slog.SetDefault(l)
}

func UseSLogJSON(logFile *os.File, opts *slog.HandlerOptions) {
	l := slog.New(slog.NewJSONHandler(logFile, opts))
	slog.SetDefault(l)
}
