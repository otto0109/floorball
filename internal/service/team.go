package service

import (
	"floorball/customError"
	"floorball/internal/entities"
)

type playerRepository interface {
	GetTeamByTeamname(teamname string) entities.Team
	CheckIfTeamExist(teamname string) bool
	GetAllTeams() []entities.Team
}

type playerService struct {
	repository playerRepository
}

func ProvidePlayerService(repository playerRepository) *playerService {
	return &playerService{repository: repository}
}

func (service *playerService) GetTeamByTeamname(teamName string) (entities.Team, error) {

	teamExist := service.repository.CheckIfTeamExist(teamName)

	if teamExist == false {
		return entities.Team{}, &customError.BadRequest{ErrorText: "team does not exist"}
	}

	return service.repository.GetTeamByTeamname(teamName), nil

}

func (service *playerService) GetAllTeams() []entities.Team {
	return service.repository.GetAllTeams()
}
