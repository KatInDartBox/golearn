package utils

import (
	"api-test/utils/setDefault"
	"log"
)

func LogError(err error, msgs ...string) {
	var errMsg string
	if len(msgs) <=1{
		setDefault.Str(&errMsg,"log Error!")
	}

	 
	
	 
	if err != nil {
		log.Fatal(errMsg)
	}
	if len(msgs) == 2 {

		log.Println(msgs[1])
	}
	
}