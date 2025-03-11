package main

import "log"

func main() {
	cfg := config{
		addr: ":8080",
	}
	app := application{
		config: cfg,
	}

	// Creating Router
	router := app.getRouter()

	// Creating Server
	if err := app.run(router); err != nil {
		log.Fatal("Failed to create Server")
	}

}
