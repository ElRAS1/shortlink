package main

import (
	"log"
	"shortlink/server"
)

func main() {

	err := runServer()

	if err != nil {
		log.Fatalln(err)
	}
}
