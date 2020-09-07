package config

import (
	"os"

	"github.com/spf13/cast"
)

// Base configuration struct
type Config struct {
	Environment string // development, staging, production
	Lang        string // uz, ru, en

	CassandraHost     string
	CassandraPort     int
	CassandraKeyspace string
	CassandraPassword string
	CassandraUser     string

	LogLevel string
	RPCPort  string
}

func Load() Config {
	c := Config{}
	c.Environment = cast.ToString(getOrReturnDefault("ENVIRONMENT", "development"))
	c.Lang = cast.ToString(getOrReturnDefault("DEFAULT_LANG", "ru"))

	c.CassandraHost = cast.ToString(getOrReturnDefault("CASSANDRA_HOST", "localhost"))
	c.CassandraPort = cast.ToInt(getOrReturnDefault("CASSANDRA_PORT", 9042))
	c.CassandraKeyspace = cast.ToString(getOrReturnDefault("CASSANDRA_KEYSPACE", "todo_service"))
	c.CassandraPassword = cast.ToString(getOrReturnDefault("CASSANDRA_PASSWORD", "cassandra"))
	c.CassandraUser = cast.ToString(getOrReturnDefault("CASSANDRA_USER", "cassandra"))

	c.LogLevel = cast.ToString(getOrReturnDefault("LOG_LEVEL", "debug"))
	c.RPCPort = cast.ToString(getOrReturnDefault("RPC_PORT", ":8005"))
	return c
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}
	return defaultValue
}
