package main

import "mirror-apt/app"

func main() {
	server := app.NewInstance()
	server.Start()
}
