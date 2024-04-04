package main

import (
	"github.com/mcisback/snitcher-events-server/socket"
)

const PORT = ":8219"

func main() {
	// Create a new Fiber instance
	// app := fiber.New()

	// h := handlers.New(app)
	// h.Init()

	// queue := queue.New()

	// queue.Add("Hello World", "data")

	// fmt.Printf("Server is running on port %s\n", PORT)

	// err := app.Listen(PORT)
	// if err != nil {
	// 	fmt.Printf("Error starting server: %v", err)
	// }

	server := socket.New()

	server.QManager.AddQ("JOB_QUEUE_1")
	server.QManager.AddQ("JOB_QUEUE_2")

	go server.Start(PORT)
}
