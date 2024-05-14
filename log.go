package log

import (
	"fmt"
	clog "log"
	"os"
	"time"
)

type Log struct {
	open bool
	root string
	app  string
	lib  string
	file *os.File
}

const WindowsRootFolder = "C:/LogFiles"

var logInstance Log

func getLogInstance() *Log {
	if logInstance.open {
		return &logInstance
	}
	panic("log must be initialised via InitLog()")
}

func InitLog(root string, app string) {
	log := &logInstance

	log.root = root
	log.app = app
	log.lib = ""

	now := time.Now()
	folder := fmt.Sprintf("%s/%4d/%02d/%02d", root, now.Year(), now.Month(), now.Day())
	fileName := fmt.Sprintf("%4d%02d%02d-%s.log", now.Year(), now.Month(), now.Day(), log.app)
	os.MkdirAll(folder, 0777)

	file, err := os.OpenFile(folder+"/"+fileName, os.O_RDWR | os.O_CREATE | os.O_APPEND, 0644)
	if err != nil {
		return
	}

	log.file = file
	clog.SetOutput(log.file)
	clog.SetFlags(clog.Ldate | clog.Ltime)
	log.open = true

	Printf("Initialising '%s' log file", log.app)
}

func Close() {
	log := getLogInstance()
	if log.open {
		log.file.Close()
		log.open = false
	}
}

func Lib(lib string) *Log {
	log := getLogInstance()
	log.lib = lib
	return log
}

func (log *Log) Printf(format string, a ...any) {
	Printf(format, a...)
}

func Printf(format string, a ...any) {
	log := getLogInstance()
	msg := fmt.Sprintf(format, a...)
	prefix := log.prefix()
	clog.Printf("[%s] %s", prefix, msg)
}

func (log *Log) Println(v ...any) {
	Println(v...)
}

func Println(v ...any) {
	log := getLogInstance()
	msg := concat(v...)
	prefix := log.prefix()
	clog.Printf("[%s] %s", prefix, msg)
}

func (log *Log) Fatalf(format string, a ...any) {
	Fatalf(format, a...)
}

func Fatalf(format string, a ...any) {
	log := getLogInstance()
	msg := fmt.Sprintf(format, a...)
	prefix := log.prefix()
	clog.Fatalf("[%s] %s", prefix, msg)
}

func (log *Log) Panic(v ...any) {
	Panic(v...)
}

func Panic(v ...any) {
	log := getLogInstance()
	msg := concat(v...)
	prefix := log.prefix()
	clog.Panicf("[%s] %s", prefix, msg)
}

func concat(v ...any) string {
	msg := ""
	for _, val := range v {
		if msg != "" {
			msg += " "
			msg += fmt.Sprintf("%v", val)
		}
	}
	return msg
}

func (log *Log) prefix() string {
	prefix := log.app
	if log.lib != "" {
		prefix += ":" + log.lib
		log.lib = ""
	}
	return prefix
}
