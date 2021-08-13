package config

import "github.com/jinzhu/configor"

type Config struct {
	RepositoryType string `env:"RepositoryType" default:"memory"`
	HttpServer
	MySQL
}

type MySQL struct {
	CfgAddress		string  `env:"DBAdress" default:"0.0.0.0"`
	CfgDBPort		int 	`env:"DBPort" default:"3306"`
	CfgDBName		string	`env:"DBName" default:"vehicle_registration"`
	CfgDBUser     string `env:"DBUser" default:"root"`
	CfgDBPassword string `env:"DBPassword" default:"123qweasdzxc"`
}

func (c Config) DBAddress() string {
	return c.CfgAddress
}

func (c Config) DBPort() int {
	return c.CfgDBPort
}

func (c Config) DBName() string {
	return c.CfgDBName
}

func (c Config) DBUser() string {
	return c.CfgDBUser
}

func (c Config) DBPassword() string {
	return c.CfgDBPassword
}

type HttpServer struct {
	CfgServerAddress string `env:"Address" default:"0.0.0.0"`
	CfgServerPort int		`env:"Port" default:"8888"`
}

func (c Config) ServerAddress() string {
	return c.CfgServerAddress
}

func (c Config) ServerPort() int {
	return c.CfgServerPort
}

func NewConfig() * Config {
	cfg := new(Config)
	if err := configor.New(&configor.Config{Debug: true}).Load(cfg); err != nil {
		panic(err)
	}
	return cfg
}