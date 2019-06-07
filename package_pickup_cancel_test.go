package go_usps

import (
	"reflect"
	"strings"
	"testing"
)

func TestUSPS_PackagePickupCancel(t *testing.T) {
	username := "USPSUsername"
	password := "USPSPassword"

	successRespXmlStr := `
<?xml version="1.0"?>
<CarrierPickupCancelResponse>
  <FirmName>ABC Corp.</FirmName>
  <SuiteOrApt>Suite 777</SuiteOrApt>
  <Address2>1390 Market Street</Address2>
  <Urbanization></Urbanization>
  <City>Houston</City>
  <State>TX</State>
  <ZIP5>77058</ZIP5>
  <ZIP4>1234</ZIP4>
  <ConfirmationNumber>ABC12345</ConfirmationNumber>
  <Status>Your pickup request was cancelled.</Status>
</CarrierPickupCancelResponse>
`

	request := CarrierPickupCancelRequest{
		USERID: username,
	}

	request.FirmName = "ABC Corp."

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
		request *CarrierPickupCancelRequest
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
			want:    "Your pickup request was cancelled.",
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
			got, err := U.PackagePickupCancel(tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("USPS.PackagePickupCancel() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.Status, tt.want) {
				t.Errorf("USPS.PackagePickupCancel() = %v, want %v", got.Status, tt.want)
			}
		})
	}
}
