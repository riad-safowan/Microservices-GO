package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"

	"github.com/riad-safowan/GO_MICROSERVICES/pkg/routes"
)

func main() {
	l := log.New(os.Stdout, "rest-api", log.LstdFlags)

	r := mux.NewRouter() // Router witn mux

	routes.RegisterProductRoutes(r)
	http.Handle("/", r)

	s := &http.Server{ // custom server
		Addr:         ":9090",
		Handler:      r,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	go func() {
		err := s.ListenAndServe() // start custom server
		if err != nil {
			l.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	l.Println("Received terminated, graceful shutdown", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)
}

