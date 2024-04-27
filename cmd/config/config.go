package config

import "os"

type Config struct {
	KafkaBootstrapServers string
	ServerAddress         string
}

func Load() (*Config, error) {
	return &Config{
		KafkaBootstrapServers: os.Getenv("KAFKA_BOOTSTRAP_SERVERS"),
		ServerAddress:         "0.0.0.0:8080",
	}, nil
}
