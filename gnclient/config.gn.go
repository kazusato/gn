package gnclient

import (
	"github.com/mitchellh/go-homedir"
	"path/filepath"
	"github.com/BurntSushi/toml"
	"log"
)

type Config struct {
	Url, UserId, Password string
}

func LoadConfig() (*Config, error) {
	homedir, err := homedir.Dir()
	if err != nil {
		return nil, err
	}
	gndir := filepath.Join(homedir, ".gn")
	gnconf := filepath.Join(gndir, "config")

	var config Config
	if _, err := toml.DecodeFile(gnconf, &config); err != nil {
		log.Println(err)
		return nil, err
	}

	return &config, nil
}
