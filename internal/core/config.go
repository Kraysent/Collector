package core

import (
	"fmt"
	"os"
	"time"

	"collector/internal/storage"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Logging struct {
		StdoutPath string `yaml:"stdout_path"`
		StderrPath string `yaml:"stderr_path"`
	} `yaml:"logging"`
	TagCleaner struct {
		Disabled bool          `yaml:"disabled"`
		TagName  string        `yaml:"tag_name"`
		Limit    int           `yaml:"limit"`
		Period   time.Duration `yaml:"period"`
	} `yaml:"tag_cleaner"`
	DurationChecker struct {
		Disabled bool          `yaml:"disabled"`
		Period   time.Duration `yaml:"period"`
	} `yaml:"duration_checker"`
	Clients struct {
		TickTick struct {
			Token string `yaml:"-"`
		}
		WakaTime struct {
			Token string `yaml:"-"`
		}
	}
	Storage storage.Config `yaml:"storage"`
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

	ticktickToken, ok := os.LookupEnv("TICKTICK_TOKEN")
	if !ok {
		return nil, fmt.Errorf("no TickTick token provided")
	}
	config.Clients.TickTick.Token = ticktickToken

	wakatimeToken, ok := os.LookupEnv("WAKATIME_TOKEN")
	if !ok {
		return nil, fmt.Errorf("no WakaTime token provided")
	}
	config.Clients.WakaTime.Token = wakatimeToken

	return &config, nil
}
