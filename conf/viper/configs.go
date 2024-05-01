package viper

import (
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	Environment string
	Database    map[string]interface{}
	// Otros campos de configuración
}

func LoadConfig() (*Config, error) {
	env := os.Getenv("SCOPE")
	if env == "" {
		// Establecer un valor predeterminado o manejar el error según sea necesario
		env = "development"
	}

	viper.SetConfigName("conf_" + env)
	viper.AddConfigPath("./conf")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	var config Config
	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
