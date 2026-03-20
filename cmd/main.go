package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Serioga111/CutterService/internal/handler"
	"github.com/Serioga111/CutterService/internal/repositorie"
)

func main() {

	var storageMod string
	flag.StringVar(&storageMod, "m", "psql", "Which type of storage should we use")

	var s repositorie.Repositorie

	flag.Parse()
	switch storageMod {
	case "im":
		s = repositorie.NewInMemoryRepositorie()
	case "psql":
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
			os.Getenv("DB_HOST"),
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_NAME"),
			os.Getenv("DB_PORT"),
			os.Getenv("SSL_MODE"),
		)
		var err error
		s, err = repositorie.NewPostgresRepositorie(dsn)
		if err != nil {
			log.Fatalf("Failed to create PostgreSQl repositorie: %v", err)
		}
	default:
		log.Fatal("Invalid storage mode:", storageMod)
	}

	h := handler.NewHandler(s)
	mux := http.NewServeMux()

	h.RegisterRoutes(mux)

	log.Printf("Server started on :8080\nStorage mode: %v", storageMod)
	log.Fatal(http.ListenAndServe(":8080", mux))
}
