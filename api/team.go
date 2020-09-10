package api

import (
	"floorball/api/mapper"
	"floorball/internal/entities"
	"github.com/gin-gonic/gin"
	"net/http"
)

type playerService interface {
	GetTeamByTeamname(teamname string) (entities.Team, error)
	GetAllTeams() []entities.Team
}

type PlayerApi struct {
	service playerService
}

func providePlayerApi(service playerService) *PlayerApi {
	return &PlayerApi{service: service}
}

func InitRouterPlayer(service playerService, router *gin.Engine) {
	playerApi := providePlayerApi(service)
	router.GET("/team/:team", playerApi.GetPlayerByTeam)
	router.GET("/team", playerApi.GetAllTeams)

}

func (api *PlayerApi) GetPlayerByTeam(requestContext *gin.Context) {

	param := requestContext.Param("team")

	if param == "first" {
		team := api.service.GetAllTeams()

		if team == nil {
			requestContext.Status(http.StatusNoContent)
			return
		}
		param = team[0].Name
	}

	team, err := api.service.GetTeamByTeamname(param)

	if err != nil {
		return
	}

	requestContext.JSON(http.StatusOK, mapper.TeamToDto(team))
}

func (api *PlayerApi) GetAllTeams(requestContext *gin.Context) {

	team := api.service.GetAllTeams()

	requestContext.JSON(http.StatusOK, mapper.TeamSliceToDto(team))
}
