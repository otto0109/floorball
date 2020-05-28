package util

import (
	"floorball/internal/entities"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
)

func InitDB(connectionString string) *gorm.DB {
	db, err := gorm.Open("mssql", connectionString)
	if err != nil {
		panic(err)
	}

	db.LogMode(true)

	db.AutoMigrate(
		&entities.Player{},
		&entities.Team{},
		&entities.PlayerTeam{},
	)

	return db
}
