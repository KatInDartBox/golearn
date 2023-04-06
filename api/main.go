package main

import (
	"learn/api/initer"
	"learn/api/server"
	"log"
)

func init() {
	initer.InitDB()
	initer.AddBookableDate()
	initer.InitKeyToken()
}

func main() {
	server := server.NewServer(initer.Qry)
	err := server.Start(":5000")
	if err != nil {
		log.Fatalf("can not connect to server %v", err)
	}
}
