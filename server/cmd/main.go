package main

import (
	"log"

	"github.com/MahatVasudev/liveStreamingProject/server/cmd/api"
)

func main() {
	addr := ":5000"
	api1 := api.NewApiServer(addr, nil, nil)

	log.Printf("Server Starting at Address %s....", addr)

	if err := api1.Run(); err != nil {
		log.Panic("Server Error!!!")
	}

	log.Println("Server is offline....")
}
