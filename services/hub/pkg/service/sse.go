package service

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	gateway "github.com/cs3org/go-cs3apis/cs3/gateway/v1beta1"
	user "github.com/cs3org/go-cs3apis/cs3/identity/user/v1beta1"
	rpc "github.com/cs3org/go-cs3apis/cs3/rpc/v1beta1"
	provider "github.com/cs3org/go-cs3apis/cs3/storage/provider/v1beta1"
	"github.com/cs3org/reva/v2/pkg/ctx"
	revactx "github.com/cs3org/reva/v2/pkg/ctx"
	"github.com/cs3org/reva/v2/pkg/events"
	"github.com/cs3org/reva/v2/pkg/rgrpc/todo/pool"
	"github.com/cs3org/reva/v2/pkg/storagespace"
	"github.com/go-chi/chi/v5"
	"github.com/r3labs/sse/v2"
	"google.golang.org/grpc/metadata"
)

// ServeSSE provides server sent events functionality
func ServeSSE(evts <-chan interface{}) func(chi.Router) {
	server := sse.New()

	// TODO: configure properly
	gwc, err := pool.GetGatewayServiceClient("127.0.0.1:9142")
	if err != nil {
		log.Fatal(err)
	}

	// TODO: start multiple eventListeners?
	go eventListener(evts, server, gwc)

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

func eventListener(evts <-chan interface{}, server *sse.Server, gwc gateway.GatewayAPIClient) {
	for e := range evts {
		rcps, ev := extractDetails(e, gwc)
		publish(server, ev, rcps)
	}
}

func publish(server *sse.Server, event []byte, rcps <-chan string) {
	for r := range rcps {
		server.Publish(r, &sse.Event{Data: event})
	}

}

// extracts recipients and builds event to send to client
func extractDetails(e interface{}, gwc gateway.GatewayAPIClient) (<-chan string, []byte) {

	// determining recipients can take longer. We spawn a seperate go routine to do it
	ch := determineRecipients(e, gwc)

	var event interface{}

	switch ev := e.(type) {
	case events.UploadReady:

		// TODO: add timestamp to event
		t := time.Now().Format("2006-01-02 15:04:05")
		id, _ := storagespace.FormatReference(ev.FileRef)
		event = UploadReady{
			FileID:    id,
			Filename:  ev.Filename,
			Timestamp: t,
			Message:   fmt.Sprintf("[%s] Hello! The file %s is ready to work with", t, ev.Filename),
		}

	}

	b, err := json.Marshal(event)
	if err != nil {
		log.Println(err)
	}

	return ch, b
}

func determineRecipients(e interface{}, gwc gateway.GatewayAPIClient) <-chan string {
	ch := make(chan string)
	go func() {
		var (
			ref  *provider.Reference
			user *user.User
		)
		switch ev := e.(type) {
		case events.UploadReady:
			ref = ev.FileRef
			user = ev.ExecutingUser

		}

		// impersonate executing user to stat the resource
		// FIXME: What to do if executing user is not member of the space?
		iRes, err := impersonate(user.GetId(), gwc)
		if err != nil {
			fmt.Println("ERROR:", err)
			return
		}

		//resp, err := gwc.Stat(metadata.AppendToOutgoingContext(context.Background(), revactx.TokenHeader, iRes.Token), &provider.StatRequest{Ref: ref, ArbitraryMetadataKeys: []string{"*"}})
		filters := []*provider.ListStorageSpacesRequest_Filter{listStorageSpacesIDFilter(ref.GetResourceId().GetSpaceId())}
		resp, err := gwc.ListStorageSpaces(metadata.AppendToOutgoingContext(context.Background(), revactx.TokenHeader, iRes.Token), &provider.ListStorageSpacesRequest{Filters: filters})
		if err != nil || resp.GetStatus().GetCode() != rpc.Code_CODE_OK || len(resp.GetStorageSpaces()) != 1 {
			fmt.Println("ERROR:", err, resp.GetStatus().GetCode())
			return
		}

		informed := make(map[string]struct{})

		// inform executing user first and foremost
		ch <- user.Id.OpaqueId
		informed[user.GetId().GetOpaqueId()] = struct{}{}

		// inform space members next
		// TODO: create utils.ReadJSONFromOpaque
		if o := resp.GetStorageSpaces()[0].GetOpaque(); o != nil && o.Map != nil {
			var m map[string]*provider.ResourcePermissions
			entry, ok := o.Map["grants"]
			if ok {
				err := json.Unmarshal(entry.Value, &m)
				if err != nil {
					fmt.Println("ERROR:", err)
				}
			}

			// FIXME: Which space permissions allow me to get this event?
			for u := range m {
				// TODO: Resolve group recipients

				if _, ok := informed[u]; ok {
					continue
				}

				ch <- u
				informed[u] = struct{}{}
			}
		}

		// TODO: inform share recipients

		close(ch)
	}()
	return ch
}

func impersonate(userID *user.UserId, gwc gateway.GatewayAPIClient) (*gateway.AuthenticateResponse, error) {
	getUserResponse, err := gwc.GetUser(context.Background(), &user.GetUserRequest{
		UserId: userID,
	})
	if err != nil {
		return nil, err
	}
	if getUserResponse.Status.Code != rpc.Code_CODE_OK {
		return nil, fmt.Errorf("error getting user: %s", getUserResponse.Status.Message)
	}

	// Get auth context
	ownerCtx := revactx.ContextSetUser(context.Background(), getUserResponse.User)
	authRes, err := gwc.Authenticate(ownerCtx, &gateway.AuthenticateRequest{
		Type:     "machine",
		ClientId: "userid:" + userID.OpaqueId,
		// TODO: use machine_auth_api_key
		ClientSecret: "wPzAX0!E.lqkI1EXodvTgMwHDiuEa!a=",
	})
	if err != nil {
		return nil, err
	}
	if authRes.GetStatus().GetCode() != rpc.Code_CODE_OK {
		return nil, fmt.Errorf("error impersonating user: %s", authRes.Status.Message)
	}
	return authRes, nil
}

func listStorageSpacesIDFilter(id string) *provider.ListStorageSpacesRequest_Filter {
	return &provider.ListStorageSpacesRequest_Filter{
		Type: provider.ListStorageSpacesRequest_Filter_TYPE_ID,
		Term: &provider.ListStorageSpacesRequest_Filter_Id{
			Id: &provider.StorageSpaceId{
				OpaqueId: id,
			},
		},
	}
}
