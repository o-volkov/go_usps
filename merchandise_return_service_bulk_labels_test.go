package go_usps

import (
	"reflect"
	"strings"
	"testing"
)

func TestUSPS_MerchandiseReturnServiceBulkLabels(t *testing.T) {
	username := "USPSUsername"
	password := "USPSPassword"

	successRespXmlStr := `
<?xml version="1.0"?>
<EMRSV4.0BulkResponse>
</EMRSV4.0BulkResponse>
`

	request := EMRSV40BulkRequest{
		USERID: username,
	}

	request.LabelCount = "3"

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
		request *EMRSV40BulkRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    EMRSV40BulkResponse
		wantErr bool
	}{
		{name: "Success flow",
			fields:  fields{Username: username, Password: password, Client: &TestClient{RequestResponseMap: requestResponseMap}},
			args:    args{request: &request},
			want:    EMRSV40BulkResponse{},
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
			got, err := U.MerchandiseReturnServiceBulkLabels(tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("USPS.MerchandiseReturnServiceBulkLabels() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("USPS.MerchandiseReturnServiceBulkLabels() = %v, want %v", got, tt.want)
			}
		})
	}
}
