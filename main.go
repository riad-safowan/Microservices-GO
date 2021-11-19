package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"p1/handlers"
	"time"
)

func main() {
	l := log.New(os.Stdout, "rest-api", log.LstdFlags)

	hh := handlers.NewHello(l)
	bh := handlers.NewBye()
	ph := handlers.NewProducts(l)

	sm := http.NewServeMux()
	sm.Handle("/", hh)
	sm.Handle("/bye", bh)
	sm.Handle("/products", ph)

	s := &http.Server{ //custom server
		Addr:         ":9090",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	// http.ListenAndServe(":9090", sm) // starts the server at mensioned port
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

//terminal
// go run main.go --- run the programme
// contrl + c --- close running programme
// curl localhost:9090 --- request to the server
// curl localhost:9090 -v -XDELETE --- specific request to the server
// curl -v localhost:9090 --- request with info
// curl -v -d 'amar sonar bangla'  localhost:9090 --- request with data to the server

//git
// git config user.name
// git config user.email
// git config --global user.email "mailAddress"
// git init --- initialize directory to git repository
// rm -rf .git --- delete git
// git status --- status with current branch
// git branch --list --- list of branch
// git branch newbranchname --- create branch locally
// git checkout branchname --- switch branch
// git checkout -b branchname --- create and switch
// git push --set-upstream origin branchname --- push with newly created branch
// git merge branchname --- merge 'branchname' branch to current branch
// git branch -d branchname --- delete local branch, -D for force delete
