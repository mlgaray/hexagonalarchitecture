package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/mlgaray/hexagonalarchitecture/applicattion"
	"github.com/mlgaray/hexagonalarchitecture/conf/viper"
)

func main() {
	cfg, err := viper.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Entorno: %s\n", cfg.Environment)
	fmt.Printf("Configuración de la base de datos: %+v\n", cfg.Database)

	env := os.Getenv("SCOPE")
	if env == "" {
		env = "development"
	}
	godotenv.Load(".env." + env)

	// err := godotenv.Load(".env." + env)

	if err != nil {
		log.Fatal("Error loading .env.development file")
	}
	fmt.Printf(os.Getenv("VERSION"))

	fmt.Println("Initializating project...")
	server := applicattion.Init()
	server.Initialize()
}
