package configs

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

var cfg *config

type config struct {
	API APIConfig
	DB  DBConfig
}

type APIConfig struct {
	Port string
}

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DataBase string
}

func Load() error {

	absolutePath, err := filepath.Abs("config.toml")
	if err != nil {
		fmt.Println("Erro ao obter o caminho absoluto:", err)
		return err
	}

	if _, err = toml.DecodeFile(absolutePath, &cfg); err != nil {
		log.Fatal(err)
		return err
	}
	fmt.Println(&cfg)

	return nil
}

func GetDB() DBConfig {
	return cfg.DB
}

func GetServerPort() string {
	return cfg.API.Port
}
