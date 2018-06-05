/**
 * Copyright 2018 Kazuhiko Sato
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

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
