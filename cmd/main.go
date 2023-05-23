/*
Author: Satria Bagus(satria.bagus18@gmail.com)
main.go (c) 2023
Desc: description
Created:  2023-05-22T08:27:05.292Z
Modified: !date!
*/

package main

import (
	"log"

	"github.com/satriabagusi/campo-sport/config"
	"github.com/satriabagusi/campo-sport/pkg/server"
)

func main() {
	cfg := config.NewConfig()

	//Load configuration values
	cfg.Load()

	//create a new instance of the server
	srv := server.NewServer()

	//initialize server with the configuration values
	err := srv.Initialize(cfg.PostgresConnectionString)
	if err != nil {
		log.Fatalf("failed to initialize server %v", err)
	}

	err = srv.Start(cfg.ServerAddress)
	if err != nil {
		log.Fatalf("Failed to start the server: %v", err)
	}
}
