package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"first-project/api"
	"first-project/api/dto"
	"first-project/internal/config"
	"net/http"
)

const (
	allCarlinesUrl = "/catalogue/carlines?tenant="
	catalogUrl     = "/catalogue/salesgroups?tenant="
	flags          = "&fetchPrices=false&fetchMedia=false&addErrorMps=false"
	carlineKeyFlag = "&carlineKey="
)

type vicciService struct {
	vicciBaseConfig config.VicciConfig
}

func SetupCarlineService(baseConfig config.VicciConfig) *vicciService {
	return &vicciService{vicciBaseConfig: baseConfig}
}

func (service *vicciService) GetCatalogByTenantAndCarline(tenant string, carline string) (catalog dto.CarlineCatalog, err api.ResponseError) {

	response, err := performVicciRequest(service.vicciBaseConfig, catalogUrl+tenant+carlineKeyFlag+carline+flags)

	if err.Error != nil {
		return
	}
	decodeError := json.Unmarshal([]byte(bodyToString(response)), &catalog)

	if decodeError != nil {
		err = api.ResponseError{
			Code:  http.StatusInternalServerError,
			Error: decodeError,
		}
	}

	return
}

func (service *vicciService) GetAllCarlinesWithTenant(tenant string) (carlines []dto.Carline, err api.ResponseError) {
	var result dto.VicciCarlineResult

	response, err := performVicciRequest(service.vicciBaseConfig, allCarlinesUrl+tenant+flags)

	if err.Error != nil {
		return
	}

	decodeError := json.Unmarshal([]byte(bodyToString(response)), &result)

	if decodeError != nil {
		err = api.ResponseError{
			Code:  http.StatusInternalServerError,
			Error: decodeError,
		}
	}

	carlines = result.Carlines

	return
}

func performVicciRequest(vicciBaseConfig config.VicciConfig, requestUrl string) (*http.Response, api.ResponseError) {

	client := http.Client{
		Timeout: 0,
	}

	req, err := http.NewRequest("GET", vicciBaseConfig.BaseURL+requestUrl, nil)

	if err != nil {
		return nil, api.ResponseError{
			Code:  http.StatusBadRequest,
			Error: err,
		}
	}

	req.Header.Set("Authorization", "Basic "+basicAuth(vicciBaseConfig.UserName, vicciBaseConfig.Password))

	response, err := client.Do(req)

	if err != nil {
		return nil, api.ResponseError{
			Code:  http.StatusBadRequest,
			Error: err,
		}
	} else if response.StatusCode != http.StatusOK {
		newString := bodyToString(response)
		return nil, api.ResponseError{
			Code:  response.StatusCode,
			Error: errors.New(newString),
		}
	}
	return response, api.ResponseError{
		Code:  http.StatusOK,
		Error: nil,
	}
}

func bodyToString(response *http.Response) string {
	buf := new(bytes.Buffer)
	buf.ReadFrom(response.Body)
	newString := buf.String()
	return newString
}
