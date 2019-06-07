package go_usps

import (
	"reflect"
	"strings"
	"testing"
)

func TestUSPS_HoldForPickupFacilityInformation(t *testing.T) {
	username := "USPSUsername"
	password := "USPSPassword"

	successRespXmlStr := `
<?xml version="1.0"?>
<HFPFacilityInfoResponse>
  <PickupCity/>
  <PickupState/>
  <PickupZIP>33952</PickupZIP>
  <PickupZIP4/>
  <Facility>
    <FacilityID>1438805</FacilityID>
    <FacilityName>PORT CHARLOTTE BRANCH</FacilityName>
    <FacilityAddress>3740 TAMIAMI TRL</FacilityAddress>
    <FacilityCity>PORT CHARLOTTE</FacilityCity>
    <FacilityState>FL</FacilityState>
    <FacilityZIP>33952</FacilityZIP>
    <FacilityZIP4>9998</FacilityZIP4>
    <Has10amCommitment>false</Has10amCommitment>
  </Facility>
  <Facility>
    <FacilityID>1378061</FacilityID>
    <FacilityName>PORT CHARLOTTE ANNEX</FacilityName>
    <FacilityAddress>18100 PAULSON DR</FacilityAddress>
    <FacilityCity>PORT CHARLOTTE</FacilityCity>
    <FacilityState>FL</FacilityState>
    <FacilityZIP>33954</FacilityZIP>
    <FacilityZIP4>9998</FacilityZIP4>
    <Has10amCommitment>false</Has10amCommitment>
   </Facility>
</HFPFacilityInfoResponse>
`

	request := HFPFacilityInfoRequest{
		USERID:   username,
		PASSWORD: password,
	}

	request.PickupZIP = "33952"

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
		request *HFPFacilityInfoRequest
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
			want:    "PORT CHARLOTTE BRANCH",
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
			got, err := U.HoldForPickupFacilityInformation(tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("USPS.HoldForPickupFacilityInformation() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.Facility[0].FacilityName, tt.want) {
				t.Errorf("USPS.HoldForPickupFacilityInformation() = %v, want %v", got.Facility[0].FacilityName, tt.want)
			}
		})
	}
}
