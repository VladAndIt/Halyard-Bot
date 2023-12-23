package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Telegram `yaml:"telegram"`
}

type Telegram struct {
	Tocken string `yaml:"tocken"`
}

// NewConfig returns a new decoded Config struct
func MustLoad() *Config {
	// Create config structure
	cfg := &Config{}

	// Open config file
	file, err := os.Open("config/config.yaml")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Init new YAML decode
	d := yaml.NewDecoder(file)

	// Start YAML decoding from file
	if err := d.Decode(&cfg); err != nil {
		log.Fatal(err)
	}

	return cfg
}
