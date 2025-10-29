package main

import (
	"log"

	"github.com/IvanRoussev/iroussev/beat-my-cluster/internal/api"
	"github.com/IvanRoussev/iroussev/beat-my-cluster/internal/db"
	"github.com/IvanRoussev/iroussev/beat-my-cluster/internal/game"
	"github.com/IvanRoussev/iroussev/beat-my-cluster/internal/util"

	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load configurations: ", err)
	}

	dbconn, err := db.DBConnect(config.ConnectionString)
	if err != nil {
		log.Fatal("Cannot connect to Database: ", err)
	}

	err = dbconn.AutoMigrate(&db.Player{}, &db.Attack{})
	if err != nil {
		log.Fatal("Failed to Migrate Database: ", err)
	}

	g := game.NewGame()

	server := &api.Server{
		DB:   dbconn,
		Game: g,
	}
	server.SetupRouter()

	log.Println("Starting server on", config.ServerAddress)
	if err := server.Start(config.ServerAddress); err != nil {
		log.Fatal("Cannot start server: ", err)
	}
}
