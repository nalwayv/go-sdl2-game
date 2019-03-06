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
	logIt *Logger
	once  sync.Once
)

// GetInstance ...
func GetInstance(fileName string) *Logger {
	once.Do(func() {
		logIt = createLogger(fileName)
	})
	return logIt
}

// SLogger ...
var SLogger = GetInstance("src/gologger/golog.log")

func createLogger(fileName string) *Logger {
	file, _ := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)

	return &Logger{
		filename: fileName,
		Logger:   log.New(file, "SDL >> ", log.Lshortfile|log.Ltime),
	}
}
