package env

import (
	"flag"
	"fmt"
	"strings"
)

var (
	active Env
	dev    = &env{value: "dev"}
	test   = &env{value: "test"}
	prod   = &env{value: "prod"}
)

type Env interface {
	Value() string
	IsDev() bool
	IsTest() bool
	IsProd() bool
	i()
}

type env struct {
	value string
}

func (e *env) Value() string {
	return e.value
}

func (e *env) IsDev() bool {
	return e.value == dev.Value()
}

func (e *env) IsTest() bool {
	return e.value == test.Value()
}

func (e *env) IsProd() bool {
	return e.value == prod.Value()
}

func (e *env) i() {}

func init() {
	env := flag.String("env", "", "运行环境:\n  dev: 开发环境\n  test: 测试环境\n  prod: 正式环境")

	flag.Parse()
	switch strings.ToLower(strings.TrimSpace(*env)) {
	case "dev":
		active = dev
	case "test":
		active = test
	case "prod":
		active = prod
	default:
		active = dev
		fmt.Println("Warning: '-env' cannot be found, or it is illegal. The default 'dev' will be used.")
	}
}

func Active() Env {
	return active
}
