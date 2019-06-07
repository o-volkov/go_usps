package go_usps

import (
	"errors"
	"io/ioutil"
	"net/http"
)

type HttpClient interface {
	Get(url string) (resp *http.Response, err error)
}

type USPSHttpClient struct {
	USPSClient
	HttpClient HttpClient
	Production bool
}

func (c *USPSHttpClient) Execute(request USPSHTTPRequest, result interface{}) error {
	reqStr, err := request.toHTTPRequestStr(c.Production)
	if err != nil {
		return err
	}

	body, err := c.callRequest(reqStr)
	if err != nil {
		return err
	}
	if body == nil {
		return errors.New("error on request")
	}

	return parseUSPSXml(body, result)
}

func (c *USPSHttpClient) callRequest(requestURL string) ([]byte, error) {
	currentURL := ""
	if c.Production {
		currentURL += defaultHttpsProd
	} else {
		currentURL += defaultHttpsDev
	}
	currentURL += requestURL

	resp, err := c.HttpClient.Get(currentURL)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	return body, nil
}

func getHttpClient() HttpClient {
	return http.DefaultClient
}
