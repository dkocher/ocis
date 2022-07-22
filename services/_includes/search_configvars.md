## Environment Variables

| Name | Type | Default Value | Description |
|------|------|---------------|-------------|
| OCIS_TRACING_ENABLED<br/>SEARCH_TRACING_ENABLED | bool | false | Activates tracing.|
| OCIS_TRACING_TYPE<br/>SEARCH_TRACING_TYPE | string |  | The type of tracing. Defaults to "", which is the same as "jaeger". Allowed tracing types are "jaeger" and "" as of now.|
| OCIS_TRACING_ENDPOINT<br/>SEARCH_TRACING_ENDPOINT | string |  | The endpoint of the tracing agent.|
| OCIS_TRACING_COLLECTOR<br/>SEARCH_TRACING_COLLECTOR | string |  | The HTTP endpoint for sending spans directly to a collector, i.e. http://jaeger-collector:14268/api/traces. Only used if the tracing endpoint is unset.|
| OCIS_LOG_LEVEL<br/>SEARCH_LOG_LEVEL | string |  | The log level. Valid values are: "panic", "fatal", "error", "warn", "info", "debug", "trace".|
| OCIS_LOG_PRETTY<br/>SEARCH_LOG_PRETTY | bool | false | Activates pretty log output.|
| OCIS_LOG_COLOR<br/>SEARCH_LOG_COLOR | bool | false | Activates colorized log output.|
| OCIS_LOG_FILE<br/>SEARCH_LOG_FILE | string |  | The path to the log file. Activates logging to this file if set.|
| SEARCH_DEBUG_ADDR | string | 127.0.0.1:9224 | Bind address of the debug server, where metrics, health, config and debug endpoints will be exposed.|
| SEARCH_DEBUG_TOKEN | string |  | Token to secure the metrics endpoint.|
| SEARCH_DEBUG_PPROF | bool | false | Enables pprof, which can be used for profiling.|
| SEARCH_DEBUG_ZPAGES | bool | false | Enables zpages, which can be used for collecting and viewing in-memory traces.|
| SEARCH_GRPC_ADDR | string | 127.0.0.1:9220 | The bind address of the GRPC service.|
| SEARCH_DATA_PATH | string | ~/.ocis/search | Path for the search persistence directory.|
| REVA_GATEWAY | string | 127.0.0.1:9142 | The CS3 gateway endpoint.|
| SEARCH_EVENTS_ENDPOINT | string | 127.0.0.1:9233 | |
| SEARCH_EVENTS_CLUSTER | string | ocis-cluster | The clusterID of the event system. The event system is the message queuing service. It is used as message broker for the microservice architecture. Mandatory when using NATS as event system.|
| SEARCH_EVENTS_GROUP | string | search | The customer group of the service. One group will only get one copy of an event|
| OCIS_MACHINE_AUTH_API_KEY<br/>SEARCH_MACHINE_AUTH_API_KEY | string |  | Machine auth API key used to validate internal requests necessary for the access to resources from other services.|