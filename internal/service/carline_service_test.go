package service

import (
	"errors"
	"first-project/api/dto"
	"first-project/internal/config"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func GetCatalog(requestContext *gin.Context) {
	if requestContext.Query("tenant") == "fail" {
		requestContext.JSON(http.StatusBadRequest, errors.New("test").Error())
		return
	} else if requestContext.Query("tenant") == "badBody" {
		requestContext.JSON(http.StatusOK, "]")
		return
	}
	requestContext.JSON(http.StatusOK, getTestCatalogResponse())
}

func GetCarlines(requestContext *gin.Context) {
	if requestContext.Query("tenant") == "fail" {
		requestContext.JSON(http.StatusBadRequest, errors.New("test").Error())
		return
	} else if requestContext.Query("tenant") == "badBody" {
		requestContext.JSON(http.StatusOK, "]")
		return
	}
	requestContext.JSON(http.StatusOK, getCarlineResponse())
}

func serverMock() *httptest.Server {
	handler := gin.Default()
	handler.GET("/catalogue/overview", GetCatalog)
	handler.GET("/catalogue/carlines", GetCarlines)
	handler.GET("/catalogue/models")

	server := httptest.NewServer(handler)

	return server
}

func TestVicciService_GetAllCarlinesWithTenant(t *testing.T) {
	const tenant = "test"
	t.Run("Get all Carlines", func(t *testing.T) {
		server := serverMock()
		defer server.Close()
		vicciBaseConfig := config.VicciConfig{
			BaseURL:  server.URL,
			UserName: "test",
			Password: "test",
		}

		result, err := SetupCarlineService(vicciBaseConfig).GetAllCarlinesWithTenant(tenant)

		assert.NoError(t, err)
		assert.Equal(t, getCarlines(), result)
	})
	t.Run("Get all Carlines with error", func(t *testing.T) {
		server := serverMock()
		defer server.Close()
		vicciBaseConfig := config.VicciConfig{
			BaseURL:  server.URL,
			UserName: "test",
			Password: "test",
		}

		_, err := SetupCarlineService(vicciBaseConfig).GetAllCarlinesWithTenant("fail")

		assert.Error(t, err)
	})
	t.Run("Get all Carlines with error", func(t *testing.T) {
		server := serverMock()
		defer server.Close()
		vicciBaseConfig := config.VicciConfig{
			BaseURL:  server.URL,
			UserName: "test",
			Password: "test",
		}

		_, err := SetupCarlineService(vicciBaseConfig).GetAllCarlinesWithTenant("badBody")

		assert.Error(t, err)
	})
}

func TestVicciService_GetCatalogByTenantAndCarline(t *testing.T) {

	const tenant = "test"
	const carline = "test"
	t.Run("Get Catalog", func(t *testing.T) {
		server := serverMock()
		defer server.Close()
		vicciBaseConfig := config.VicciConfig{
			BaseURL:  server.URL,
			UserName: "test",
			Password: "test",
		}

		result, err := SetupCarlineService(vicciBaseConfig).GetCatalogByTenantAndCarline(tenant, carline)

		assert.Equal(t, getTestCatalog(), result)
		assert.NoError(t, err)
	})
	t.Run("Get Catalog with returned error", func(t *testing.T) {
		server := serverMock()
		defer server.Close()
		vicciBaseConfig := config.VicciConfig{
			BaseURL:  server.URL,
			UserName: "test",
			Password: "test",
		}

		_, err := SetupCarlineService(vicciBaseConfig).GetCatalogByTenantAndCarline("fail", carline)

		assert.Error(t, err)
	})

	t.Run("Get Catalog bad decode", func(t *testing.T) {
		server := serverMock()
		defer server.Close()
		vicciBaseConfig := config.VicciConfig{
			BaseURL:  server.URL,
			UserName: "test",
			Password: "test",
		}

		_, err := SetupCarlineService(vicciBaseConfig).GetCatalogByTenantAndCarline("badBody", carline)

		assert.Error(t, err)
	})
}

func getTestCatalogResponse() dto.VicciCarlineCatalogResult {
	return dto.VicciCarlineCatalogResult{
		Carline: dto.CarlineCatalog{
			Name:        "VW_GOLF",
			Code:        "123456",
			Salesgroups: nil,
		},
		Salesgroups: nil,
	}
}

func getTestCatalog() dto.CarlineCatalog {
	return dto.CarlineCatalog{
		Name:        "VW_GOLF",
		Code:        "123456",
		Salesgroups: nil,
	}
}

func getCarlineResponse() dto.VicciCarlineResult {
	return dto.VicciCarlineResult{
		Carlines: []dto.Carline{
			{
				Name: "VW",
				Code: "123456",
			},
			{
				Name: "VW",
				Code: "123456",
			},
		},
	}
}

func getCarlines() []dto.Carline {
	return []dto.Carline{
		{
			Name: "VW",
			Code: "123456",
		},
		{
			Name: "VW",
			Code: "123456",
		},
	}
}

type CarlineCatalogResponse struct {
	Name            string
	Code            string
	referenceModels []referenceModel
}

type referenceModel struct {
}
