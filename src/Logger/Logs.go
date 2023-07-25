package cpu

import (
	"io"
	"log"
	"os"
)

func DebugLogger() *log.Logger {
	// f, err := os.OpenFile("../logs/debug.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	f, err := os.OpenFile("./debug.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	w := io.MultiWriter(f)
	debugLog := log.New(w, "", 0)
	return debugLog
}

func StateLogger() *log.Logger {
	// f, err := os.OpenFile("../logs/state.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	// if err != nil {
	// 	panic(err)
	// }
	// w := io.MultiWriter(f)
	w := io.MultiWriter(os.Stdout)
	stateLog := log.New(w, "", 0)
	return stateLog
}
