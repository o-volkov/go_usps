package go_usps

import (
	"reflect"
	"strings"
	"testing"
)

func TestUSPS_TrackField(t *testing.T) {
	username := "USPSUsername"
	password := "USPSPassword"

	successRespXmlStr := `
<?xml version="1.0"?>
<TrackResponse>
  <TrackInfo ID="XXXXXXXXXXXX1">
    <TrackSummary> Your item was delivered at 6:50 am on February 6 in BARTOW FL 33830.</TrackSummary>
    <TrackDetail>February 6 6:49 am NOTICE LEFT BARTOW FL 33830</TrackDetail>
    <TrackDetail>February 6 6:48 am ARRIVAL AT UNIT BARTOW FL 33830</TrackDetail>
    <TrackDetail>February 6 3:49 am ARRIVAL AT UNIT LAKELAND FL 33805</TrackDetail>
    <TrackDetail>February 5 7:28 pm ENROUTE 33699</TrackDetail>
    <TrackDetail>February 5 7:18 pm ACCEPT OR PICKUP 33699</TrackDetail>
  </TrackInfo>
  <TrackInfo ID="XXXXXXXXXXXX2">
    <TrackSummary> There is no record of that mail item. If it was mailed recently, It may not yet be tracked. Please try again later. </TrackSummary>
  </TrackInfo>
  <TrackInfo ID="XXXXXXXXXXXX3">
    <TrackSummary> That's not a valid number. Please check to make sure you entered it correctly.</TrackSummary>
  </TrackInfo>
</TrackResponse>
`

	request := TrackFieldRequest{
		USERID: username,
	}

	request.TrackID = append(request.TrackID, TrackIDTrackFieldRequest{ID: "XXXXXXXXXXXX1"})

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
		request *TrackFieldRequest
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
			want:    "XXXXXXXXXXXX2",
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
			got, err := U.TrackField(tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("USPS.TrackField() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.TrackInfo[1].ID, tt.want) {
				t.Errorf("USPS.TrackField() = %v, want %v", got.TrackInfo[1].ID, tt.want)
			}
		})
	}
}
