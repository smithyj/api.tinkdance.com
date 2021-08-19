package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"tinkdance/pkg/tinkgo/logx"
	"tinkdance/service/app/api/internal/assets"
)

type Config struct {
	Name      string      `yaml:"Name"`
	Mode      string      `yaml:"Mode"`
	Addr      string      `yaml:"Addr"`
	LogConfig logx.Config `yaml:"LogConfig"`
}

func NewConfig(env string) (config *Config, err error) {
	var out []byte
	filename := fmt.Sprintf("config/%v.yaml", env)
	if out, err = assets.FS.ReadFile(filename); err != nil {
		return
	}
	if err = yaml.Unmarshal(out, &config); err != nil {
		return
	}
	if config.Mode == "debug" {
		config.LogConfig.Debug = true
	}
	return
}
