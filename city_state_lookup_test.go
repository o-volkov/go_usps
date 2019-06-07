package go_usps

import (
	"reflect"
	"strings"
	"testing"
)

func TestUSPS_CityStateLookup(t *testing.T) {
	username := "USPSUsername"
	password := "USPSPassword"
	appId := "USPSAppId"

	successRespXmlStr := `
<?xml version="1.0"?>
<CityStateLookupResponse>
  <ZipCode ID="0">
    <Zip5>20024</Zip5>
    <City>BEVERLY HILLS</City>
    <State>CA</State>
    <FinanceNumber></FinanceNumber>
    <ClassificationCode></ClassificationCode>
  </ZipCode>
  <ZipCode ID="1">
    <Zip5>20770</Zip5>
    <City>GREENBELT</City>
    <State>MD</State>
    <FinanceNumber></FinanceNumber>
    <ClassificationCode></ClassificationCode>
  </ZipCode>
</CityStateLookupResponse>
`

	request := CityStateLookupRequest{
		USERID:   username,
		PASSWORD: password,
		APPID:    appId,
	}

	request.ZipCode.Zip5 = "20024"

	rStr, _ := request.toHTTPRequestStr(false)
	requestResponseMap := map[string][]byte{
		rStr: []byte(strings.TrimSuffix(successRespXmlStr, "\n")),
	}

	type fields struct {
		Username   string
		Password   string
		AppId      string
		Production bool
		Client     USPSClient
	}
	type args struct {
		request *CityStateLookupRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{name: "Success flow",
			fields:  fields{Username: username, Password: password, AppId: appId, Client: &TestClient{RequestResponseMap: requestResponseMap}},
			args:    args{request: &request},
			want:    request.ZipCode.Zip5,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			U := &USPS{
				Username:   tt.fields.Username,
				Password:   tt.fields.Password,
				AppId:      tt.fields.AppId,
				Production: tt.fields.Production,
				Client:     tt.fields.Client,
			}
			got, err := U.CityStateLookup(tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("USPS.CityStateLookup() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.ZipCode[0].Zip5, tt.want) {
				t.Errorf("USPS.CityStateLookup() = %v, want %v", got.ZipCode[0].Zip5, tt.want)
			}
		})
	}
}
