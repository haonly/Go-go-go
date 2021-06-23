package config

import "github.com/jinzhu/configor"

type Config struct {
	RepositoryType string `env:"RepositoryType" default:"memory"`
	HttpServer
	MariaDB
}

func New() *Config {
	cfg := new(Config)
	if err := configor.New(&configor.Config{Debug: true}).Load(cfg); err != nil {
		panic(err)
	}
	return cfg
}

type HttpServer struct {
	CfgServerAddress string `env:"Address" default:"0.0.0.0"`
	CfgServerPort    int    `env:"Port" default:"8888"`
}

func (c Config) ServerAddress() string {
	return c.CfgServerAddress
}

func (c Config) ServerPort() int {
	return c.CfgDBPort
}

type MariaDB struct {
	CfgAddress    string `env:"DBAddress" default:"0.0.0.0"`
	CfgDBPort     int    `env:"DBPort" default:"3306"`
	CfgDBName     string `env:"DBName" default:"mydb"`
	CfgDBUser     string `env:"DBUser" default:"root"`
	CfgDBPassword string `env:"DBPassword" default:"1234"`
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
