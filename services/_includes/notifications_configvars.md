## Environment Variables

| Name | Type | Default Value | Description |
|------|------|---------------|-------------|
| OCIS_LOG_LEVEL<br/>NOTIFICATIONS_LOG_LEVEL | string |  | The log level. Valid values are: "panic", "fatal", "error", "warn", "info", "debug", "trace".|
| OCIS_LOG_PRETTY<br/>NOTIFICATIONS_LOG_PRETTY | bool | false | Activates pretty log output.|
| OCIS_LOG_COLOR<br/>NOTIFICATIONS_LOG_COLOR | bool | false | Activates colorized log output.|
| OCIS_LOG_FILE<br/>NOTIFICATIONS_LOG_FILE | string |  | The path to the log file. Activates logging to this file if set.|
| NOTIFICATIONS_DEBUG_ADDR | string | 127.0.0.1:9174 | Bind address of the debug server, where metrics, health, config and debug endpoints will be exposed.|
| NOTIFICATIONS_DEBUG_TOKEN | string |  | Token to secure the metrics endpoint.|
| NOTIFICATIONS_DEBUG_PPROF | bool | false | Enables pprof, which can be used for profiling.|
| NOTIFICATIONS_DEBUG_ZPAGES | bool | false | Enables zpages, which can be used for collecting and viewing in-memory traces.|
| NOTIFICATIONS_SMTP_HOST | string |  | SMTP host to connect to.|
| NOTIFICATIONS_SMTP_PORT | int | 1025 | Port of the SMTP host to connect to.|
| NOTIFICATIONS_SMTP_SENDER | string | ownCloud &lt;noreply@example.com&gt; | Sender address of emails that will be sent.|
| NOTIFICATIONS_SMTP_USERNAME | string |  | Username for the SMTP host to connect to.|
| NOTIFICATIONS_SMTP_PASSWORD | string |  | Password for the SMTP host to connect to.|
| NOTIFICATIONS_SMTP_INSECURE | bool | false | Allow insecure connections to the SMTP server.|
| NOTIFICATIONS_SMTP_AUTHENTICATION | string | none | Authentication method for the SMTP communication. Possible values are 'login', 'plain', 'crammd5', 'none'|
| NOTIFICATIONS_SMTP_ENCRYPTION | string | none | Encryption method for the SMTP communication. Possible values  are 'starttls', 'ssl', 'ssltls', 'tls'  and 'none'.|
| NOTIFICATIONS_EVENTS_ENDPOINT | string | 127.0.0.1:9233 | The address of the event system. The event system is the message queuing service. It is used as message broker for the microservice architecture.|
| NOTIFICATIONS_EVENTS_CLUSTER | string | ocis-cluster | The clusterID of the event system. The event system is the message queuing service. It is used as message broker for the microservice architecture. Mandatory when using NATS as event system.|
| NOTIFICATIONS_EVENTS_GROUP | string | notifications | Name of the event group / queue on the event system.|
| REVA_GATEWAY<br/>NOTIFICATIONS_REVA_GATEWAY | string | 127.0.0.1:9142 | CS3 gateway used to look up user metadata|
| OCIS_MACHINE_AUTH_API_KEY<br/>NOTIFICATIONS_MACHINE_AUTH_API_KEY | string |  | Machine auth API key used to validate internal requests necessary to access resources from other services.|