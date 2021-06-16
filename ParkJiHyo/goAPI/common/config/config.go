package config

import (
	"github.com/jinzhu/configor"
)

func ReadEnv(envFileName string) *Environment {
	env := new(Environment)
	if err := configor.New(&configor.Config{}).Load(env, envFileName); err != nil {
		panic(err)
	}
	return env
}

type Environment struct {
	Port     int      `toml:"port" default:"8083"`
	Debug    bool     `toml:"debug" default:"false"`
	Database Database `toml:"database"`
}

type Database struct {
	DriverName         string `toml:"driver_name" env:"DRIVER_NAME"`
	DataSourceName     string `toml:"data_source_name" env:"DATA_SOURCE_NAME"`
	MaxIdleConnections int    `toml:"max_idle_conns"`
	MaxOpenConnections int    `toml:"max_open_conns"`
}

func GetDatabase() (driverName, dataSourceName string) {
	user := "root"
	password := "12345"
	driver := "mysql"
	dbName := "go_study"
	dataSource := user + ":" + password + "@tcp(127.0.0.1:3306)/" + dbName + "?charset=utf8&parseTime=True"
	return driver, dataSource
}
