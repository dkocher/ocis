## Environment Variables

| Name | Type | Default Value | Description |
|------|------|---------------|-------------|
| OCIS_TRACING_ENABLED<br/>PROXY_TRACING_ENABLED | bool | false | |
| OCIS_TRACING_TYPE<br/>PROXY_TRACING_TYPE | string |  | |
| OCIS_TRACING_ENDPOINT<br/>PROXY_TRACING_ENDPOINT | string |  | |
| OCIS_TRACING_COLLECTOR<br/>PROXY_TRACING_COLLECTOR | string |  | |
| OCIS_LOG_LEVEL<br/>PROXY_LOG_LEVEL | string |  | |
| OCIS_LOG_PRETTY<br/>PROXY_LOG_PRETTY | bool | false | |
| OCIS_LOG_COLOR<br/>PROXY_LOG_COLOR | bool | false | |
| OCIS_LOG_FILE<br/>PROXY_LOG_FILE | string |  | |
| PROXY_DEBUG_ADDR | string | 127.0.0.1:9205 | |
| PROXY_DEBUG_TOKEN | string |  | |
| PROXY_DEBUG_PPROF | bool | false | |
| PROXY_DEBUG_ZPAGES | bool | false | |
| PROXY_HTTP_ADDR | string | 0.0.0.0:9200 | |
| PROXY_HTTP_ROOT | string | / | |
| PROXY_TRANSPORT_TLS_CERT | string | ~/.ocis/proxy/server.crt | |
| PROXY_TRANSPORT_TLS_KEY | string | ~/.ocis/proxy/server.key | |
| PROXY_TLS | bool | true | |
| REVA_GATEWAY | string | 127.0.0.1:9142 | |
| OCIS_URL<br/>OCIS_OIDC_ISSUER<br/>PROXY_OIDC_ISSUER | string | https://localhost:9200 | |
| OCIS_INSECURE<br/>PROXY_OIDC_INSECURE | bool | true | |
| PROXY_OIDC_USERINFO_CACHE_SIZE | int | 1024 | |
| PROXY_OIDC_USERINFO_CACHE_TTL | int | 10 | |
| OCIS_JWT_SECRET<br/>PROXY_JWT_SECRET | string |  | |
| PROXY_ENABLE_PRESIGNEDURLS | bool | true | |
| PROXY_ACCOUNT_BACKEND_TYPE | string | cs3 | |
| PROXY_USER_OIDC_CLAIM | string | email | |
| PROXY_USER_CS3_CLAIM | string | mail | |
| OCIS_MACHINE_AUTH_API_KEY<br/>PROXY_MACHINE_AUTH_API_KEY | string |  | |
| PROXY_AUTOPROVISION_ACCOUNTS | bool | false | |
| PROXY_ENABLE_BASIC_AUTH | bool | false | |
| PROXY_INSECURE_BACKENDS | bool | false | |