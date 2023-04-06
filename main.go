package main

import (
	"fmt"
	// "learn/utils/uToken"
	// "time"
	// "aidanwoods.dev/go-paseto"
)

func main() {
	str := "abc123456"

	strStrip := str[3:]

	fmt.Println(str)
	fmt.Println(strStrip)

}

func panicIf(err error) {
	if err != nil {
		panic(err)
	}
}
