package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/lightster/finch/internal/pkg/web"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("No .env file detected")
	}

	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		log.Fatal("provide a database URL via DATABASE_URL env var")
	}
	port := os.Getenv("HTTP_PORT")
	if port == "" {
		port = "80"
	}

	config := &web.Config{
		DatabaseURL: databaseURL,
	}

	http.HandleFunc("/song/", web.ServeSong(config))
	//http.HandleFunc("/stream/", web.ServeStream(config))

	err = http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
