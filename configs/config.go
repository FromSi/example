package configs

import (
	"os"
	"strconv"
)

const (
	EnvHost = "GO_EXAMPLE_HOST"
	EnvPort = "GO_EXAMPLE_PORT"

	EnvDefaultHost = "localhost"
	EnvDefaultPort = 8080
)

type Config struct {
	Host string
	Port uint16
}

type BuilderConfig struct {
	config *Config
}

func NewBuilderConfig() *BuilderConfig {
	return &BuilderConfig{
		config: &Config{
			Host: EnvDefaultHost,
			Port: EnvDefaultPort,
		},
	}
}

func (builder *BuilderConfig) Build() (*Config, error) {
	if env, exists := os.LookupEnv(EnvHost); builder.config.Host == EnvDefaultHost && exists {
		builder.SetHost(env)
	}

	if env, exists := os.LookupEnv(EnvPort); builder.config.Port == EnvDefaultPort && exists {
		port, err := strconv.Atoi(env)

		if err != nil {
			return nil, err
		}

		builder.SetPort(uint16(port))
	}

	return builder.config, nil
}

func (builder *BuilderConfig) SetHost(host string) *BuilderConfig {
	builder.config.Host = host

	return builder
}

func (builder *BuilderConfig) SetPort(port uint16) *BuilderConfig {
	builder.config.Port = port

	return builder
}
