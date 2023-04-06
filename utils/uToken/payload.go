package uToken

import "encoding/json"

type Payload struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

func GetPayloadString(p Payload) (string, error) {
	str, err := json.Marshal(p)
	if err != nil {
		return "", err
	}
	return string(str), nil
}
func ParseJson(str string) (Payload, error) {
	var payload Payload
	err := json.Unmarshal([]byte(str), &payload)
	if err != nil {
		return Payload{}, err
	}
	return payload, nil
}
