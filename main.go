package main

import (
	"log"
	"strconv"

	"github.com/spf13/viper"

	"github.com/i1kondratiuk/visitors-counter/application"
	"github.com/i1kondratiuk/visitors-counter/config"
	"github.com/i1kondratiuk/visitors-counter/domain/repository"
	"github.com/i1kondratiuk/visitors-counter/domain/service"
	"github.com/i1kondratiuk/visitors-counter/infrastructure/persistence"
	"github.com/i1kondratiuk/visitors-counter/interface/web"
)

func init() {
	// To load environmental variables.
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	var configuration config.Configuration

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	err := viper.Unmarshal(&configuration)
	if err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}

	db, err := persistence.NewDbConnection(configuration.Database.Driver,
		configuration.Database.Host,
		strconv.Itoa(int(configuration.Database.Port)),
		configuration.Database.Name,
		configuration.Database.User,
		configuration.Database.Password,
	)

	repository.InitAuthRepository(persistence.NewAuthRepository(db))
	repository.InitUserRepository(persistence.NewUserRepository(db))

	if err != nil {
		panic(err.Error())
	}
}

func main() {
	application.InitAuthApp(&application.AuthAppImpl{})
	application.InitUsersCounterApp(&application.UsersCounterAppImpl{})

	service.InitAuthService(&service.AuthServiceImpl{})

	web.Run(8080)
}
