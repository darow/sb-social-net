package server

import (
	"encoding/json"
	"os"
	"time"
)

type ConfigDB struct {
	URI     string        `json:"uri"`
	Name    string        `json:"name"`
	Timeout time.Duration `json:"timeout"`
}

type Config struct {
	DB ConfigDB `json:"db"`
}

func NewConfig(configPath string) (*Config, error) {
	f, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}

	var buf []byte
	f.Read(buf)

	cfg := &Config{}
	err = json.NewDecoder(f).Decode(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
