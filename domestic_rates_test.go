package go_usps

import (
	"reflect"
	"strings"
	"testing"
)

func TestUSPS_RateDomestic(t *testing.T) {
	username := "USPSUsername"
	password := "USPSPassword"

	successRespXmlStr := `
<?xml version="1.0"?>
<RateV4Response>
  <Package ID="1ST">
    <ZipOrigination>44106</ZipOrigination>
    <ZipDestination>20770</ZipDestination>
  </Package>
  <Package ID="2ND">
    <ZipOrigination>44107</ZipOrigination>
    <ZipDestination>20771</ZipDestination>
    <Pounds>1</Pounds>
    <Ounces>8</Ounces>
  </Package>
</RateV4Response>
`

	request := RateV4Request{
		USERID: username,
	}

	request.Package = append(request.Package, PackageRateV4Request{
		ID:             "1ST",
		ZipOrigination: "44106",
	}, PackageRateV4Request{
		ID:             "2ST",
		ZipOrigination: "44107",
	})

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
		request *RateV4Request
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
			want:    "44106",
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
			got, err := U.RateDomestic(tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("USPS.RateDomestic() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.Package[0].ZipOrigination, tt.want) {
				t.Errorf("USPS.RateDomestic() = %v, want %v", got, tt.want)
			}
		})
	}
}
