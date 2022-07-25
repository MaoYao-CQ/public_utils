package services

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type ApiReturnJson struct {
	Status  bool        `json:"status"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Errors  []string    `json:"errors"`
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