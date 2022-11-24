package utilHttp

import (
	"crypto/tls"
	"encoding/json"
	//"fmt"
	//godump "github.com/favframework/debug"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func Request(method string, remoteUrl string, params url.Values, headers map[string]string, timeout ...time.Duration) (resp *http.Response, err error) {
	method = strings.ToUpper(method)
	var httpTimeout time.Duration
	if len(timeout) > 0 {
		httpTimeout = timeout[0]
	} else {
		httpTimeout = 0
	}
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{
		Timeout:   httpTimeout,
		Transport: tr,
	}

	if "GET" == method || "DELETE" == method {
		urlParse, err := url.Parse(remoteUrl)
		if nil != err {
			return nil, err
		}
		query := urlParse.Query()
		for qk, _ := range params {
			query.Set(qk, params.Get(qk))
		}
		urlParse.RawQuery = query.Encode()

		remoteUrl = urlParse.String()
		params = url.Values{}
	}

	//fmt.Println(remoteUrl,params.Encode())

	req, err := http.NewRequest(method, remoteUrl, strings.NewReader(params.Encode()))

	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	for hk, hv := range headers {
		req.Header.Set(hk, hv)
	}
	//godump.Dump(req)
	return client.Do(req)
}

func RequestJson(v interface{}, method string, remoteUrl string, params url.Values, headers map[string]string, timeout ...time.Duration) error {
	resp, err := Request(method, remoteUrl, params, headers, timeout...)
	if err != nil {
		return err
	}
	//godump.Dump(resp)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &v)
	if err != nil {
		return err
	}
	return nil
}
