package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Hello struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "Welcome to Riad Safowan (Backend developer)") //write to http response
	h.l.Println("HomePage is called")                               // print on log with time
	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		// rw.WriteHeader(http.StatusBadRequest)
		// rw.Write([]byte("Oooops"))
		http.Error(rw, "Oooops", http.StatusBadRequest) // print the error msg with status code. And works equivalent to previos 2 lines
		return
	}

	h.l.Printf("Data: %s", d) // write the data came with http request
	fmt.Fprintf(rw, "Data: %s", d)
}
