package api

import (
	"floorball/api/mapper"
	"floorball/internal/entities"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type articleService interface {
	GetArticleByTeamAndCategory(teamId int64, category string) []entities.Article
}

type ArticleApi struct {
	service articleService
}

func provideArticleApi(service articleService) *ArticleApi {
	return &ArticleApi{service: service}
}

func InitRouterArticle(service articleService, router *gin.Engine) {
	articleApi := provideArticleApi(service)
	router.GET("/articles", articleApi.GetArticles)
}

func (api *ArticleApi) GetArticles(requestContext *gin.Context) {
	var teamId int64

	teamId, _ = strconv.ParseInt(requestContext.Query("teamId"), 0, 64)

	category := requestContext.Query("category")

	article := api.service.GetArticleByTeamAndCategory(teamId, category)

	requestContext.JSON(http.StatusOK, mapper.ArticleSliceToDto(article))
}
