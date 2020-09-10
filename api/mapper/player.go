package mapper

import (
	"floorball/api/dto"
	"floorball/internal/entities"
)

func playerToDto(playerEntity entities.Player) (playerDto dto.Player) {
	playerDto.ID = playerEntity.ID
	playerDto.Name = playerEntity.Name
	playerDto.Surname = playerEntity.Surname
	playerDto.Position = playerEntity.Position
	playerDto.Qoute = playerEntity.Qoute
	playerDto.Instagram = playerEntity.Instagram
	playerDto.PictureUrl = playerEntity.PictureUrl

	return playerDto
}

func slicePlayerToDto(playerEntitySlice []entities.Player) []dto.Player {

	playerDto := make([]dto.Player, len(playerEntitySlice))

	for index, player := range playerEntitySlice {
		playerDto[index] = playerToDto(player)
	}

	return playerDto
}
