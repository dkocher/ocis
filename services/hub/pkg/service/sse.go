package service

import (
	"fmt"
	"net/http"
	"time"

	"github.com/cs3org/reva/v2/pkg/ctx"
	"github.com/cs3org/reva/v2/pkg/events"
	"github.com/go-chi/chi/v5"
	"github.com/r3labs/sse/v2"
)

// ServeSSE provides server sent events functionality
func ServeSSE(evts <-chan interface{}) func(chi.Router) {
	server := sse.New()
	go func() {
		for e := range evts {
			switch ev := e.(type) {
			case events.UploadReady:
				u := ev.ExecutingUser
				t := time.Now()
				server.Publish(u.Id.OpaqueId, &sse.Event{
					Data: []byte(fmt.Sprintf("[%s] Hello %s, your file %s is ready to work with", t.Format("2006-01-02 15:04:05"), u.Username, ev.Filename)),
				})
			}
		}
	}()

	return func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			u, ok := ctx.ContextGetUser(r.Context())
			if !ok {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			uid := u.GetId().GetOpaqueId()
			if uid == "" {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			stream := server.CreateStream(uid)
			stream.AutoReplay = false

			// add stream to URL
			q := r.URL.Query()
			q.Set("stream", uid)
			r.URL.RawQuery = q.Encode()

			server.ServeHTTP(w, r)
		})
	}

}
