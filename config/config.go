package config

import (
	"fmt"
	"github.com/tyrm/gemini-forum/util"
	"os"
	"strconv"
	"strings"
)

// Config hold config collected from the environment.
type Config struct {
	LoggerConfig string

	PostgresDsn string

	RedisAddress  string
	RedisDB       int
	RedisPassword string
}

// CollectConfig will gather configuration from env vars and return a Config object
func CollectConfig(requiredVars []string) (*Config, error) {
	var config Config

	// LOG_LEVEL
	config.LoggerConfig = os.Getenv("LOG_LEVEL")
	if config.LoggerConfig == "" {
		config.LoggerConfig = "<root>=INFO"
	} else {
		config.LoggerConfig = fmt.Sprintf("<root>=%s", strings.ToUpper(config.LoggerConfig))
	}

	// POSTGRES_DSN
	config.PostgresDsn = os.Getenv("POSTGRES_DSN")
	if config.PostgresDsn != "" {
		requiredVars = util.FastPopString(requiredVars, "POSTGRES_DSN")
	}

	// REDIS_DNS_ADDRESS
	config.RedisAddress = os.Getenv("REDIS_ADDRESS")
	if config.RedisAddress != "" {
		requiredVars = util.FastPopString(requiredVars, "REDIS_ADDRESS")
	}

	// REDIS_DNS_DB
	if os.Getenv("REDIS_DB") == "" {
		config.RedisDB = 0
	} else {
		var err error
		config.RedisDB, err = strconv.Atoi(os.Getenv("REDIS_DB"))
		if err != nil {
			return nil, err
		}
	}

	// REDIS_DNS_PASSWORD
	config.RedisPassword = os.Getenv("REDIS_PASSWORD")

	// Validation
	if len(requiredVars) > 0 {
		return nil, fmt.Errorf("Environment variables missing: %v", requiredVars)
	}

	return &config, nil
}
