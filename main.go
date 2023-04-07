package main

import (
	"fmt"
	"strings"
	// "learn/utils/uToken"
	// "time"
	// "aidanwoods.dev/go-paseto"
)

func main() {
	str := "abc 123456"

	fields := strings.Fields(str)

	error := fmt.Errorf("fields %v", fields)
	panicIf(error)

	fmt.Println("after error f")
}

func panicIf(err error) {
	if err != nil {
		panic(err)
	}
}
