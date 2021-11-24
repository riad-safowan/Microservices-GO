package handlers

import (
	"fmt"
	"net/http"
	"time"
)

type Bye struct {
}

func NewBye() *Bye {
	return &Bye{}
}

func (b *Bye) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

	for i := 0; i < 4; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Fprintln(rw, "Good Bye from the server")
	}
}