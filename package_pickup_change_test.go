package go_usps

import (
	"reflect"
	"strings"
	"testing"
)

func TestUSPS_PackagePickupChange(t *testing.T) {
	username := "USPSUsername"
	password := "USPSPassword"

	successRespXmlStr := `
<?xml version="1.0"?>
<CarrierPickupChangeResponse>
  <FirstName>John</FirstName>
  <LastName>Doe</LastName>
  <FirmName>ABC Corp.</FirmName>
  <SuiteOrApt>Suite 777</SuiteOrApt>
  <Address2>1390 Market Street</Address2>
  <Urbanization></Urbanization>
  <City>Houston</City>
  <State>TX</State>
  <ZIP5>77058</ZIP5>
  <ZIP4>1234</ZIP4>
  <Phone>(555) 555-1234</Phone>
  <Extension>201</Extension>
  <Package>
       <ServiceType>PriorityMailExpress</ServiceType>
       <Count>2</Count>
  </Package>
  <Package>
       <ServiceType>PriorityMail</ServiceType>
       <Count>1</Count>
  </Package>
  <EstimatedWeight>14</EstimatedWeight>
  <PackageLocation>Front Door</PackageLocation>
  <SpecialInstructions>Packages are behind the screen door.</SpecialInstructions>
  <ConfirmationNumber>ABC12345</ConfirmationNumber>
  <DayOfWeek>Monday</DayOfWeek>
  <Date>04/01/2004</Date>
  <Status>Update successful</Status>
</CarrierPickupChangeResponse>
`

	request := CarrierPickupChangeRequest{
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
		request *CarrierPickupChangeRequest
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
			want:    "Update successful",
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
			got, err := U.PackagePickupChange(tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("USPS.PackagePickupChange() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.Status, tt.want) {
				t.Errorf("USPS.PackagePickupChange() = %v, want %v", got.Status, tt.want)
			}
		})
	}
}
