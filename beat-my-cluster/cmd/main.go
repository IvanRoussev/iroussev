package main

import (
	"log"

	"github.com/IvanRoussev/iroussev/beat-my-cluster/internal/api"
	"github.com/IvanRoussev/iroussev/beat-my-cluster/internal/util"
	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")

	if err != nil {
		log.Fatal("Cannot load configurations: ", err)
	}

	server, err := api.NewServer(config)
	if err != nil {
		log.Fatal("cannot create server", err)
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("Cannot start server: ", err)
	}
}
