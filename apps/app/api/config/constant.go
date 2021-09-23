package config

const (
	// Version 项目版本号
	Version = "1.0.0"

	// Name 项目名称
	Name = "tinkdance"

	// AccessLogFile 项目访问日志存放文件
	AccessLogFile = "./logs/" + Name + "-access.log"

	// TraceLogFile 项目链路日志存放文件
	TraceLogFile = "./logs/" + Name + "-trace.log"
)
