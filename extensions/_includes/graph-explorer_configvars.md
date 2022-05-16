## Environment Variables

| Name | Type | Default Value | Description |
|------|------|---------------|-------------|
| OCIS_TRACING_ENABLED<br/>GRAPH_EXPLORER_TRACING_ENABLED | bool | false | Enable tracing.|
| OCIS_TRACING_TYPE<br/>GRAPH_EXPLORER_TRACING_TYPE | string |  | The sampler type: remote, const, probabilistic, ratelimiting (default remote). See also https://www.jaegertracing.io/docs/latest/sampling/.|
| OCIS_TRACING_ENDPOINT<br/>GRAPH_EXPLORER_TRACING_ENDPOINT | string |  | The endpoint of the tracing service.|
| OCIS_TRACING_COLLECTOR<br/>GRAPH_EXPLORER_TRACING_COLLECTOR | string |  | The HTTP endpoint for sending spans directly to a collector, i.e. http://jaeger-collector:14268/api/traces. If specified, the tracing endpoint is ignored.|
| OCIS_LOG_LEVEL<br/>GRAPH_EXPLORER_LOG_LEVEL | string |  | The log level.|
| OCIS_LOG_PRETTY<br/>GRAPH_EXPLORER_LOG_PRETTY | bool | false | Enable pretty logs.|
| OCIS_LOG_COLOR<br/>GRAPH_EXPLORER_LOG_COLOR | bool | false | Enable colored logs.|
| OCIS_LOG_FILE<br/>GRAPH_EXPLORER_LOG_FILE | string |  | The path to the log file when logging to file.|
| GRAPH_EXPLORER_DEBUG_ADDR | string | 127.0.0.1:9136 | |
| GRAPH_EXPLORER_DEBUG_TOKEN | string |  | |
| GRAPH_EXPLORER_DEBUG_PPROF | bool | false | |
| GRAPH_EXPLORER_DEBUG_ZPAGES | bool | false | |
| GRAPH_EXPLORER_HTTP_ADDR | string | 127.0.0.1:9135 | The HTTP service address.|
| GRAPH_EXPLORER_HTTP_ROOT | string | /graph-explorer | The HTTP service root path.|
| GRAPH_EXPLORER_CLIENT_ID | string | ocis-explorer.js | |
| OCIS_URL<br/>OCIS_OIDC_ISSUER<br/>GRAPH_EXPLORER_ISSUER | string | https://localhost:9200 | |
| OCIS_URL<br/>GRAPH_EXPLORER_GRAPH_URL_BASE | string | https://localhost:9200 | |
| GRAPH_EXPLORER_GRAPH_URL_PATH | string | /graph | |