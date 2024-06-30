package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/ELRAS1/shortlink/internal/server"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	srv := server.StartServer(ctx)
	go func() {
		if err := srv.Srv.ListenAndServe(); err != nil {
			log.Fatalln(err)
		}
	}()
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig
	cancel()
	if err := server.Finish(ctx, srv); err != nil && err != http.ErrServerClosed {
		log.Println(err)
	}
}
