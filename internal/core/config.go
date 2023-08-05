package core

import (
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

type Config struct {
	TagCleaner struct {
		TagName string        `yaml:"tag_name"`
		Limit   int           `yaml:"limit"`
		Period  time.Duration `yaml:"period"`
	} `yaml:"tag_cleaner"`
	Clients struct {
		TickTick struct {
			Token string `yaml:"-"`
		}
	}
}

func ParseConfig(filename string) (*Config, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var config Config

	if err := yaml.Unmarshal(file, &config); err != nil {
		return nil, err
	}

	return &config, nil
}
