## Environment Variables

| Name | Type | Default Value | Description |
|------|------|---------------|-------------|
| OCIS_LOG_LEVEL<br/>NATS_LOG_LEVEL | string |  | The log level. Valid values are: "panic", "fatal", "error", "warn", "info", "debug", "trace".|
| OCIS_LOG_PRETTY<br/>NATS_LOG_PRETTY | bool | false | Activates pretty log output.|
| OCIS_LOG_COLOR<br/>NATS_LOG_COLOR | bool | false | Activates colorized log output.|
| OCIS_LOG_FILE<br/>NATS_LOG_FILE | string |  | The path to the log file. Activates logging to this file if set.|
| NATS_DEBUG_ADDR | string | 127.0.0.1:9234 | Bind address of the debug server, where metrics, health, config and debug endpoints will be exposed.|
| NATS_DEBUG_TOKEN | string |  | Token to secure the metrics endpoint.|
| NATS_DEBUG_PPROF | bool | false | Enables pprof, which can be used for profiling.|
| NATS_DEBUG_ZPAGES | bool | false | Enables zpages, which can be used for collecting and viewing in-memory traces.|
| NATS_NATS_HOST | string | 127.0.0.1 | Bind address.|
| NATS_NATS_PORT | int | 9233 | Bind port.|
| NATS_NATS_CLUSTER_ID | string | ocis-cluster | ID of the NATS cluster.|
| NATS_NATS_STORE_DIR | string | ~/.ocis/nats | Path for the NATS JetStream persistence directory.|