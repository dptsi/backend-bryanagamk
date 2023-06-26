package main

import "log"

func main() {
	container := container.Setup()

	if err := server.InitGRPC(container); err != nil {
		log.Fatal(err)
	}
}
