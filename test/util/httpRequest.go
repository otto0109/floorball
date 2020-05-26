package util

import (
	"net/http"
	"net/http/httptest"
)

func PerformRequest(handler http.Handler, method, path string) *httptest.ResponseRecorder {

	request, _ := http.NewRequest(method, path, nil)

	recorder := httptest.NewRecorder()
	handler.ServeHTTP(recorder, request)
	return recorder
}
