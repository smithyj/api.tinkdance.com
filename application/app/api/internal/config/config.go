package config

import (
	"embed"
	"fmt"
	"gopkg.in/yaml.v2"
	"tinkdance/pkg/logger"
	"tinkdance/pkg/redisx"
)

//go:embed *
var fs embed.FS

type Config struct {
	Name  string        `yaml:"Name"`
	Mode  string        `yaml:"Mode"`
	Addr  string        `yaml:"Addr"`
	Log   logger.Config `yaml:"Log"`
	Redis redisx.Config `yaml:"Redis"`
}

func NewConfig(env string) (config Config, err error) {
	var out []byte
	filename := fmt.Sprintf("%v.yaml", env)
	if out, err = fs.ReadFile(filename); err != nil {
		return
	}
	if err = yaml.Unmarshal(out, &config); err != nil {
		return
	}
	if config.Mode == "debug" {
		config.Log.Debug = true
	}
	return
}
