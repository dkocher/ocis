package config

// NOTE: Most of this configuration is not needed to keep it as simple as possible
// TODO: Clean up unneeded configuration

func DefaultConfig() *Config {
	return &Config{
		Service: Service{
			Name: "notifications",
		},
	}
}
