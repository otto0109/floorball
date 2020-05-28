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

func (datasource *playerDatasource) GetPlayerByTeam(teamname string) entities.Team {
	var team entities.Team

	datasource.db.Where("name = ?", teamname).Preload("PlayerTeam").First(&team)

	players := make([]entities.Player, len(team.PlayerTeam))

	for index, playerTeam := range team.PlayerTeam {
		var player entities.Player
		datasource.db.Where("id = ?", playerTeam.Id).Find(&player)
		players[index] = player
	}

	team.Player = players

	return team
}

func (datasource *playerDatasource) CheckIfTeamExist(teamname string) bool {
	var team entities.Team

	datasource.db.Table("teams").Where("name = ?", teamname).First(&team)

	return team.Name == teamname
}
