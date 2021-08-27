package redisx

type Config struct {
	Prefix   string `yaml:"Prefix"`
	Host     string `yaml:"Host"`
	Port     string `yaml:"Port"`
	Password string `yaml:"Password"`
	DB       int    `yaml:"DB"`
}
