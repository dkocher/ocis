## Environment Variables

| Name | Type | Default Value | Description |
|------|------|---------------|-------------|
| OCIS_TRACING_ENABLED<br/>STORAGE_USERS_TRACING_ENABLED | bool | false | Activates tracing.|
| OCIS_TRACING_TYPE<br/>STORAGE_USERS_TRACING_TYPE | string |  | |
| OCIS_TRACING_ENDPOINT<br/>STORAGE_USERS_TRACING_ENDPOINT | string |  | The endpoint to the tracing collector.|
| OCIS_TRACING_COLLECTOR<br/>STORAGE_USERS_TRACING_COLLECTOR | string |  | |
| OCIS_LOG_LEVEL<br/>STORAGE_USERS_LOG_LEVEL | string |  | The log level.|
| OCIS_LOG_PRETTY<br/>STORAGE_USERS_LOG_PRETTY | bool | false | Activates pretty log output.|
| OCIS_LOG_COLOR<br/>STORAGE_USERS_LOG_COLOR | bool | false | Activates colorized log output.|
| OCIS_LOG_FILE<br/>STORAGE_USERS_LOG_FILE | string |  | The target log file.|
| STORAGE_USERS_DEBUG_ADDR | string | 127.0.0.1:9159 | |
| STORAGE_USERS_DEBUG_TOKEN | string |  | |
| STORAGE_USERS_DEBUG_PPROF | bool | false | |
| STORAGE_USERS_DEBUG_ZPAGES | bool | false | |
| STORAGE_USERS_GRPC_ADDR | string | 127.0.0.1:9157 | The address of the grpc service.|
| STORAGE_USERS_GRPC_PROTOCOL | string | tcp | The transport protocol of the grpc service.|
| STORAGE_USERS_GRPC_ADDR | string | 127.0.0.1:9158 | The address of the grpc service.|
| STORAGE_USERS_GRPC_PROTOCOL | string | tcp | The transport protocol of the grpc service.|
| OCIS_JWT_SECRET<br/>STORAGE_USERS_JWT_SECRET | string |  | |
| REVA_GATEWAY | string | 127.0.0.1:9142 | |
| STORAGE_USERS_SKIP_USER_GROUPS_IN_TOKEN | bool | false | |
| STORAGE_USERS_DRIVER | string | ocis | The storage driver which should be used by the service|
| STORAGE_USERS_OCIS_ROOT | string | ~/.ocis/storage/users | |
| STORAGE_USERS_OCIS_USER_LAYOUT | string | {{.Id.OpaqueId}} | |
| STORAGE_USERS_OCIS_PERMISSIONS_ENDPOINT | string | 127.0.0.1:9191 | |
| STORAGE_USERS_OCIS_PERSONAL_SPACE_ALIAS_TEMPLATE | string | {{.SpaceType}}/{{.User.Username \| lower}} | |
| STORAGE_USERS_OCIS_GENERAL_SPACE_ALIAS_TEMPLATE | string | {{.SpaceType}}/{{.SpaceName \| replace &#34; &#34; &#34;-&#34; \| lower}} | |
| STORAGE_USERS_OCIS_SHARE_FOLDER | string | /Shares | |
| STORAGE_USERS_S3NG_ROOT | string | ~/.ocis/storage/users | |
| STORAGE_USERS_S3NG_USER_LAYOUT | string | {{.Id.OpaqueId}} | |
| STORAGE_USERS_PERMISSION_ENDPOINT<br/>STORAGE_USERS_S3NG_USERS_PROVIDER_ENDPOINT | string | 127.0.0.1:9191 | |
| STORAGE_USERS_S3NG_REGION | string | default | |
| STORAGE_USERS_S3NG_ACCESS_KEY | string |  | |
| STORAGE_USERS_S3NG_SECRET_KEY | string |  | |
| STORAGE_USERS_S3NG_ENDPOINT | string |  | |
| STORAGE_USERS_S3NG_BUCKET | string |  | |
| STORAGE_USERS_S3NG_PERSONAL_SPACE_ALIAS_TEMPLATE | string | {{.SpaceType}}/{{.User.Username \| lower}} | |
| STORAGE_USERS_S3NG_GENERAL_SPACE_ALIAS_TEMPLATE | string | {{.SpaceType}}/{{.SpaceName \| replace &#34; &#34; &#34;-&#34; \| lower}} | |
| STORAGE_USERS_S3NG_SHARE_FOLDER | string | /Shares | |
| STORAGE_USERS_OWNCLOUDSQL_DATADIR | string | ~/.ocis/storage/owncloud | |
| STORAGE_USERS_OWNCLOUDSQL_SHARE_FOLDER | string | /Shares | |
| STORAGE_USERS_OWNCLOUDSQL_LAYOUT | string | {{.Username}} | |
| STORAGE_USERS_UPLOADINFO_DIR | string | ~/.ocis/storage/uploadinfo | |
| STORAGE_USERS_OWNCLOUDSQL_DB_USERNAME | string | owncloud | |
| STORAGE_USERS_OWNCLOUDSQL_DB_PASSWORD | string | owncloud | |
| STORAGE_USERS_OWNCLOUDSQL_DB_HOST | string |  | |
| STORAGE_USERS_OWNCLOUDSQL_DB_PORT | int | 3306 | |
| STORAGE_USERS_OWNCLOUDSQL_DB_NAME | string | owncloud | |
| STORAGE_USERS_PERMISSION_ENDPOINT<br/>STORAGE_USERS_OWNCLOUDSQL_USERS_PROVIDER_ENDPOINT | string |  | |
| STORAGE_USERS_DATA_SERVER_URL | string | http://localhost:9158/data | |
| STORAGE_USERS_TEMP_FOLDER | string | ~/.ocis/tmp/users | |
| OCIS_INSECURE<br/>STORAGE_USERS_DATAPROVIDER_INSECURE | bool | false | |
| STORAGE_USERS_EVENTS_ENDPOINT | string | 127.0.0.1:9233 | the address of the streaming service|
| STORAGE_USERS_EVENTS_CLUSTER | string | ocis-cluster | the clusterID of the streaming service. Mandatory when using nats|
| STORAGE_USERS_MOUNT_ID | string | 1284d238-aa92-42ce-bdc4-0b0000009157 | |
| STORAGE_USERS_EXPOSE_DATA_SERVER | bool | false | |
| STORAGE_USERS_READ_ONLY | bool | false | |