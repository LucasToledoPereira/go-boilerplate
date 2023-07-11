package main

import (
	gbp "github.com/LucasToledoPereira/go-boilerplate"
	gameapi "github.com/LucasToledoPereira/go-boilerplate/apis/game"
)

// @title Go Boilerplate API Example
// @version 1.0
// @description
// @host localhost:8080
// @BasePath /api/v1
func main() {
	builder := gbp.Default()
	gameapi.LoadGameAPI(builder)
	server, _ := builder.Bootstrap()
	server.Run()
}
