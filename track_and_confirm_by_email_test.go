package go_usps

import (
	"reflect"
	"strings"
	"testing"
)

func TestUSPS_PTSEmail(t *testing.T) {
	username := "USPSUsername"
	password := "USPSPassword"

	successRespXmlStr := `
<?xml version="1.0"?>
<PTSEmailResult>
  <ResultText>Your request for all activity to-date will be processed within four hours. Any future activity   will be processed whenever there is new delivery related event activity.</ResultText>
  <ReturnCode>0</ReturnCode>
</PTSEmailResult>
`

	request := PTSEmailRequest{
		USERID: username,
	}

	request.TrackID = "XXXXXXXXXX1"

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
		request *PTSEmailRequest
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
			want:    "0",
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
			got, err := U.PTSEmail(tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("USPS.PTSEmail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.ReturnCode, tt.want) {
				t.Errorf("USPS.PTSEmail() = %v, want %v", got.ReturnCode, tt.want)
			}
		})
	}
}
