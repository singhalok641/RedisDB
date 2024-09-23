package main

import (
	"diceDb/config"
	"diceDb/server"
	"flag"
	"log"
)

func setupFlags() {
	flag.StringVar(&config.Host, "host", "0.0.0.0", "host for the dice server")
	flag.IntVar(&config.Port, "port", 7379, "port for the dice server")
	flag.Parse()
}

func main() {
	setupFlags()
	log.Println("Starting dice server")
	server.RunSyncTCPServer()
}
