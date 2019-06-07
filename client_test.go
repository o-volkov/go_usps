package go_usps

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"reflect"
	"testing"
)

type TestHttpClient struct {
}

func (hc *TestHttpClient) Get(url string) (resp *http.Response, err error) {
	resp = new(http.Response)
	resp.Status = "200"
	resp.Body = ioutil.NopCloser(bytes.NewReader([]byte("<Test></Test>")))

	return resp, nil
}

type TestClient struct {
	RequestResponseMap map[string][]byte
}

func (c *TestClient) Execute(request USPSHTTPRequest, result interface{}) error {
	reqStr, err := request.toHTTPRequestStr(false)
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

func (c *TestClient) callRequest(requestURL string) ([]byte, error) {
	if c.RequestResponseMap[requestURL] != nil {
		return c.RequestResponseMap[requestURL], nil
	}

	return nil, nil
}

type testUSPSHTTPRequest struct {
}

func (r *testUSPSHTTPRequest) toHTTPRequestStr(bool) (string, error) {
	return "", nil
}

type testResult struct{}

func TestUSPSHttpClient_Execute(t *testing.T) {
	type fields struct {
		USPSClient USPSClient
		HttpClient HttpClient
		Production bool
	}
	type args struct {
		request USPSHTTPRequest
		result  interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{name: "Success flow",
			fields:  fields{USPSClient: new(TestClient), HttpClient: new(TestHttpClient), Production: false},
			args:    args{request: new(testUSPSHTTPRequest), result: new(testResult)},
			wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &USPSHttpClient{
				USPSClient: tt.fields.USPSClient,
				HttpClient: tt.fields.HttpClient,
				Production: tt.fields.Production,
			}
			if err := c.Execute(tt.args.request, tt.args.result); (err != nil) != tt.wantErr {
				t.Errorf("USPSHttpClient.Execute() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUSPSHttpClient_callRequest(t *testing.T) {
	type fields struct {
		USPSClient USPSClient
		HttpClient HttpClient
		Production bool
	}
	type args struct {
		requestURL string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		{name: "Success flow",
			fields: fields{USPSClient: new(TestClient), HttpClient: new(TestHttpClient), Production: false},
			args:   args{""},
			want:   []byte("<Test></Test>"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &USPSHttpClient{
				USPSClient: tt.fields.USPSClient,
				HttpClient: tt.fields.HttpClient,
				Production: tt.fields.Production,
			}
			got, err := c.callRequest(tt.args.requestURL)
			if (err != nil) != tt.wantErr {
				t.Errorf("USPSHttpClient.callRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("USPSHttpClient.callRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}
