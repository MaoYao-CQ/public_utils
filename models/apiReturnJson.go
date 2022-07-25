package models

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type ApiReturnJson struct {
	Status  bool        `json:"status"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Errors  []string    `json:"errors"`
	Data    interface{} `json:"data"`
}

type ApiReturnJsonDataServiceStatus struct {
	ServiceType   string                 `json:"service_type"`
	ServiceStatus string                 `json:"service_status"`
	Services      map[string]interface{} `json:"services"`
}

func GetApiReturn(resp *http.Response) (*ApiReturnJson, error) {

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	result := ApiReturnJson{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

type ListData struct {
	Models interface{} `json:"models"`
	Total  int64       `json:"total"`
}
