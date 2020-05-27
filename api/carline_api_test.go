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

func (mock *CarlineServiceMock) GetAllCarlinesWithTenant(tenant string) ([]dto.Carline, error) {
	args := mock.Called(tenant)
	return args.Get(0).([]dto.Carline), args.Error(1)
}

func (mock *CarlineServiceMock) GetCatalogByTenantAndCarline(tenant string, carline string) (dto.CarlineCatalog, error) {
	args := mock.Called(tenant, carline)
	return args.Get(0).(dto.CarlineCatalog), args.Error(1)
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
		}, nil)

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
		}, errors.New("sample"))

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
			}, nil)

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
			}, errors.New("sample"))

		router := gin.Default()

		ProvideCarlineApi(router, &mockCarlines)
		result := util.PerformRequest(router, "GET", "http://localhost:8080/tenant/"+tenant+"/catalog?carline="+carline)

		assert.Equal(t, http.StatusInternalServerError, result.Code)
		mockCarlines.AssertNumberOfCalls(t, getCatalog, 1)
	})
}
