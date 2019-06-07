package go_usps

import (
	"reflect"
	"strings"
	"testing"
)

func TestUSPS_ExpressMailServiceCommitments(t *testing.T) {
	username := "USPSUsername"
	password := "USPSPassword"

	successRespXmlStr := `
<?xml version="1.0"?>
<ExpressMailCommitmentResponse>
  <OriginZIP>90201</OriginZIP>
  <OriginCity>BELL GARDENS</OriginCity>
  <OriginState>CA</OriginState>
  <DestinationZIP>21114</DestinationZIP>
  <DestinationCity>CROFTON</DestinationCity>
  <DestinationState>MD</DestinationState>
  <Date>3-Jul-2014</Date>
  <Time>7:51AM</Time>
  <EffectiveAcceptanceDate>2014-07-03</EffectiveAcceptanceDate>
  <Commitment>
    <Name>2-Day</Name>
    <Time>3:00 PM</Time>
    <Sequence>A0215</Sequence>
    <Location>
      <ScheduledDeliveryDate>2014-07-05</ScheduledDeliveryDate>
      <CutOff>5:00 PM</CutOff>
      <Facility>POST OFFICE</Facility>
      <Street>7001 GARFIELD AVE</Street>
      <City>BELL GARDENS</City>
      <State>CA</State>
      <Zip>90201</Zip>
      <IsGuaranteed>1</IsGuaranteed>
    </Location>
    <Location>
      <ScheduledDeliveryDate>2014-07-05</ScheduledDeliveryDate>
      <CutOff>4:45 PM</CutOff>
      <Facility>POST OFFICE</Facility>
      <Street>4619 ELIZABETH ST</Street>
      <City>CUDAHY</City>
      <State>CA</State>
      <Zip>90201</Zip>
      <IsGuaranteed>1</IsGuaranteed>
    </Location>
    <Location>
      <ScheduledDeliveryDate>2014-07-05</ScheduledDeliveryDate>
      <CutOff>2:00 PM</CutOff>
      <Facility>POST OFFICE</Facility>
      <Street>5555 BANDINI BLVD</Street>
      <City>BELL GARDENS</City>
      <State>CA</State>
      <Zip>90201</Zip>
      <IsGuaranteed>1</IsGuaranteed>
    </Location>
  </Commitment>
  <Commitment>
    <Name>2-Day</Name>
    <Time>3:00 PM</Time>
    <Sequence>B0215</Sequence>
    <Location>
      <ScheduledDeliveryDate>2014-07-07</ScheduledDeliveryDate>
      <CutOff>5:00 PM</CutOff>
      <Facility>POST OFFICE</Facility>
      <Street>7001 GARFIELD AVE</Street>
      <City>BELL GARDENS</City>
      <State>CA</State>
      <Zip>90201</Zip>
      <IsGuaranteed>1</IsGuaranteed>
    </Location>
    <Location>
      <ScheduledDeliveryDate>2014-07-07</ScheduledDeliveryDate>
      <CutOff>4:45 PM</CutOff>
      <Facility>POST OFFICE</Facility>
      <Street>4619 ELIZABETH ST</Street>
      <City>CUDAHY</City>
      <State>CA</State>
      <Zip>90201</Zip>
      <IsGuaranteed>1</IsGuaranteed>
    </Location>
    <Location>
      <ScheduledDeliveryDate>2014-07-07</ScheduledDeliveryDate>
      <CutOff>2:00 PM</CutOff>
      <Facility>POST OFFICE</Facility>
      <Street>5555 BANDINI BLVD</Street>
      <City>BELL GARDENS</City>
      <State>CA</State>
      <Zip>90201</Zip>
      <IsGuaranteed>1</IsGuaranteed>
    </Location>
  </Commitment>
</ExpressMailCommitmentResponse>
`

	request := ExpressMailCommitmentRequest{
		USERID: username,
	}

	request.OriginZip = "90201"

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
		request *ExpressMailCommitmentRequest
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
			want:    "2-Day",
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
			got, err := U.ExpressMailServiceCommitments(tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("USPS.ExpressMailServiceCommitments() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.Commitment[0].Name, tt.want) {
				t.Errorf("USPS.ExpressMailServiceCommitments() = %v, want %v", got.Commitment[0].Name, tt.want)
			}
		})
	}
}
