package main

import (
	"log"
	"log/slog"

	flog "github.com/pjsoftware/go-file-logger"
)

func main() {
	testMixedLogOutput()
	testMixedSLogJSONOutput()
	testMixedSLogTextOutput()
}

func testMixedLogOutput() {
	logFile := flog.LogWriter("test-mixed-log", "")
	defer logFile.Close()
	flog.UseLog(logFile)

	log.Print("Log message via log.Print()")
	slog.Debug("Log message via slog.Debug()")
	slog.Info("Log message via slog.Info()")
}

func testMixedSLogJSONOutput() {
	logFile := flog.LogWriter("test-mixed-slogJSON", "")
	defer logFile.Close()

	opts := &slog.HandlerOptions{Level: slog.LevelDebug, AddSource: true}
	flog.UseSLogJSON(logFile, opts)

	log.Print("Log message via log.Print()")
	slog.Debug("Log message via slog.Debug()")
	slog.Info("Log message via slog.Info()")
}

func testMixedSLogTextOutput() {
	logFile := flog.LogWriter("test-mixed-slogText", "")
	defer logFile.Close()

	lvl := new(slog.LevelVar)
	lvl.Set(slog.LevelInfo)
	opts := &slog.HandlerOptions{AddSource: true, Level: lvl}
	flog.UseSLogText(logFile, opts)

	log.Print("Log message 1 via log.Print()")
	slog.Debug("Log message 1 via slog.Debug() -- not logged")
	slog.Info("Log message 1 via slog.Info()")

	lvl.Set(slog.LevelDebug)
	log.Print("Log message 2 via log.Print()")
	slog.Debug("Log message 2 via slog.Debug() -- should be logged")
	slog.Info("Log message 2 via slog.Info()")
}
