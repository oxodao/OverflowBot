package log

import (
	"fmt"
	"runtime"
	"time"
)

func log(level string, str interface{}) {
	pc, fn, line, _ := runtime.Caller(2)
	fmt.Printf("[%v] %v: (%v:%v:%v) %v \n", level, time.Now().Format("2006/01/02 - 15:04:00"), runtime.FuncForPC(pc), fn, line, str)
}

func Error(err interface{}) {
	log("ERROR", err)
}

func Warn(err interface{}) {
	log("WARN", err)
}

func Normal(err interface{}) {
	log("NORMAL", err)
}

func Info(err interface{}) {
	log("INFO", err)
}
