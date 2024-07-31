package config

import (
	"flag"
	"github.com/spf13/viper"
	"os"
)

type Config struct {
	App              ConfigApp
	RelationDatabase ConfigRelationDatabase
}

func (config *Config) Init() *Config {
	currentKey := getAbsoluteKey("", "")

	config.App.Init(currentKey)
	config.RelationDatabase.Init(currentKey)

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
	currentKey := getAbsoluteKey(parentKey, "app")

	config.Host = viper.GetString(getAbsoluteKey(currentKey, "host"))
	config.Port = viper.GetInt(getAbsoluteKey(currentKey, "port"))

	return config
}

type ConfigRelationDatabase struct {
	Connection  ConfigRelationDatabaseConnection
	Connections ConfigRelationDatabaseConnections
}

func (config *ConfigRelationDatabase) Init(parentKey string) *ConfigRelationDatabase {
	currentKey := getAbsoluteKey(parentKey, "relation_database")

	config.Connection.Init(currentKey)
	config.Connections.Init(currentKey)

	return config
}

type ConfigRelationDatabaseConnection struct {
	Master    string
	MasterORM string
	Slave     string
	SlaveORM  string
	Test      string
}

func (config *ConfigRelationDatabaseConnection) Init(parentKey string) *ConfigRelationDatabaseConnection {
	currentKey := getAbsoluteKey(parentKey, "connection")

	config.Master = viper.GetString(getAbsoluteKey(currentKey, "master"))
	config.MasterORM = viper.GetString(getAbsoluteKey(currentKey, "master_orm"))
	config.Slave = viper.GetString(getAbsoluteKey(currentKey, "slave"))
	config.SlaveORM = viper.GetString(getAbsoluteKey(currentKey, "slave_orm"))
	config.Test = viper.GetString(getAbsoluteKey(currentKey, "test"))

	return config
}

type ConfigRelationDatabaseConnections struct {
	Sqlite ConfigRelationDatabaseConnectionsSqlite
}

func (config *ConfigRelationDatabaseConnections) Init(parentKey string) *ConfigRelationDatabaseConnections {
	currentKey := getAbsoluteKey(parentKey, "connections")

	config.Sqlite.Init(currentKey)

	return config
}

type ConfigRelationDatabaseConnectionsSqlite struct {
	MasterDSN string
	SlaveDSN  string
	TestDSN   string
}

func (config *ConfigRelationDatabaseConnectionsSqlite) Init(parentKey string) *ConfigRelationDatabaseConnectionsSqlite {
	currentKey := getAbsoluteKey(parentKey, "sqlite")

	config.MasterDSN = viper.GetString(getAbsoluteKey(currentKey, "master_dsn"))
	config.SlaveDSN = viper.GetString(getAbsoluteKey(currentKey, "slave_dsn"))
	config.TestDSN = viper.GetString(getAbsoluteKey(currentKey, "test_dsn"))

	return config
}

func getAbsoluteKey(parentKey string, currentKey string) string {
	if len(parentKey) == 0 {
		return currentKey
	}

	return parentKey + "." + currentKey
}
