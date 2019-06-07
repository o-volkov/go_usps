package go_usps

import (
	"reflect"
	"strings"
	"testing"
)

func TestUSPS_AddressVerification(t *testing.T) {
	username := "USPSUsername"
	password := "USPSPassword"
	appId := "USPSAppId"

	successRespXmlStr := `
<?xml version="1.0"?>
<AddressValidateResponse>
  <Address ID="0">
    <Address2>29851 AVENTURA STE K</Address2>
    <City>RANCHO SANTA MARGARITA</City>
    <State>CA</State>
    <Zip5>92688</Zip5>
    <Zip4>2014</Zip4>
  </Address>
</AddressValidateResponse>
`

	request := AddressValidateRequest{
		USERID:   username,
		PASSWORD: password,
		APPID:    appId,
	}

	request.Address.Address2 = "29851 AVENTURA STE K"
	request.Address.City = ""
	request.Address.State = "CA"
	request.Address.Zip5 = "92688"
	request.Address.Zip4 = ""

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
		request *AddressValidateRequest
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
			want:    request.Address.Address2,
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
			got, err := U.AddressVerification(tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("USPS.AddressVerification() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.Address.Address2, tt.want) {
				t.Errorf("USPS.AddressVerification() got = %v, want %v", got.Address.Address2, tt.want)
			}
		})
	}
}
