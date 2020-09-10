package main

import (
	"floorball/api"
	"floorball/internal/config"
	"floorball/internal/repository"
	"floorball/internal/service"
	"floorball/internal/util"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func main() {
	configFolder := "./configs/"
	env, envExists := os.LookupEnv("ENV")
	if envExists == false {
		env = "production"
	}
	databaseConfig, err := config.LoadDatabaseConfig(env, configFolder)
	if err != nil {
		log.Fatal(err) // Terminate the application if the config is broken
	}

	router := gin.Default()

	db := util.InitDB(databaseConfig.ConnectionString())

	defer func() {
		err = db.Close()
		if err != nil {
			panic(err)
		}
	}()

	playerRepository := repository.ProvidePlayerRepository(db)
	playerService := service.ProvidePlayerService(playerRepository)
	api.InitRouterPlayer(playerService, router)

	articleRepository := repository.ProvideArticleRepository(db)
	articleService := service.ProvideArticleService(articleRepository)
	api.InitRouterArticle(articleService, router)

	if err := router.Run(); err != nil {
		panic(err)
	}
}
