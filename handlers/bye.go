package handlers

import (
	"fmt"
	"net/http"
	"time"
)

type Bye struct{

}

func NewBye() *Bye{
	return &Bye{}
}

func (b *Bye) ServeHTTP(rw http.ResponseWriter, r *http.Request){

	for i := 0; i < 100; i++ {
		time.Sleep(1000)
		fmt.Fprintln(rw, "Good Bye from the server")
	}
}