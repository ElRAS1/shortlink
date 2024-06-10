package main

import (
	"log"

	"github.com/ELRAS1/shortlink/internal/server"
)

func main() {
	srv := server.StartServer()
	err := srv.ListenAndServe()

	if err != nil {
		log.Fatalln(err)
	}
}
