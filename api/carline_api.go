package api

import (
	"first-project/api/dto"
	"github.com/gin-gonic/gin"
	"net/http"
)

type carlineServiceImpl interface {
	GetAllCarlinesWithTenant(tenant string) ([]dto.Carline, ResponseError)
	GetCatalogByTenantAndCarline(tenant string, carline string) (dto.CarlineCatalog, ResponseError)
}

type carlineApi struct {
	carlineService carlineServiceImpl
}

func ProvideCarlineApi(router *gin.Engine, service carlineServiceImpl) {
	api := &carlineApi{carlineService: service}
	router.GET("/tenant/:vicci_tenant/carlines", api.GetAllCarlines)
	router.GET("/tenant/:vicci_tenant/catalog", api.GetCatalog)
}

func (api *carlineApi) GetAllCarlines(requestContext *gin.Context) {

	tenant := requestContext.Param("vicci_tenant")

	carlines, err := api.carlineService.GetAllCarlinesWithTenant(tenant)

	if err.Code != http.StatusOK {
		requestContext.JSON(err.Code, err.Error.Error())
		return
	}
	requestContext.JSON(http.StatusOK, carlines)

}

func (api *carlineApi) GetCatalog(requestContext *gin.Context) {

	tenant := requestContext.Param("vicci_tenant")
	carline := requestContext.Query("carline")

	carlines, err := api.carlineService.GetCatalogByTenantAndCarline(tenant, carline)

	if err.Code != http.StatusOK {
		requestContext.JSON(err.Code, err.Error.Error())
		return
	}
	requestContext.JSON(http.StatusOK, carlines)
}
