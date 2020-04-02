package config

import (
	"github.com/BurntSushi/toml"
	"github.com/sirupsen/logrus"
)

// Setting ...
var Setting Config

const (
	// Path ...
	// Right setting:Working Directory G:/GoWorkspace/src/gin_web
	//Path = "config/config.toml"

	// Wrong setting:Working Directory G:/GoWorkspace/src/
	//Path = "gin_web/config/config.toml"

	Path = "pkg/config/config.toml"
)

// Init ...
func Init() {
	if err := LoadConfigs(Path); err != nil {
		panic(err)
	}
}

// LoadConfigs ...
func LoadConfigs(files ...string) error {
	for _, file := range files {
		if _, err := toml.DecodeFile(file, &Setting); err != nil {
			logrus.Error("decode error: ", err)
			return err
		}
	}

	return nil
}
