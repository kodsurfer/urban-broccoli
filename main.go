package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	timeoutArg := os.Args[1]
	fmt.Println("Timeout arg: ", timeoutArg)

	serveMux := http.NewServeMux()

	serveMux.HandleFunc("/", http.HandlerFunc(handler))

	http.ListenAndServe(":8080", serveMux)
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
