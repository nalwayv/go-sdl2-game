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

// getInstance ...
func getInstance(fileName string) *Logger {
	once.Do(func() {
		// logIt = createLogger(fileName)

		logIt = func(filename string) *Logger {
			file, _ := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)

			return &Logger{
				filename: fileName,
				Logger:   log.New(file, "SDL >> ", log.Lshortfile|log.Ltime),
			}
		}(fileName)

	})
	return logIt
}

// SLogger ...
var SLogger = getInstance("src/gologger/golog.log")
