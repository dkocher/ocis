package service

import (
	"fmt"
	"net/http"
	"time"

	"github.com/cs3org/reva/v2/pkg/ctx"
	"github.com/go-chi/chi/v5"
	"github.com/r3labs/sse/v2"
)

// ServeSSE provides server sent events functionality
func ServeSSE(r chi.Router) {
	server := sse.New()
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

		go func() {
			for range time.Tick(time.Second * 4) {
				t := time.Now()
				server.Publish(uid, &sse.Event{
					Data: []byte(fmt.Sprintf("[%s] Hello %s, new push notification from server!", t.Format("2006-01-02 15:04:05"), u.Username)),
				})
			}
		}()

		// add stream to URL
		q := r.URL.Query()
		q.Set("stream", uid)
		r.URL.RawQuery = q.Encode()

		server.ServeHTTP(w, r)
	})

}
