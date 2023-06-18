package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

const defaultPort = "8080"

type MessageQueue struct {
	mq map[string][]string
}

func main() {
	port := defaultPort
	if len(os.Args) > 1 {
		port = os.Args[1]
	}

	serveMux := http.NewServeMux()

	serveMux.HandleFunc("/", handler)

	err := http.ListenAndServe(fmt.Sprintf("127.0.0.1:%s", port), serveMux)
	if err != nil {
		log.Fatalln(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	messageQueue := NewMessageQueue()

	switch r.Method {
	case http.MethodPut:
		{
			message := r.URL.Query().Get("v")
			if message != "" {
				queue := r.URL.Path
				messageQueue.push(queue, message)
			} else {
				w.WriteHeader(400)
			}
		}
	case http.MethodGet:
		queue := r.URL.Path

		if len(messageQueue.mq) > 0 {
			message := messageQueue.pop(queue)

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

func (mq *MessageQueue) pop(queue string) string {
	message := mq.mq[queue][0]
	mq.mq[queue] = mq.mq[queue][1:]

	return message
}

func (mq *MessageQueue) push(queue string, message string) {
	_, ok := mq.mq[queue]
	if ok {
		mq.mq[queue] = append(mq.mq[queue], message)
	} else {
		mq.mq[queue] = append(make([]string, 0), message)
	}
}

func NewMessageQueue() *MessageQueue {
	return &MessageQueue{
		mq: make(map[string][]string, 0),
	}
}
