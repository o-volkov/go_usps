package go_usps

import (
	"reflect"
	"strings"
	"testing"
)

func TestUSPS_EVSDomesticCancel(t *testing.T) {
	username := "USPSUsername"
	password := "USPSPassword"

	successRespXmlStr := `
<?xml version="1.0"?>
<eVSCancelResponse>
  <BarcodeNumber>420902109411202901089817001111</BarcodeNumber>
  <Status>Cancelled</Status>
  <Reason>Order Cancelled Successfully</Reason>
</eVSCancelResponse>
`

	request := EVSCancelRequest{
		USERID:   username,
		PASSWORD: password,
	}

	request.BarcodeNumber = "420902109411202901089817001111"

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
		request *EVSCancelRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{name: "Success flow",
			fields:  fields{Username: username, Password: password, Client: &TestClient{RequestResponseMap: requestResponseMap}},
			args:    args{request: &request},
			want:    "420902109411202901089817001111",
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
			got, err := U.EVSDomesticCancel(tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("USPS.EVSDomesticCancel() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.BarcodeNumber, tt.want) {
				t.Errorf("USPS.EVSDomesticCancel() = %v, want %v", got.BarcodeNumber, tt.want)
			}
		})
	}
}
