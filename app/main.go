package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/Serioga111/CutterService/internal/handler"
	storage "github.com/Serioga111/CutterService/internal/repositorie"
)

func main() {

	var storageMod string
	flag.StringVar(&storageMod, "m", "psql", "Which type of storage should we use")

	var s storage.Storage

	flag.Parse()
	switch storageMod {
	case "im":
		s = storage.NewMemoryStorage()
	default:
		s = storage.NewPostgresStorage()
	}

	h := handler.NewHandler(s)
	mux := http.NewServeMux()

	h.RegisterRoutes(mux)

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
