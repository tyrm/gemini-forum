package config

import (
	"log"
	"os"
	"testing"
)

func TestCollectConfig_Empty(t *testing.T) {
	unsetEnv()

	cfg, err := CollectConfig([]string{})
	if err != nil {
		t.Errorf("enexpected error, got: %v, want: nil.", err.Error())
	}

	if cfg.LoggerConfig != "<root>=INFO" {
		t.Errorf("enexpected config value for LoggerConfig, got: '%s', want: '<root>=INFO'.", cfg.LoggerConfig)
	}
	if cfg.PostgresDsn != "" {
		t.Errorf("enexpected config value for PostgresDsn, got: '%s', want: ''.", cfg.PostgresDsn)
	}
	if cfg.RedisAddress != "" {
		t.Errorf("enexpected config value for RedisDNSAddress, got: '%s', want: ''.", cfg.RedisAddress)
	}
	if cfg.RedisDB != 0 {
		t.Errorf("enexpected config value for RedisDNSDB, got: %d, want: 0.", cfg.RedisDB)
	}
	if cfg.RedisPassword != "" {
		t.Errorf("enexpected config value for RedisDNSPassword, got: '%s', want: ''.", cfg.RedisPassword)
	}
}

func TestCollectConfig_EmptyRequireAll(t *testing.T) {
	unsetEnv()

	requiredEnvVars := []string{
		"POSTGRES_DSN",
		"REDIS_ADDRESS",
	}

	cfg, err := CollectConfig(requiredEnvVars)
	if err == nil {
		t.Errorf("expected error, got: nil, want: err.")
	}

	if cfg != nil {
		t.Errorf("expected config, got: %v, want: nil.", cfg)
	}
}

func TestCollectConfig_InvalidRedisDNSDB(t *testing.T) {
	unsetEnv()

	setEnvVars := map[string]string{
		"REDIS_DB": "astring",
	}

	for k, v := range setEnvVars {
		err := os.Setenv(k, v)
		if err != nil {
			log.Fatal(err)
		}
	}

	cfg, err := CollectConfig([]string{})
	if err == nil {
		t.Errorf("expected error, got: nil, want: error.")
	}
	if cfg != nil {
		t.Errorf("expected config, got: %v, want: nil.", cfg)
	}
}

func TestCollectConfig_Loaded(t *testing.T) {
	unsetEnv()

	setEnvVars := map[string]string{
		"LOG_LEVEL":      "trace",
		"POSTGRES_DSN":   "postgresql://test:test@127.0.0.1:5432/test",
		"REDIS_ADDRESS":  "localhost:6379",
		"REDIS_DB":       "8",
		"REDIS_PASSWORD": "P@ssw0rd!",
	}

	for k, v := range setEnvVars {
		err := os.Setenv(k, v)
		if err != nil {
			log.Fatal(err)
		}
	}

	cfg, err := CollectConfig([]string{})
	if err != nil {
		t.Errorf("enexpected error, got: %v, want: nil.", err.Error())
	}

	if cfg.LoggerConfig != "<root>=TRACE" {
		t.Errorf("enexpected config value for LoggerConfig, got: '%s', want: '<root>=TRACE'.", cfg.LoggerConfig)
	}
	if cfg.PostgresDsn != "postgresql://test:test@127.0.0.1:5432/test" {
		t.Errorf("enexpected config value for PostgresDsn, got: '%s', want: 'postgresql://test:test@127.0.0.1:5432/test'.", cfg.PostgresDsn)
	}
	if cfg.RedisAddress != "localhost:6379" {
		t.Errorf("enexpected config value for RedisDNSAddress, got: '%s', want: 'localhost:6379'.", cfg.RedisAddress)
	}
	if cfg.RedisDB != 8 {
		t.Errorf("enexpected config value for RedisDNSDB, got: %d, want: 8.", cfg.RedisDB)
	}
	if cfg.RedisPassword != "P@ssw0rd!" {
		t.Errorf("enexpected config value for RedisDNSPassword, got: '%s', want: 'P@ssw0rd!'.", cfg.RedisPassword)
	}
}

func unsetEnv() {
	envVars := []string{
		"LOG_LEVEL",
		"POSTGRES_DSN",
		"PRIMARY_NS",
		"REDIS_ADDRESS",
		"REDIS_DB",
		"REDIS_PASSWORD",
	}

	for _, ev := range envVars {
		err := os.Unsetenv(ev)
		if err != nil {
			log.Fatal(err)
		}
	}
}
