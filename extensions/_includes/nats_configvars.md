## Environment Variables

| Name | Type | Default Value | Description |
|------|------|---------------|-------------|
| OCIS_LOG_LEVEL<br/>NATS_LOG_LEVEL | string |  | |
| OCIS_LOG_PRETTY<br/>NATS_LOG_PRETTY | bool | false | |
| OCIS_LOG_COLOR<br/>NATS_LOG_COLOR | bool | false | |
| OCIS_LOG_FILE<br/>NATS_LOG_FILE | string |  | |
| NATS_DEBUG_ADDR | string | 127.0.0.1:9234 | Bind address of the debug server, where metrics, health, config and debug endpoints will be exposed.|
| NATS_DEBUG_TOKEN | string |  | Token to secure the metrics endpoint|
| NATS_DEBUG_PPROF | bool | false | Enables pprof, which can be used for profiling|
| NATS_DEBUG_ZPAGES | bool | false | Enables zpages, which can  be used for collecting and viewing traces in-me|
| NATS_NATS_HOST | string | 127.0.0.1 | |
| NATS_NATS_PORT | int | 9233 | |
| NATS_NATS_CLUSTER_ID | string | ocis-cluster | |
| NATS_NATS_STORE_DIR | string | ~/.ocis/nats | |