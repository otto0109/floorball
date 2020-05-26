package main

import (
	"first-project/api"
	"first-project/internal/config"
	"first-project/internal/service"
	_ "github.com/aws/aws-sdk-go/aws/session"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"os"
)

func main() {
	configFolder := "./configs/"
	env, envExists := os.LookupEnv("ENV")
	if envExists == false {
		env = "production"
	}
	vicciBaseConfig, err := config.LoadVicciBaseURL(env, configFolder)
	if err != nil {
		log.Fatal(err) // Terminate the application if the config is broken
	}

	router := gin.Default()

	carlineService := service.SetupCarlineService(*vicciBaseConfig)
	api.ProvideCarlineApi(router, carlineService)

	if err := router.Run(); err != nil {
		panic(err)
	}
}
