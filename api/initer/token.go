package initer

import "aidanwoods.dev/go-paseto"

var PasetoKey paseto.V4SymmetricKey

func InitKeyToken() {
	PasetoKey = paseto.NewV4SymmetricKey()
}
