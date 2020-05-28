package api

import (
	"floorball/api/mapper"
	"floorball/internal/entities"
	"github.com/gin-gonic/gin"
	"net/http"
)

type playerService interface {
	GetPlayerByTeam(teamname string) (entities.Team, error)
}

type PlayerApi struct {
	service playerService
}

func providePlayerApi(service playerService) *PlayerApi {
	return &PlayerApi{service: service}
}

func InitRouterPlayer(service playerService, router *gin.Engine) {
	playerApi := providePlayerApi(service)
	router.GET("/player/:team", playerApi.GetPlayerByTeam)
}

func (api *PlayerApi) GetPlayerByTeam(requestContext *gin.Context) {

	param := requestContext.Param("team")

	players, err := api.service.GetPlayerByTeam(param)

	if err != nil {
		return
	}

	requestContext.JSON(http.StatusOK, mapper.TeamToDto(players))
}
