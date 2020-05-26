package api

import (
	"errors"
	"first-project/api/dto"
	"first-project/test/util"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/stretchr/testify/mock"
	"net/http"
	"testing"
)

type CarlineServiceMock struct {
	mock.Mock
}

func (mock *CarlineServiceMock) GetAllCarlinesWithTenant(tenant string) ([]dto.Carline, ResponseError) {
	args := mock.Called(tenant)
	return args.Get(0).([]dto.Carline), args.Get(1).(ResponseError)
}

func (mock *CarlineServiceMock) GetCatalogByTenantAndCarline(tenant string, carline string) (dto.CarlineCatalog, ResponseError) {
	args := mock.Called(tenant, carline)
	return args.Get(0).(dto.CarlineCatalog), args.Get(1).(ResponseError)
}

func TestCarlineApi_GetAllCarlines(t *testing.T) {
	const getAllCarline = "GetAllCarlinesWithTenant"
	const tenant = "okapi-vw-es-es"

	t.Run("Get all Carlines", func(t *testing.T) {
		mockCarlines := CarlineServiceMock{}
		mockCarlines.On(getAllCarline, tenant).Return([]dto.Carline{
			{
				Name: "VW Golf",
				Code: "879613",
			},
		}, ResponseError{
			Code:  200,
			Error: nil,
		})

		router := gin.Default()

		ProvideCarlineApi(router, &mockCarlines)
		result := util.PerformRequest(router, "GET", "http://localhost:8080/tenant/"+tenant+"/carlines")
		expected := `[{"Name":"VW Golf","Code":"879613"}]`

		assert.Equal(t, http.StatusOK, result.Code)
		assert.Equal(t, expected, result.Body.String())
		mockCarlines.AssertNumberOfCalls(t, getAllCarline, 1)
	})

	t.Run("Get all Carlines with returned error", func(t *testing.T) {
		mockCarlines := CarlineServiceMock{}
		mockCarlines.On(getAllCarline, tenant).Return([]dto.Carline{
			{
				Name: "VW Golf",
				Code: "879613",
			},
		}, ResponseError{
			Code:  500,
			Error: errors.New("testError"),
		})

		router := gin.Default()

		ProvideCarlineApi(router, &mockCarlines)
		result := util.PerformRequest(router, "GET", "http://localhost:8080/tenant/"+tenant+"/carlines")

		assert.Equal(t, http.StatusInternalServerError, result.Code)
		mockCarlines.AssertNumberOfCalls(t, getAllCarline, 1)
	})
}

func TestCarlineApi_GetCatalog(t *testing.T) {
	const getCatalog = "GetCatalogByTenantAndCarline"
	const tenant = "okapi-vw-es-es"
	const carline = "799863"

	t.Run("Get Catalog", func(t *testing.T) {
		mockCarlines := CarlineServiceMock{}
		mockCarlines.On(getCatalog, tenant, carline).Return(
			dto.CarlineCatalog{
				Name:        "VW Golf",
				Code:        carline,
				Salesgroups: nil,
			},
			ResponseError{
				Code:  200,
				Error: nil,
			})

		router := gin.Default()

		ProvideCarlineApi(router, &mockCarlines)
		result := util.PerformRequest(router, "GET", "http://localhost:8080/tenant/"+tenant+"/catalog?carline="+carline)

		expected := `{"Name":"VW Golf","Code":"799863","Salesgroups":null}`
		assert.Equal(t, http.StatusOK, result.Code)
		assert.Equal(t, expected, result.Body.String())
		mockCarlines.AssertNumberOfCalls(t, getCatalog, 1)
	})

	t.Run("Get Catalog", func(t *testing.T) {
		mockCarlines := CarlineServiceMock{}
		mockCarlines.On(getCatalog, tenant, carline).Return(
			dto.CarlineCatalog{
				Name:        "VW Golf",
				Code:        carline,
				Salesgroups: nil,
			},
			ResponseError{
				Code:  500,
				Error: errors.New("test error"),
			})

		router := gin.Default()

		ProvideCarlineApi(router, &mockCarlines)
		result := util.PerformRequest(router, "GET", "http://localhost:8080/tenant/"+tenant+"/catalog?carline="+carline)

		assert.Equal(t, http.StatusInternalServerError, result.Code)
		mockCarlines.AssertNumberOfCalls(t, getCatalog, 1)
	})
}
