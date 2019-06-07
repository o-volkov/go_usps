package go_usps

import (
	"reflect"
	"strings"
	"testing"
)

func TestUSPS_ServiceDeliveryCalculatorGetLocations(t *testing.T) {
	username := "USPSUsername"
	password := "USPSPassword"

	successRespXmlStr := `
<?xml version="1.0"?>
<SDCGetLocationsResponse>
  <Release>2.0</Release>
  <MailClass>0</MailClass>
  <OriginZIP>70601</OriginZIP>
  <OriginCity>LAKE CHARLES</OriginCity>
  <OriginState>LA</OriginState>
  <DestZIP>98101</DestZIP>
  <DestCity>SEATTLE</DestCity>
  <DestState>WA</DestState>
  <AcceptDate>2014-07-01</AcceptDate>
  <AcceptTime>0900</AcceptTime>
  <Expedited>
    <EAD>2014-07-01</EAD>
      <Commitment>
      <MailClass>1</MailClass>
        <CommitmentName>1-Day</CommitmentName>
        <CommitmentTime>1030</CommitmentTime>
        <CommitmentSeq>A0110</CommitmentSeq>
        <Location>
          <SDD>2014-07-02</SDD>
          <COT>1700</COT>
          <FacType>EXPRESS MAIL COLLECTION BOX</FacType>
          <Street>604 PUJO ST</Street>
          <City>LAKE CHARLES</City>
          <State>LA</State>
          <ZIP>70601</ZIP>
          <IsGuaranteed>1</IsGuaranteed>
        </Location>
        <Location>
         <SDD>2014-07-02</SDD>
          <COT>1700</COT>
          <FacType>EXPRESS MAIL COLLECTION BOX</FacType>
          <Street>619 KIRBY ST</Street>
          <City>LAKE CHARLES</City>
          <State>LA</State>
          <ZIP>70601</ZIP>
          <IsGuaranteed>1</IsGuaranteed>
        </Location>
      </Commitment>
      <Commitment>
        <CommitmentName>1-Day</CommitmentName>
        <CommitmentTime>1500</CommitmentTime>
        <CommitmentSeq>A0115</CommitmentSeq>
        <Location>
          <SDD>2014-07-02</SDD>
          <COT>1700</COT>
          <FacType>EXPRESS MAIL COLLECTION BOX</FacType>
          <Street>604 PUJO ST</Street>
          <City>LAKE CHARLES</City>
          <State>LA</State>
          <ZIP>70601</ZIP>
          <IsGuaranteed>1</IsGuaranteed>
        </Location>
        <Location>
          <SDD>2014-07-02</SDD>
          <COT>1700</COT>
          <FacType>EXPRESS MAIL COLLECTION BOX</FacType>
          <Street>619 KIRBY ST</Street>
          <City>LAKE CHARLES</City>
          <State>LA</State>
          <ZIP>70601</ZIP>
          <IsGuaranteed>1</IsGuaranteed>
        </Location>
      </Commitment>
  </Expedited>
  <NonExpedited>
    <MailClass>3</MailClass>
    <NonExpeditedDestType>1</NonExpeditedDestType>
    <EAD>2014-07-01</EAD>
    <COT>1700</COT>
    <SvcStdMsg>3 Days</SvcStdMsg>
    <SvcStdDays>3</SvcStdDays>
    <SchedDlvryDate>2014-07-05</SchedDlvryDate>
    <NonExpeditedExceptions>
    </NonExpeditedExceptions>
  </NonExpedited>
  <NonExpedited>
    <MailClass>6</MailClass>
    <NonExpeditedDestType>1</NonExpeditedDestType>
    <EAD>2014-07-01</EAD>
    <COT>1700</COT>
    <SvcStdMsg>7 Days</SvcStdMsg>
    <SvcStdDays>7</SvcStdDays>
    <SchedDlvryDate>2014-07-08</SchedDlvryDate>
  </NonExpedited>
</SDCGetLocationsResponse>
`

	request := SDCGetLocationsRequest{}

	request.OriginZIP = "70601"
	request.DestinationZIP = "21817"

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
		request *SDCGetLocationsRequest
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
			want:    "LAKE CHARLES",
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
			got, err := U.ServiceDeliveryCalculatorGetLocations(tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("USPS.ServiceDeliveryCalculatorGetLocations() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.OriginCity, tt.want) {
				t.Errorf("USPS.ServiceDeliveryCalculatorGetLocations() = %v, want %v", got.OriginCity, tt.want)
			}
		})
	}
}
