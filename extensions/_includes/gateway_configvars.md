## Environment Variables

| Name | Type | Default Value | Description |
|------|------|---------------|-------------|
| OCIS_TRACING_ENABLED<br/>GATEWAY_TRACING_ENABLED | bool | false | Activates tracing.|
| OCIS_TRACING_TYPE<br/>GATEWAY_TRACING_TYPE | string |  | |
| OCIS_TRACING_ENDPOINT<br/>GATEWAY_TRACING_ENDPOINT | string |  | The endpoint to the tracing collector.|
| OCIS_TRACING_COLLECTOR<br/>GATEWAY_TRACING_COLLECTOR | string |  | |
| OCIS_LOG_LEVEL<br/>GATEWAY_LOG_LEVEL | string |  | The log level.|
| OCIS_LOG_PRETTY<br/>GATEWAY_LOG_PRETTY | bool | false | Activates pretty log output.|
| OCIS_LOG_COLOR<br/>GATEWAY_LOG_COLOR | bool | false | Activates colorized log output.|
| OCIS_LOG_FILE<br/>GATEWAY_LOG_FILE | string |  | The target log file.|
| GATEWAY_DEBUG_ADDR | string | 127.0.0.1:9143 | Bind address of the debug server, where metrics, health, config and debug endpoints will be exposed.|
| GATEWAY_DEBUG_TOKEN | string |  | Token to secure the metrics endpoint|
| GATEWAY_DEBUG_PPROF | bool | false | Enables pprof, which can be used for profiling|
| GATEWAY_DEBUG_ZPAGES | bool | false | Enables zpages, which can  be used for collecting and viewing traces in-me|
| GATEWAY_GRPC_ADDR | string | 127.0.0.1:9142 | The address of the grpc service.|
| GATEWAY_GRPC_PROTOCOL | string | tcp | The transport protocol of the grpc service.|
| OCIS_JWT_SECRET<br/>GATEWAY_JWT_SECRET | string |  | |
| REVA_GATEWAY | string | 127.0.0.1:9142 | |
| GATEWAY_SKIP_USER_GROUPS_IN_TOKEN | bool | false | |
| GATEWAY_COMMIT_SHARE_TO_STORAGE_GRANT | bool | true | |
| GATEWAY_COMMIT_SHARE_TO_STORAGE_REF | bool | true | |
| GATEWAY_SHARE_FOLDER_NAME | string | Shares | |
| GATEWAY_DISABLE_HOME_CREATION_ON_LOGIN | bool | true | |
| STORAGE_TRANSFER_SECRET | string |  | |
| GATEWAY_TRANSFER_EXPIRES | int | 86400 | |
| GATEWAY_HOME_MAPPING | string |  | |
| GATEWAY_ETAG_CACHE_TTL | int | 0 | |
| OCIS_URL<br/>GATEWAY_FRONTEND_PUBLIC_URL | string | https://localhost:9200 | |
| GATEWAY_USERS_ENDPOINT | string | localhost:9144 | |
| GATEWAY_GROUPS_ENDPOINT | string | localhost:9160 | |
| GATEWAY_PERMISSIONS_ENDPOINT | string | localhost:9191 | |
| GATEWAY_SHARING_ENDPOINT | string | localhost:9150 | |
| GATEWAY_AUTH_BASIC_ENDPOINT | string | localhost:9146 | |
| GATEWAY_AUTH_BEARER_ENDPOINT | string | localhost:9148 | |
| GATEWAY_AUTH_MACHINE_ENDPOINT | string | localhost:9166 | |
| GATEWAY_STORAGE_PUBLIC_LINK_ENDPOINT | string | localhost:9178 | |
| GATEWAY_STORAGE_USERS_ENDPOINT | string | localhost:9157 | |
| GATEWAY_STORAGE_SHARES_ENDPOINT | string | localhost:9154 | |
| GATEWAY_APP_REGISTRY_ENDPOINT | string | localhost:9242 | |