package mapper

import (
	"floorball/api/dto"
	"floorball/internal/entities"
)

func ArticleToDto(article entities.Article) (articleDto dto.Article) {

	articleDto.ID = article.ID
	articleDto.Thumbnail = article.Thumbnail
	articleDto.Article = article.Article
	articleDto.Title = article.Title
	articleDto.Teams = TeamSliceToDto(article.Teams)
	articleDto.Category = article.Category

	return articleDto
}

func ArticleSliceToDto(articleSlice []entities.Article) []dto.Article {
	articleDtoSlice := make([]dto.Article, len(articleSlice))

	for index, article := range articleSlice {
		articleDtoSlice[index] = ArticleToDto(article)
	}

	return articleDtoSlice
}
