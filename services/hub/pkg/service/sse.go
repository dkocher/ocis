package service

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/r3labs/sse/v2"
	"time"
)

type SSE struct{}

func ServeSSE(r chi.Router) {
	server := sse.New()
	stream := server.CreateStream("messages")
	stream.AutoReplay = false

	r.Get("/", server.ServeHTTP)

	go func() {
		for range time.Tick(time.Second * 4) {
			t := time.Now()
			server.Publish("messages", &sse.Event{
				Data: []byte(fmt.Sprintf("[%s] New push notification from server", t.Format("2006-01-02 15:04:05"))),
			})
		}
	}()
}
