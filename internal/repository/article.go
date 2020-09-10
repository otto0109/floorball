package repository

import (
	"floorball/internal/entities"
	"github.com/jinzhu/gorm"
)

type articleDatasource struct {
	db *gorm.DB
}

func ProvideArticleRepository(db *gorm.DB) *articleDatasource {
	return &articleDatasource{db: db}
}

func (db *articleDatasource) GetArticleByTeamAndCategory(teamId int64, category string) (articleSlice []entities.Article) {

	if category == "" {
		db.db.Preload("ArticleTeam").Find(&articleSlice)
	} else {
		db.db.Where("category = ?", category).Preload("ArticleTeam").Find(&articleSlice)
	}
	if teamId != 0 {
		for index, article := range articleSlice {
			isArticleForTeam := false

			for _, articleTeam := range article.ArticleTeam {
				if articleTeam.TeamID == teamId {
					isArticleForTeam = true
				}
			}

			if !isArticleForTeam {
				copy(articleSlice[index:], articleSlice[index+1:])     // Shift a[i+1:] left one index.
				articleSlice[len(articleSlice)-1] = entities.Article{} // Erase last element (write zero value).
				articleSlice = articleSlice[:len(articleSlice)-1]
			}

		}
	}

	for index, article := range articleSlice {
		teams := make([]entities.Team, len(article.ArticleTeam))

		for index, articleTeam := range article.ArticleTeam {
			var team entities.Team

			db.db.Where("id = ?", articleTeam.Id).Preload("PlayerTeam").First(&team)

			players := make([]entities.Player, len(team.PlayerTeam))

			for index, playerTeam := range team.PlayerTeam {
				var player entities.Player
				db.db.Where("id = ?", playerTeam.Id).Find(&player)
				players[index] = player
			}
			team.Player = players

			teams[index] = team
		}

		articleSlice[index].Teams = teams
	}

	return
}
