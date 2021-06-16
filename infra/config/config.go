package config

import "fmt"

var config Config

type DBConfig struct {
	Host     string
	Port     string
	User     string
	DBName   string
	Password string
}

func setDefaultConfig() {
	config.Db = &DBConfig{
		Host:     "localhost",
		Port:     "3306",
		User:     "tonu",
		DBName:   "first_go",
		Password: "tonu",
	}
}

type Config struct {
	Db *DBConfig
}

func Db() *DBConfig {
	return config.Db
}

func DbURL(dbConfig *DBConfig) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)
}

func LoadConfig() {
	setDefaultConfig()
}
