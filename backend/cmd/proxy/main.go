package main

import "github.com/bayu-aditya/ideagate/backend/internal/proxy/app"

func main() {
	configFileName := "./config.yaml"
	app.NewServer(configFileName)
}
