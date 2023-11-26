package config

type Configuration struct {
	App    App    `mapstructure:"app" json:"app" yaml:"app"`
	Logger Logger `mapstructure:"logger" json:"logger" yaml:"logger"`
	Jwt    Jwt    `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Redis  Redis  `mapstructure:"redis" json:"redis" yaml:"redis"`
	Mysql  MySQL  `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Pgsql  PGSQL  `mapstructure:"pgsql" json:"pgsql" yaml:"pgsql"`
}
