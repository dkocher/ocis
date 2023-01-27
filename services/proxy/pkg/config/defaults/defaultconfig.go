package defaults

import (
	"path"
	"strings"

	"github.com/owncloud/ocis/v2/ocis-pkg/config/defaults"
	"github.com/owncloud/ocis/v2/ocis-pkg/shared"
	"github.com/owncloud/ocis/v2/services/proxy/pkg/config"
)

func FullDefaultConfig() *config.Config {
	cfg := DefaultConfig()
	EnsureDefaults(cfg)
	Sanitize(cfg)
	return cfg
}

func DefaultConfig() *config.Config {
	return &config.Config{
		Debug: config.Debug{
			Addr:  "127.0.0.1:9205",
			Token: "",
		},
		HTTP: config.HTTP{
			Addr:      "0.0.0.0:9200",
			Root:      "/",
			Namespace: "com.owncloud.web",
			TLSCert:   path.Join(defaults.BaseDataPath(), "proxy", "server.crt"),
			TLSKey:    path.Join(defaults.BaseDataPath(), "proxy", "server.key"),
			TLS:       true,
		},
		Service: config.Service{
			Name: "proxy",
		},
		OIDC: config.OIDC{
			Issuer: "https://localhost:9200",

			AccessTokenVerifyMethod: config.AccessTokenVerificationJWT,
			UserinfoCache: config.UserinfoCache{
				Size: 1024,
				TTL:  10,
			},
			JWKS: config.JWKS{
				RefreshInterval:   60, // minutes
				RefreshRateLimit:  60, // seconds
				RefreshTimeout:    10, // seconds
				RefreshUnknownKID: true,
			},
		},
		PolicySelector: nil,
		Reva:           shared.DefaultRevaConfig(),
		PreSignedURL: config.PreSignedURL{
			AllowedHTTPMethods: []string{"GET"},
			Enabled:            true,
		},
		AccountBackend:        "cs3",
		UserOIDCClaim:         "preferred_username",
		UserCS3Claim:          "username",
		AutoprovisionAccounts: false,
		EnableBasicAuth:       false,
		InsecureBackends:      false,
	}
}

func DefaultPolicies() []config.Policy {
	return []config.Policy{
		{
			Name: "ocis",
			Routes: []config.Route{
				{
					Endpoint:    "/",
					Service:     "com.owncloud.web.web",
					Unprotected: true,
				},
				{
					Endpoint:    "/.well-known/",
					Service:     "com.owncloud.web.idp",
					Unprotected: true,
				},
				{
					Endpoint:    "/konnect/",
					Service:     "com.owncloud.web.idp",
					Unprotected: true,
				},
				{
					Endpoint:    "/signin/",
					Service:     "com.owncloud.web.idp",
					Unprotected: true,
				},
				{
					Endpoint: "/archiver",
					Service:  "com.owncloud.web.frontend",
				},
				{
					Type:     config.RegexRoute,
					Endpoint: "/ocs/v[12].php/cloud/user/signing-key", // only `user/signing-key` is left in ocis-ocs
					Service:  "com.owncloud.web.ocs",
				},
				{
					Type:        config.RegexRoute,
					Endpoint:    "/ocs/v[12].php/config",
					Service:     "com.owncloud.web.frontend",
					Unprotected: true,
				},
				{
					Endpoint: "/ocs/",
					Service:  "com.owncloud.web.frontend",
				},
				{
					Type:     config.QueryRoute,
					Endpoint: "/remote.php/?preview=1",
					Service:  "com.owncloud.web.webdav",
				},
				// TODO the actual REPORT goes to /dav/files/{username}, which is user specific ... how would this work in a spaces world?
				// TODO what paths are returned? the href contains the full path so it should be possible to return urls from other spaces?
				// TODO or we allow a REPORT on /dav/spaces to search all spaces and /dav/space/{spaceid} to search a specific space
				// send webdav REPORT requests to search service
				{
					Method:   "REPORT",
					Endpoint: "/remote.php/dav/",
					Service:  "com.owncloud.web.webdav",
				},
				{
					Method:   "REPORT",
					Endpoint: "/remote.php/webdav",
					Service:  "com.owncloud.web.webdav",
				},
				{
					Method:   "REPORT",
					Endpoint: "/dav/spaces",
					Service:  "com.owncloud.web.webdav",
				},
				{
					Type:     config.QueryRoute,
					Endpoint: "/dav/?preview=1",
					Service:  "com.owncloud.web.webdav",
				},
				{
					Type:     config.QueryRoute,
					Endpoint: "/webdav/?preview=1",
					Service:  "com.owncloud.web.webdav",
				},
				{
					Endpoint: "/remote.php/",
					Service:  "com.owncloud.web.ocdav",
				},
				{
					Endpoint: "/dav/",
					Service:  "com.owncloud.web.ocdav",
				},
				{
					Endpoint: "/webdav/",
					Service:  "com.owncloud.web.ocdav",
				},
				{
					Endpoint:    "/status",
					Service:     "com.owncloud.web.ocdav",
					Unprotected: true,
				},
				{
					Endpoint:    "/status.php",
					Service:     "com.owncloud.web.ocdav",
					Unprotected: true,
				},
				{
					Endpoint: "/index.php/",
					Service:  "com.owncloud.web.ocdav",
				},
				{
					Endpoint: "/apps/",
					Service:  "com.owncloud.web.ocdav",
				},
				{
					Endpoint:    "/data",
					Service:     "com.owncloud.web.frontend",
					Unprotected: true,
				},
				{
					Endpoint:    "/app/list",
					Service:     "com.owncloud.web.frontend",
					Unprotected: true,
				},
				{
					Endpoint: "/app/", // /app or /apps? ocdav only handles /apps
					Service:  "com.owncloud.web.frontend",
				},
				{
					Endpoint: "/graph/",
					Service:  "com.owncloud.graph.graph",
				},
				{
					Endpoint: "/api/v0/settings",
					Service:  "com.owncloud.web.settings",
				},
			},
		},
	}
}

