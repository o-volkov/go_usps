package go_usps

import (
	"reflect"
	"strings"
	"testing"
)

func TestUSPS_PriorityMail(t *testing.T) {
	username := "USPSUsername"
	password := "USPSPassword"

	successRespXmlStr := `
<?xml version="1.0"?>
<PriorityMailResponse>
  <OriginZip>90201</OriginZip>
  <DestinationZip>21114</DestinationZip>
  <Days>4</Days>
</PriorityMailResponse>
`

	request := PriorityMailRequest{
		USERID: username,
	}

	request.OriginZip = "90201"
	request.DestinationZip = "21114"

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
		request *PriorityMailRequest
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
			want:    "4",
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
			got, err := U.PriorityMail(tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("USPS.PriorityMail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.Days, tt.want) {
				t.Errorf("USPS.PriorityMail() = %v, want %v", got.Days, tt.want)
			}
		})
	}
}
