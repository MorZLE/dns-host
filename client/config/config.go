package config

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type Config struct {
	GRPC *grpcConfig `yaml:"grpc"`
}

type grpcConfig struct {
	Host string `yaml:"host"`
}

func NewConfig() *Config {
	return parseConfig()
}

func parseConfig() *Config {
	file, err := os.ReadFile("config/config.yaml")
	if err != nil {
		log.Fatalf("failed to read config file: %v", err)
	}
	var cfg Config

	if err := yaml.Unmarshal(file, &cfg); err != nil {
		log.Fatalf("failed to parse config file: %v", err)
	}

	if cfg.GRPC.Host == "" {
		host := os.Getenv("GRPC_HOST")
		if host == "" {
			log.Fatal("failed to parse GRPC_HOST")
		}
		cfg.GRPC.Host = host
	}

	return &cfg
}
