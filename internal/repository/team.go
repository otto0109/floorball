package repository

import (
	"floorball/internal/entities"
	"github.com/jinzhu/gorm"
)

type playerDatasource struct {
	db *gorm.DB
}

func ProvidePlayerRepository(db *gorm.DB) *playerDatasource {
	return &playerDatasource{db: db}
}

func (datasource *playerDatasource) GetTeamByTeamname(teamname string) entities.Team {
	var team entities.Team

	datasource.db.Where("name = ?", teamname).Preload("PlayerTeam").Preload("TeamTraining").First(&team)

	players := make([]entities.Player, len(team.PlayerTeam))

	for index, playerTeam := range team.PlayerTeam {
		var player entities.Player
		datasource.db.Where("id = ?", playerTeam.PlayerID).Find(&player)
		players[index] = player
	}

	trainings := make([]entities.Training, len(team.TeamTraining))

	for index, teamTraining := range team.TeamTraining {
		var training entities.Training
		datasource.db.Where("id = ?", teamTraining.TrainingID).Find(&training)
		trainings[index] = training
	}

	team.Player = players
	team.Training = trainings

	return team
}

func (datasource *playerDatasource) CheckIfTeamExist(teamname string) bool {
	var team entities.Team

	datasource.db.Table("teams").Where("name = ?", teamname).First(&team)

	return team.Name == teamname
}

func (datasource *playerDatasource) GetAllTeams() []entities.Team {
	var teamArray []entities.Team

	datasource.db.Table("teams").Find(&teamArray)

	return teamArray
}
