package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

const defaultPort = "8080"

var mq = make(map[string][]string, 0)

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
			if message != "" {
				queue := r.URL.Path
				if _, ok := mq[queue]; !ok {
					mq[queue] = make([]string, 0)
					mq[queue] = append(mq[queue], message)
				}
			} else {
				w.WriteHeader(400)
			}
		}
	case http.MethodGet:
		queue := r.URL.Path

		if len(mq) > 0 {
			message := mq[queue][0]
			mq[queue] = mq[queue][1:]

			w.Header().Set("Content-Type", "application/json")

			bytes, err := json.Marshal(message)
			if err != nil {
				log.Fatalln(err)
			}

			_, err = w.Write(bytes)
			if err != nil {
				log.Fatalln(err)
			}
		} else {
			w.WriteHeader(404)
		}
	}
}
