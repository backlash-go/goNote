package factory_pattern

import (
	"net/http"
	"net/http/httptest"
)

type Doer interface {
	Do(req *http.Request) (*http.Response, error)
}

func NewHTTPClient()  Doer {
	return &http.Client{}
}


type mockHTTPClient struct{}

func (m mockHTTPClient) Do(req *http.Request) (*http.Response, error) {
	res := httptest.NewRecorder()
	return res.Result(), nil
}


func NewMockHTTPClient() Doer { return &mockHTTPClient{}}









