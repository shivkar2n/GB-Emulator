package Logger

import (
	"io"
	"log"
	"os"
)

func DebugLogger() *log.Logger {
	// f, err := os.OpenFile("../logs/debug.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	// if err != nil {
	// 	panic(err)
	// }
	w := io.MultiWriter(os.Stdout)
	// w := io.MultiWriter(os.Stdout, f)
	logger := log.New(w, "", 0)
	return logger
}

func StateLogger() *log.Logger {
	f, err := os.OpenFile("../logs/state.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	// w := io.MultiWriter(os.Stdout)
	w := io.MultiWriter(f)
	logger := log.New(w, "", 0)
	return logger
}
