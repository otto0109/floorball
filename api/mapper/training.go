package mapper

import (
	"floorball/api/dto"
	"floorball/internal/entities"
)

func trainingToDto(trainingEntity entities.Training) (trainingDto dto.Training) {
	trainingDto.ID = trainingEntity.ID
	trainingDto.End = trainingEntity.End
	trainingDto.Start = trainingEntity.Start
	trainingDto.Gym = trainingEntity.Gym
	trainingDto.GymMaps = trainingEntity.GymMaps
	trainingDto.Day = trainingEntity.Day

	return
}

func SliceTrainingToDto(trainingEntitySlice []entities.Training) []dto.Training {

	trainingDto := make([]dto.Training, len(trainingEntitySlice))

	for index, training := range trainingEntitySlice {
		trainingDto[index] = trainingToDto(training)
	}

	return trainingDto
}
