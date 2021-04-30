package main

import (
	"github.com/Leonardo-Antonio/api-technology.online.store/src/server"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	if err := godotenv.Load(); err != nil { log.Fatalln(err) }
	app := server.NewApp()
	app.Middlewares()
	app.Routes()
	app.Listening()
}