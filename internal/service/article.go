package service

import (
	"floorball/internal/entities"
)

type articleRepository interface {
	GetArticleByTeamAndCategory(teamId int64, category string) []entities.Article
}

type articleService struct {
	repository articleRepository
}

func ProvideArticleService(repository articleRepository) *articleService {
	return &articleService{repository: repository}
}

func (service *articleService) GetArticleByTeamAndCategory(teamId int64, category string) []entities.Article {
	return service.repository.GetArticleByTeamAndCategory(teamId, category)
}
