package config

import (
	"log"
	"os"

	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v2"
)

type Config struct {
	App      `yaml:"app"`
	Telegram `yaml:"telegram"`
}

type App struct {
	Profile string `yaml:"profile"`
}

type Telegram struct {
	Tocken string `yaml:"tocken" envconfig:"TELEGRAM_HALYARD_TOCKEN"`
}

// NewConfig returns a new decoded Config struct
func MustLoad(cfg *Config) {
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
	// env
	error := envconfig.Process("", cfg)
	if error != nil {
		log.Fatal("Can't read env")
	}
}
