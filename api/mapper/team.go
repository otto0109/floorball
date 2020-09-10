package mapper

import (
	"floorball/api/dto"
	"floorball/internal/entities"
)

func TeamToDto(team entities.Team) dto.Team {
	var teamDto dto.Team

	teamDto.ID = team.ID
	teamDto.Name = team.Name
	teamDto.TeamFoto = team.TeamFoto
	teamDto.Player = slicePlayerToDto(team.Player)
	teamDto.Training = SliceTrainingToDto(team.Training)

	return teamDto
}

func TeamSliceToDto(teamSlice []entities.Team) []dto.Team {
	teamDtoSlice := make([]dto.Team, len(teamSlice))

	for index, team := range teamSlice {
		teamDtoSlice[index] = TeamToDto(team)
	}

	return teamDtoSlice
}
