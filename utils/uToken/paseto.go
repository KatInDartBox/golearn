package uToken

import (
	"time"

	"aidanwoods.dev/go-paseto"
)

const preFix = "v4.local."

func Encrypt(key paseto.V4SymmetricKey, payload Payload, life time.Duration) (string, error) {
	token := paseto.NewToken()
	// token.SetIssuedAt(time.Now())
	token.SetExpiration(time.Now().Add(life))
	strPayload, err := GetPayloadString(payload)
	if err != nil {
		return "", err
	}
	token.SetString("payload", strPayload)
	tokenStr := token.V4Encrypt(key, nil)
	strStripPrefix := tokenStr[len(preFix):]
	return strStripPrefix, nil
}

func Decrypt(key paseto.V4SymmetricKey, encryption string) (Payload, error) {
	// add prefix to string
	encrypted := preFix + encryption
	token, err := paseto.NewParser().ParseV4Local(key, encrypted, nil)
	if err != nil {
		return Payload{}, err
	}
	payload, err := token.GetString("payload")
	if err != nil {
		return Payload{}, err
	}
	return ParseJson(payload)
}
