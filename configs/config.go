package configs

import (
	"github.com/spf13/viper"
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

func init() {

	viper.SetDefault("api.port", "9000")
	viper.SetDefault("database.host", "localhost")
	viper.SetDefault("database.port", "5432")
}

func Load() error {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		//valida tipo de erro for diferente de nao encontrei o arquivo
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return err
		}
	}

	//cria um ponteiro da nossa struct
	cfg = new(config)

	cfg.API = APIConfig{
		Port: viper.GetString("api.port"),
	}

	cfg.DB = DBConfig{
		Host:     viper.GetString("database.host"),
		Port:     viper.GetString("database.port"),
		User:     viper.GetString("database.user"),
		Password: viper.GetString("database.password"),
		DataBase: viper.GetString("database.dbname"),
	}

	//porque nao houve nenhum erro, no golang sempre retornamos nil quando a funcao nao retorna erro
	return nil

}

// accessar as informacoes
func GetDB() DBConfig {
	return cfg.DB
}

func GetServerPort() string {
	return cfg.API.Port
}
