package main

import (
	"fmt"
	"net/http"
	"time"
)

const delaySeconds = 60

var isInitialize = false

func main() {
	http.HandleFunc("/health", func(w http.ResponseWriter, req *http.Request) {
		if !isInitialize {
			w.WriteHeader(503)
			w.Write([]byte("Not initialized application"))
			return
		}

		w.Write([]byte("OK"))
	})

	time.AfterFunc(delaySeconds*time.Second, func() {
		isInitialize = true
		fmt.Println("[INFO] Initialize successfully")
	})

	http.ListenAndServe(":8000", nil)
}
