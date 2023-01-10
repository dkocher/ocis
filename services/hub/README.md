# Hub Service

The hub service provides `server sent events` (`sse`) functionality to clients

## Subscribing

A client can use the `hub/sse` endpoint to subscribe to events. It will open an http connection for the server to send events.
These events can inform the client about different changes on the server, for example: file uploads, shares, space memberships, ...
See `Available Events` for details

## Available Events
For complete and up-to-date list of available events see `pkg/service/events.go`

For starters the `hub` service only serves the `UploadReady` event which is emitted when a files postprocessing is finished and the file is ready to work with
