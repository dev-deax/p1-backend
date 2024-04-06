package config

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func NewConfig() *Config {
	config := new(Config)
	config.initialize()
	return config
}

type Config struct {
	username   string
	password   string
	dbName     string
	dbHost     string
	dbPort     string
	ServerHost string
	AppKeys    []string
	RateMinute int
	JwtSecret  string
}

func (config *Config) initialize() {

	config.username = os.Getenv("MYSQL_USER")
	config.password = os.Getenv("MYSQL_PASS")
	config.dbHost = os.Getenv("MYSQL_HOST")
	config.dbPort = os.Getenv("MYSQL_PORT")
	config.dbName = os.Getenv("MYSQL_DB")
	config.ServerHost = os.Getenv("SERVER_PORT")
	config.AppKeys = strings.Split(os.Getenv("APP_KEYS"), ",")
	config.JwtSecret = os.Getenv("JWT_SECRET")
	rateLimit, errorRateLimit := strconv.Atoi(os.Getenv("RATE_MINUTE"))
	if errorRateLimit != nil {
		log.Fatal(errorRateLimit)
	}
	config.RateMinute = rateLimit
}

func (config *Config) DatabaseURI() string {
	dbURI := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.username,
		config.password,
		config.dbHost, config.dbPort,
		config.dbName)
	return dbURI
}
