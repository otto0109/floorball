package mapper

import (
	"floorball/api/dto"
	"floorball/internal/entities"
)

func TeamToDto(team entities.Team) dto.Team {
	var teamDto dto.Team

	teamDto.ID = team.ID
	teamDto.Name = team.Name
	teamDto.Player = slicePlayerToDto(team.Player)

	return teamDto
}
