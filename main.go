package main

import (
	i "ShelterChatBackend/Api/internal"
	"ShelterChatBackend/Api/database"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	database.SetupDatabase()
	handler := i.NewRouter()

	port := os.Getenv("API_PORT")

	if port == "" {
		port = "8808"
		log.Printf("Port auf Default gesetzt: %s\n", port)
	}

	srv := &http.Server{
		Addr:         ":" + port,
		Handler:      handler,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}
	
	go func() {
		log.Println("ShelterChat Api gestartet")
		srv.ListenAndServe()
	}()

	c := make(chan os.Signal, 1)

	signal.Notify(c, os.Interrupt)
	<-c

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	srv.Shutdown(ctx)

}
