package main

import (
	"fmt"
	"net/http"
	"os"

	"usermanagement/models"
	userHandler "usermanagement/user/handler"
	userRepository "usermanagement/user/repository"
	userService "usermanagement/user/service"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile("config.json")
	err := viper.ReadInConfig()
	if err != nil {
		logrus.Fatal(err)
	}
}

func main() {
	dbHost := viper.GetString("database.host")
	dbPort := viper.GetString("database.port")
	dbUser := viper.GetString("database.user")
	dbPass := viper.GetString("database.pass")
	dbName := viper.GetString("database.name")
	serverAddress := os.Getenv("PORT")
	if serverAddress == "" {
		logrus.Fatal("$PORT must be set")
	}

	connUri := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=require", dbHost, dbPort, dbUser, dbName, dbPass)
	dbConnect, err := gorm.Open("postgres", connUri)
	if err != nil {
		logrus.Fatal("[main.connection.open] ", err)
	}

	err = dbConnect.DB().Ping()
	if err != nil {
		logrus.Error("[main.connection.ping] ", err)
	}

	defer func() {
		err := dbConnect.Close()
		if err != nil {
			logrus.Error("[main.connection.close] ", err)
		}
	}()

	dbConnect.Debug().AutoMigrate(
		&models.User{},
	)

	router := mux.NewRouter().StrictSlash(true)

	userRepository := userRepository.CreateUserRepositoryImpl(dbConnect)
	userService := userService.CreateUserService(userRepository)
	userHandler.CreateUserHandler(router, userService)

	logrus.Info("Starting web server at ", serverAddress)
	logrus.Fatal(http.ListenAndServe(":"+serverAddress, router))
}
