package config

type App struct {
	Env  string `mapstructure:"env" json:"env" yaml:"env"`
	Host string `mapstructure:"host" json:"host" yaml:"host"`
	Port string `mapstructure:"port" json:"port" yaml:"port"`
	Name string `mapstructure:"name" json:"name" yaml:"name"`
}
