package go_usps

import (
	"reflect"
	"strings"
	"testing"
)

func TestUSPS_MerchandiseReturnServiceLabels(t *testing.T) {
	username := "USPSUsername"
	password := "USPSPassword"

	successRespXmlStr := `
<?xml version="1.0"?>
<EMRSV4.0Response>
  <MerchandiseReturnLabel>SUkqAAgAAAASAP4ABAABAAAAAAAAAAABBAABAAAArgYAAAEBBA…
    <!--65255 skipped-->
  </MerchandiseReturnLabel>
  <CustomerAddress2>7 N Wilkes Barre Blvd</CustomerAddress2>
</EMRSV4.0Response>
`

	request := EMRSV40Request{
		USERID: username,
	}
	request.CustomerAddress2 = "7 N Wilkes Barre Blvd"

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
		request *EMRSV40Request
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
			want:    "7 N Wilkes Barre Blvd",
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
			got, err := U.MerchandiseReturnServiceLabels(tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("USPS.MerchandiseReturnServiceLabels() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.CustomerAddress2, tt.want) {
				t.Errorf("USPS.MerchandiseReturnServiceLabels() = %v, want %v", got.CustomerAddress2, tt.want)
			}
		})
	}
}
