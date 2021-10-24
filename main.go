package main

import (
	"context"
	"flayshon/micro/handlers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
	"github.com/nicholasjackson/env"
)

var bindAddress = env.String("BIND_ADDRESS", false, ":9090", "Bind address for the server")

func main() {
	env.Parse()

	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	hh := handlers.NewHello(l)
	ph := handlers.NewProducts(l)

	sm := http.NewServeMux()
	sm.Handle("/hello", hh)
	sm.Handle("/", ph)

	s := &http.Server{
		Addr: *bindAddress,
		Handler: sm,
		ErrorLog: l,
		ReadTimeout: 5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout: 120 * time.Second,
	}

	go func () {
		err := s.ListenAndServe()
		if err != nil {
			l.Printf("Error starting the server: %s\n", err)
			os.Exit(1)
		}
	}()

	sigChan := make(chan os.Signal, 1)

	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	l.Println("Received terminate. Graceful shutdown", sig)

	tc, _ := context.WithTimeout(context.Background(), 30 * time.Second)

	s.Shutdown(tc)

}