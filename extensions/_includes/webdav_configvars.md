## Environment Variables

| Name | Type | Default Value | Description |
|------|------|---------------|-------------|
| OCIS_TRACING_ENABLED<br/>WEBDAV_TRACING_ENABLED | bool | false | Enable tracing.|
| OCIS_TRACING_TYPE<br/>WEBDAV_TRACING_TYPE | string |  | The tracing type.|
| OCIS_TRACING_ENDPOINT<br/>WEBDAV_TRACING_ENDPOINT | string |  | The tracing service endpoint.|
| OCIS_TRACING_COLLECTOR<br/>WEBDAV_TRACING_COLLECTOR | string |  | The tracing collector.|
| OCIS_LOG_LEVEL<br/>WEBDAV_LOG_LEVEL | string |  | The log level.|
| OCIS_LOG_PRETTY<br/>WEBDAV_LOG_PRETTY | bool | false | Enable pretty log output.|
| OCIS_LOG_COLOR<br/>WEBDAV_LOG_COLOR | bool | false | Enable colored log output.|
| OCIS_LOG_FILE<br/>WEBDAV_LOG_FILE | string |  | The path to the file if the log should write to file.|
| WEBDAV_DEBUG_ADDR | string | 127.0.0.1:9119 | |
| WEBDAV_DEBUG_TOKEN | string |  | |
| WEBDAV_DEBUG_PPROF | bool | false | |
| WEBDAV_DEBUG_ZPAGES | bool | false | |
| WEBDAV_HTTP_ADDR | string | 127.0.0.1:9115 | The HTTP API address.|
| WEBDAV_HTTP_ROOT | string | / | The HTTP API root path.|
| OCIS_URL<br/>OCIS_PUBLIC_URL | string | https://127.0.0.1:9200 | |
| WEBDAV_WEBDAV_NAMESPACE | string | /users/{{.Id.OpaqueId}} | CS3 path layout to use when forwarding /webdav requests|
| REVA_GATEWAY | string | 127.0.0.1:9142 | |