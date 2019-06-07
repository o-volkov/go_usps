package go_usps

import (
	"reflect"
	"strings"
	"testing"
)

func TestUSPS_PackagePickupAvailability(t *testing.T) {
	username := "USPSUsername"
	password := "USPSPassword"

	successRespXmlStr := `
<?xml version="1.0"?>
<CarrierPickupAvailabilityResponse>
  <FirmName>ABC Corp.</FirmName>
  <SuiteOrApt>Suite 777</SuiteOrApt>
  <Address2>1390 Market Street</Address2>
  <Urbanization></Urbanization>
  <City>Houston</City>
  <State>TX</State>
  <ZIP5>77058</ZIP5>
  <ZIP4>1234</ZIP4>
  <DayOfWeek>Monday</DayOfWeek>
  <Date>4/01/2004</Date>
  <CarrierRoute>C</CarrierRoute>
</CarrierPickupAvailabilityResponse>
`

	request := CarrierPickupAvailabilityRequest{
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
		request *CarrierPickupAvailabilityRequest
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
			want:    "ABC Corp.",
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
			got, err := U.PackagePickupAvailability(tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("USPS.PackagePickupAvailability() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.FirmName, tt.want) {
				t.Errorf("USPS.PackagePickupAvailability() = %v, want %v", got.FirmName, tt.want)
			}
		})
	}
}
