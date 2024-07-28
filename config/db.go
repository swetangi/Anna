package config

import (
	"anna/osutils"
	"anna/repo/todorepo"
	"anna/repo/userrepo"
	"fmt"
	"log"
	"strconv"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type dbConfig struct {
	host     string
	port     int
	username string
	password string
	schema   string
}

func getMySqlDBUrl(host string, port int, username, password, schema string) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, schema)
}

func newDb(dbConfig *dbConfig) *gorm.DB {
	dsn := getMySqlDBUrl(dbConfig.host, dbConfig.port, dbConfig.username, dbConfig.password, dbConfig.schema)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(todorepo.Todo{}, userrepo.User{})
	return db
}

func newDbConfig() *dbConfig {
	host, err := osutils.GetEnvVar("DB_HOST")
	if err != nil {
		log.Fatal(err)
	}
	port, err := osutils.GetEnvVar("DB_PORT")
	if err != nil {
		log.Fatal(err)
	}
	portInt, err := strconv.ParseInt(port, 0, 32)
	if err != nil {
		log.Fatal(err)
	}
	schema, err := osutils.GetEnvVar("DB_SCHEMA")
	if err != nil {
		log.Fatal(err)
	}
	username, err := osutils.GetEnvVar("DB_USERNAME")
	if err != nil {
		log.Fatal(err)
	}
	password, err := osutils.GetEnvVar("DB_PASSWORD")
	if err != nil {
		log.Fatal(err)
	}

	return &dbConfig{
		host:     host,
		port:     int(portInt),
		schema:   schema,
		username: username,
		password: password,
	}
}
