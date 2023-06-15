package main

import (
	"net/http"
)

func main() {

	serveMux := http.NewServeMux()

	serveMux.HandleFunc("/", http.HandlerFunc(handler))

	http.ListenAndServe(":8080", serveMux)
}

func handler(w http.ResponseWriter, r *http.Request) {

}
