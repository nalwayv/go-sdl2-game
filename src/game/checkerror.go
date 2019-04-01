package game

/*
Info
---
a simple check error function insted of writing
{{{ go
	if err != nil {

	}
}}}
all over the place
**/

import (
	"../gologger"
)

func checkError(err error) {
	if err != nil {
		gologger.SLogger.Fatalln(err)
	}
}
