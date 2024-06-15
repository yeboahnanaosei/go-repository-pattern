package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	app, err := setup()
	if err != nil {
		log.Fatalln("failed to setup server:", err)
	}

	server := &http.Server{
		Addr:        ":8080",
		Handler:     app.router,
		ReadTimeout: time.Second * 5,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatalln("server failed:", err)
	}
}
