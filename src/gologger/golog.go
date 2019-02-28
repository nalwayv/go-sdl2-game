package gologger

/*
	Singleton
*/

import (
	"log"
	"os"
	"sync"
)

// Logger ...
type Logger struct {
	filename string
	*log.Logger
}

var (
	logit *Logger
	once  sync.Once
)

// GetInstance ...
func GetInstance(fname string) *Logger {
	once.Do(func() {
		logit = createLogger(fname)
	})
	return logit
}

// SLogger ...
var SLogger = GetInstance("src/gologger/golog.log")

func createLogger(fname string) *Logger {
	file, _ := os.OpenFile(fname, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)

	return &Logger{
		filename: fname,
		Logger:   log.New(file, "SDL >> ", log.Lshortfile|log.Ltime),
	}
}
