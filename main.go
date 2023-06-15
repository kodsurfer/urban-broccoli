package main

import (
	"fmt"
	"net/http"
	"os"
)

const defaultPort = "8080"

func main() {
	port := os.Args[1]
	if port == "" {
		port = defaultPort
	}

	serveMux := http.NewServeMux()

	serveMux.HandleFunc("/", http.HandlerFunc(handler))

	http.ListenAndServe(fmt.Sprintf(":%s", port), serveMux)
}

func handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPut:
		{
			message := r.URL.Query().Get("v")
			if message == "" {
				w.WriteHeader(400)
			}
		}

	}
}
