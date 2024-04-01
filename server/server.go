package server

import (
	"log"
	"net/http"
)

func runServer() error {
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatalln(err)
		return err
	}
	return nil
}
