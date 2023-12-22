package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Telegram struct {
		Tocken string `yaml:"tocken"`
	} `yaml:"telegram"`
}

// NewConfig returns a new decoded Config struct
func NewConfig() *Config {
	// Create config structure
	config := &Config{}

	// Open config file
	file, err := os.Open("config/config.yaml")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Init new YAML decode
	d := yaml.NewDecoder(file)

	// Start YAML decoding from file
	if err := d.Decode(&config); err != nil {
		log.Fatal(err)
	}

	return config
}
