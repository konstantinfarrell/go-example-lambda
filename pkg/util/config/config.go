package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	DB 		*Database
	Cache 	*Cache
}

type Database struct {
	User	string	`envconfig:"PG_USER"`
	Pass	string	`envconfig:"PG_PASS"`
	Name	string	`envconfig:"PG_NAME"`
	Port	int		`envconfig:"PG_PORT"`
	Addr	string	`envconfig:"PG_ADDR"`
}

type Cache struct {
	Addr	string	`envconfig:"REDIS_ADDR"`
	Pass	string	`envconfig:"REDIS_PASS"`
	DB		int		`envconfig:"REDIS_DB"`
}

func LoadFromEnvVar() (*Config, error){
	var cache Cache
	err := envconfig.Process("", &cache)
	if err != nil {
		return nil, err
	}

	var db Database
	err = envconfig.Process("", &db)
	if err != nil {
		return nil, err
	}

	return &Config{DB: &db, Cache: &cache}, nil
}