package main

import (
	"flag"
	"github.com/spf13/viper"
	"os"
)

type Config struct {
	App      ConfigApp
	Database ConfigDatabase
}

func (config *Config) Init() *Config {
	currentKey := getCurrentKey("", "")

	config.App.Init(currentKey)
	config.Database.Init(currentKey)

	return config
}

func NewConfig() (*Config, error) {
	dirPath, exists := os.LookupEnv("GO_EXAMPLE_CONFIG_DIR_PATH")

	if !exists {
		flag.StringVar(&dirPath, "config_dir_path", ".", "configuration file directory path")
		flag.Parse()
	}

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(dirPath)

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	config := Config{}

	return config.Init(), nil
}

type ConfigApp struct {
	Host string
	Port int
}

func (config *ConfigApp) Init(parentKey string) *ConfigApp {
	currentKey := getCurrentKey(parentKey, "app.")

	config.Host = viper.GetString(currentKey + "host")
	config.Port = viper.GetInt(currentKey + "port")

	return config
}

type ConfigDatabase struct {
	Connection  ConfigDatabaseConnection
	Connections ConfigDatabaseConnections
}

func (config *ConfigDatabase) Init(parentKey string) *ConfigDatabase {
	currentKey := getCurrentKey(parentKey, "database.")

	config.Connection.Init(currentKey)
	config.Connections.Init(currentKey)

	return config
}

type ConfigDatabaseConnection struct {
	Master string
	Slave  string
}

func (config *ConfigDatabaseConnection) Init(parentKey string) *ConfigDatabaseConnection {
	currentKey := getCurrentKey(parentKey, "connection.")

	config.Master = viper.GetString(currentKey + "master")
	config.Master = viper.GetString(currentKey + "slave")

	return config
}

type ConfigDatabaseConnections struct {
	Sqlite ConfigDatabaseConnectionsSqlite
}

func (config *ConfigDatabaseConnections) Init(parentKey string) *ConfigDatabaseConnections {
	currentKey := getCurrentKey(parentKey, "connections.")

	config.Sqlite.Init(currentKey)

	return config
}

type ConfigDatabaseConnectionsSqlite struct {
	Dsn string
}

func (config *ConfigDatabaseConnectionsSqlite) Init(parentKey string) *ConfigDatabaseConnectionsSqlite {
	currentKey := getCurrentKey(parentKey, "sqlite.")

	config.Dsn = viper.GetString(currentKey + "dsn")

	return config
}

func getCurrentKey(parentKey string, currentKey string) string {
	if len(parentKey) > 0 {
		return currentKey
	}

	return parentKey + "." + currentKey
}