// EnsureDefaults adds default values to the configuration if they are not set yet
func EnsureDefaults(cfg *config.Config) {
	// provide with defaults for shared logging, since we need a valid destination address for "envdecode".
	if cfg.Log == nil && cfg.Commons != nil && cfg.Commons.Log != nil {
		cfg.Log = &config.Log{
			Level:  cfg.Commons.Log.Level,
			Pretty: cfg.Commons.Log.Pretty,
			Color:  cfg.Commons.Log.Color,
			File:   cfg.Commons.Log.File,
		}
	} else if cfg.Log == nil {
		cfg.Log = &config.Log{}
	}
	// provide with defaults for shared tracing, since we need a valid destination address for "envdecode".
	if cfg.Tracing == nil && cfg.Commons != nil && cfg.Commons.Tracing != nil {
		cfg.Tracing = &config.Tracing{
			Enabled:   cfg.Commons.Tracing.Enabled,
			Type:      cfg.Commons.Tracing.Type,
			Endpoint:  cfg.Commons.Tracing.Endpoint,
			Collector: cfg.Commons.Tracing.Collector,
		}
	} else if cfg.Tracing == nil {
		cfg.Tracing = &config.Tracing{}
	}

	if cfg.TokenManager == nil && cfg.Commons != nil && cfg.Commons.TokenManager != nil {
		cfg.TokenManager = &config.TokenManager{
			JWTSecret: cfg.Commons.TokenManager.JWTSecret,
		}
	} else if cfg.TokenManager == nil {
		cfg.TokenManager = &config.TokenManager{}
	}

	if cfg.MachineAuthAPIKey == "" && cfg.Commons != nil && cfg.Commons.MachineAuthAPIKey != "" {
		cfg.MachineAuthAPIKey = cfg.Commons.MachineAuthAPIKey
	}

	if cfg.Reva == nil && cfg.Commons != nil && cfg.Commons.Reva != nil {
		cfg.Reva = &shared.Reva{
			Address: cfg.Commons.Reva.Address,
			TLS:     cfg.Commons.Reva.TLS,
		}
	} else if cfg.Reva == nil {
		cfg.Reva = &shared.Reva{}
	}

	if cfg.GRPCClientTLS == nil {
		cfg.GRPCClientTLS = &shared.GRPCClientTLS{}
		if cfg.Commons != nil && cfg.Commons.GRPCClientTLS != nil {
			cfg.GRPCClientTLS.Mode = cfg.Commons.GRPCClientTLS.Mode
			cfg.GRPCClientTLS.CACert = cfg.Commons.GRPCClientTLS.CACert
		}
	}
}

// Sanitize sanitizes the configuration
func Sanitize(cfg *config.Config) {
	if cfg.Policies == nil {
		cfg.Policies = DefaultPolicies()
	}

	if cfg.PolicySelector == nil {
		cfg.PolicySelector = &config.PolicySelector{
			Static: &config.StaticSelectorConf{
				Policy: "ocis",
			},
		}
	}

	if cfg.HTTP.Root != "/" {
		cfg.HTTP.Root = strings.TrimSuffix(cfg.HTTP.Root, "/")
	}
}
