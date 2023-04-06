package uString

import "math/rand"

/*
a-z -> 97-122
A-Z -> 65-90
0-9 -> 48-57
*/

// this is the fastest amount all rnd
func Rnd(no uint8) string {
	b := make([]byte, no)

	var i uint8
	for i = 0; i < no; i++ {
		b[i] = byte(97 + rand.Intn(26))
	}
	return string(b)
}

func init() {
	rand.Seed(12345)
}

func Rnd3(no uint8) string {
	bytes := []byte{}
	var i uint8
	var b byte

	for i = 0; i < no; i++ {
		b = byte(97 + rand.Intn(26))
		bytes = append(bytes, b)
	}
	return string(bytes)
}

func Rnd2(no uint8) string {
	s := ""
	var i uint8
	var b byte
	for i = 0; i < no; i++ {
		b = byte(97 + rand.Intn(26))
		s = s + string(b)
	}
	return s
}

var alphabet = []byte("abcdefghijklmnopqrstuvwxyz")

func Rnd4(no uint8) string {
	b := make([]byte, no)

	var i uint8
	for i = 0; i < no; i++ {
		b[i] = alphabet[rand.Intn(26)]
	}
	return string(b)
}
