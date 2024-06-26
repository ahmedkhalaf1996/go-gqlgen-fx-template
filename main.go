package main

import (
	"fmt"

	"github.com/dan6erbond/go-gqlgen-fx-template/pkg"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	viper.SetDefault("server.port", 5001)
	viper.SetDefault("server.host", "0.0.0.0")

	viper.SetDefault("db.port", 5432)
	viper.SetDefault("db.sslmode", "disable")

	viper.AutomaticEnv()

	// User:     "root",
	// Password: "secret",
	// Database: "meetmeup_dev",
	viper.BindEnv("db.host", "DB_HOST", "POSTGRES_HOST")
	viper.BindEnv("db.port", "DB_PORT", "POSTGRES_PORT")
	viper.BindEnv("db.user", "DB_USER", "POSTGRES_USER")
	viper.BindEnv("db.password", "DB_PASSWORD", "POSTGRES_PASSWORD")
	viper.BindEnv("db.database", "DB_DATABASE", "POSTGRES_DATABASE")

	viper.BindEnv("environment", "GO_ENV")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	// var DataB = pkg.DB
	app := pkg.NewApp()
	app.Run()
}
