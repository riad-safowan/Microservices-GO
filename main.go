package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(rw, "Welcome to Riad Safowan (Backend developer)") //write to http response
		log.Println("HomePage is called")                               // print on log with time
		d, err := ioutil.ReadAll(r.Body)
		if err != nil {
			// rw.WriteHeader(http.StatusBadRequest)
			// rw.Write([]byte("Oooops"))
			http.Error(rw, "Oooops", http.StatusBadRequest) // print the error msg with status code. And works equivalent to previos 2 lines
			return
		}

		log.Printf("Data: %s", d) // write the data came with http request
		fmt.Fprintf(rw, "Data: %s", d)
	})

	http.ListenAndServe(":9090", nil) // starts the server at mensioned port

}

//terminal
//rm -rf .git --- delete git
// go run main.go --- run the programme
// contrl + c --- close running programme
// curl -v -d 'amar sonar bangla'  localhost:9090 --- request with data to the server
