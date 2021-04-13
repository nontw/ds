package main

import (
	"log"

	"github.com/thanonchai/ds/internal/server"
)

func main() {
	srv := server.NewHTTPServer(":8080")
	log.Fatal(srv.ListenAndServe())
}
