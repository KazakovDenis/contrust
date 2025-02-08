package contrust

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	ServerURL string `yaml:"serverURL"`
}

var config Config

func parseConfig(configPath string) error {
	data, err := os.ReadFile(configPath)
	if err != nil {
		return fmt.Errorf("failed to read config file: %w", err)
	}

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return nil
}

func init() {
	err := parseConfig("contrust.yaml")
	if err != nil {
		fmt.Printf("Failed to load config file: %s\n", err)
		os.Exit(1)
	}
}
