package main

import (
	"github.com/better-than-yours/video-thumbnail/api"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	server := api.Server{}

	server.Run(3000)
}
