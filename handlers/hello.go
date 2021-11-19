package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

//Hello is a simple handler
type Hello struct {
	l *log.Logger
}

// NewHello creates a new hallo handler with the given logger
func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

// ServeHttp implements the go http.Handler interface
func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "Welcome to Riad Safowan (Backend developer)") //write to http response
	h.l.Println("HomePage is called")                               // print on log with time

	// read the body
	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		// rw.WriteHeader(http.StatusBadRequest)
		// rw.Write([]byte("Oooops"))
		http.Error(rw, "Oooops", http.StatusBadRequest) // print the error msg with status code. And works equivalent to previos 2 lines
		return
	}

	h.l.Printf("Data: %s", d)      // write the data came with http request
	fmt.Fprintf(rw, "Data: %s", d) // write to the response
}
