package uJson

import (
	"encoding/json"
	"fmt"
	"os"
)

type user struct {
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Password string `json:"password"`
}

type userClient struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func FilterUser1() userClient {
	User := user{
		Name:     "jack",
		Age:      19,
		Password: "abc12345",
	}
	UserClient := userClient{
		Name: User.Name,
		Age:  User.Age,
	}
	return UserClient
}
func FilterUser2() userClient {
	User := user{
		Name:     "jack",
		Age:      19,
		Password: "abc12345",
	}
	//stringify
	jsonUser, err := json.Marshal(User)
	if err != nil {
		fmt.Println("marshal to json string err", err)
		os.Exit(1)
	}
	var UserClient userClient
	err = json.Unmarshal(jsonUser, &UserClient)
	if err != nil {
		fmt.Println("unmarshal json string err", err)
		os.Exit(1)
	}

	return UserClient
}
