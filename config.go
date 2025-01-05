package main

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Config struct {
	Port           int    `json:"port"`
	Host           string `json:"host"`
	LogLevel       string `json:"log_level"`
	DataDir        string `json:"data_dir"`
	RetentionDays  int    `json:"retention_days"`
	MaxRequestSize int64  `json:"max_request_size"`
	RequestTimeout int    `json:"request_timeout"`
	EnableMetrics  bool   `json:"enable_metrics"`
	MetricsPort    int    `json:"metrics_port"`
}

var DefaultConfig = Config{
	Port:           9988,
	Host:           "127.0.0.1",
	LogLevel:       "info",
	DataDir:        getDefaultDataDir(),
	RetentionDays:  30,
	MaxRequestSize: 10 * 1024 * 1024, // 10MB
	RequestTimeout: 300,              // 300 seconds
	EnableMetrics:  true,
	MetricsPort:    9989,
}

func getDefaultDataDir() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "data"
	}
	return filepath.Join(homeDir, ".deepspace")
}

func LoadConfig(path string) (*Config, error) {
	config := DefaultConfig

	if path == "" {
		return &config, nil
	}

	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return &config, nil
		}
		return nil, err
	}

	if err := json.Unmarshal(data, &config); err != nil {
		return nil, err
	}

	return &config, nil
}

func (c *Config) Save(path string) error {
	data, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return err
	}

	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	return os.WriteFile(path, data, 0644)
}

func GetConfigPath() string {
	if configPath := os.Getenv("DEEPSPACE_CONFIG"); configPath != "" {
		return configPath
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "config.json"
	}

	return filepath.Join(homeDir, ".deepspace", "config.json")
}
