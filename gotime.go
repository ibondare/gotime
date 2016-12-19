// A trivial microservice in Golang
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Timestamp struct {
	Time   string `json:"time"`
	Source string `json:"source"`
}

func main() {
	http.HandleFunc("/time/", timeHandler)
	http.ListenAndServe(":8080", nil)
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	m := Timestamp{time.Now().Format(time.RFC850), "System clock"}

	b, err := json.Marshal(m)

	if err != nil {
		panic(err)
	}

	w.Write(b)
}
