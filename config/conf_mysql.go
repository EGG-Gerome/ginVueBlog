package config

type Mysql struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	DB       string `yaml:"db"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	LogLevel string `yaml:"log_level"` // 日志等级，debug 就是输出全部的 sql，dev 是输出开发环境关键信息，如函数调用、变量值等，release 是只输出生产环境关键业务操作和错误信息
}
