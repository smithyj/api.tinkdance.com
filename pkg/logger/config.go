package logger

type Config struct {
	// 日志调试
	Debug bool `yaml:"Debug"`
	// 日志存放路径
	// 默认值: ./logs
	Path string `yaml:"Path"`
	// 是否压缩
	Compress bool `yaml:"Compress"`
	// 单个日志文件大小
	// 单位: MB
	// 默认值: 100
	MaxSize int `yaml:"MaxSize"`
	// 日志文件保存周期
	// 默认永久保存
	MaxAge int `yaml:"MaxAge"`
	// 日志文件最大保存文件数
	// 默认最大文件数
	MaxBackups int `yaml:"MaxBackups"`
}
