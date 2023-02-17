package config

import (
	"io/ioutil"
	"time"

	"gopkg.in/yaml.v2"
)

// Config project config.
type Config struct {
	Tool  Tool  `yaml:"tool"`
	Shell Shell `yaml:"shell"`
}

// Tool Tool config.
type Tool struct {
	Interval time.Duration `yaml:"interval"`
}

// Shell Shell config.
type Shell struct {
	Path string `yaml:"path"`
	Name string `yaml:"name"`
}

// New builds project config from specified file.
func New(filename string) (*Config, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	c := new(Config)
	if err := yaml.UnmarshalStrict(data, c); err != nil {
		return nil, err
	}

	return c, nil
}
