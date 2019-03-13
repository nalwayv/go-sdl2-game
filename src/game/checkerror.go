package game

import (
	"../gologger"
)

func checkError(err error) {
	if err != nil {
		gologger.SLogger.Fatalln(err)
	}
}
