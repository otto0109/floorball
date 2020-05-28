package service

import (
	"floorball/customError"
	"floorball/internal/entities"
)

type playerRepository interface {
	GetPlayerByTeam(teamname string) entities.Team
	CheckIfTeamExist(teamname string) bool
}

type playerService struct {
	repository playerRepository
}

func ProvidePlayerService(repository playerRepository) *playerService {
	return &playerService{repository: repository}
}

func (service *playerService) GetPlayerByTeam(teamName string) (entities.Team, error) {

	teamExist := service.repository.CheckIfTeamExist(teamName)

	if teamExist == false {
		return entities.Team{}, &customError.BadRequest{ErrorText: "team does not exist"}
	}

	return service.repository.GetPlayerByTeam(teamName), nil

}
