package database

import (
	"fmt"
	"os"

)
type DBConfig interface{
	Dsn() string
	DbName() string
}
type config struct {
	dbUser string
	dbPass string
	dbHost string
	dbName string
	dbPort string
	dsn string
}

func NewConfig() DBConfig {
	cfg := config{
		dbUser : os.Getenv("DATABASE_USER"),
		dbPass : os.Getenv("DATABASE_PASS"),
		dbHost : os.Getenv("DATABASE_HOST"),
		dbName  :os.Getenv("DATABASE_NAME"),
		dbPort : os.Getenv("DATABASE_PORT"),
	}
	cfg.dsn = fmt.Sprintf("mongodb://%s:%s@%s:%d/%s", cfg.dbUser, cfg.dbPass, cfg.dbHost, cfg.dbPort, cfg.dbName)
	return &cfg
}

func (c *config) Dsn() string {
	return c.dsn
}

func (c *config) DbName() string {
	return c.dbName
}