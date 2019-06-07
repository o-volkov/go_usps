package go_usps

import (
	"reflect"
	"strings"
	"testing"
)

func TestUSPS_InternationalRates(t *testing.T) {
	username := "USPSUsername"
	password := "USPSPassword"

	successRespXmlStr := `
<?xml version="1.0" encoding="UTF-8"?>
<IntlRateV2Response>
  <Package ID="1ST">
    <Prohibitions>Coins; bank notes; currency notes (paper money);...</Prohibitions>
    <Restrictions>Jewelry is permitted only when sent as an insured parcel using Priority Mail International service...</Restrictions>
    <Observations>Duty may be levied on catalogs, price lists, circulars...</Observations>
    <CustomsForms>First-Class Mail International items and Priority Mail International Flat Rate Envelopes...</CustomsForms>
    <ExpressMail>Country Code: AU Reciprocal Service Name: Express Post Required Customs Form/Endorsement...</ExpressMail>
    <AreasServed>Please reference Express Mail for Areas Served.</AreasServed>
    <AdditionalRestrictions>No Additional Restrictions Data found.</AdditionalRestrictions>
    <Service ID="12">
      <Pounds>15.12345678</Pounds>
      <Ounces>0</Ounces>
      <Machinable>True</Machinable>
      <MailType>Package</MailType>
      <Container>RECTANGULAR</Container>
      <Size>LARGE</Size>
      <Width>10</Width>
      <Length>15</Length>
      <Height>10</Height>
      <Girth>0</Girth>
      <Country>AUSTRALIA</Country>
      <Postage>211.50</Postage>
      <ExtraServices>
        <ExtraService>
          <ServiceID>1</ServiceID>
          <ServiceName>Insurance</ServiceName>
          <Available>True</Available>
          <Price>2.00</Price>
          <DeclaredValueRequired>True</DeclaredValueRequired>
        </ExtraService>
      </ExtraServices>
      <ValueOfContents>200.00</ValueOfContents>
      <MaxWeight>70</MaxWeight>
      <GXGLocations>
        <PostOffice>
          <Name>WILKES BARRE PDC</Name>
          <Address>300 S MAIN ST</Address>
          <City>WILKES BARRE</City>
          <State>PA</State>
          <ZipCode>18701</ZipCode>
          <RetailGXGCutOffTime>5:00 PM</RetailGXGCutOffTime>
          <SaturDayCutOffTime>2:00 PM</SaturDayCutOffTime>
        </PostOffice>
      </GXGLocations>
    </Service>
  </Package>
</IntlRateV2Response>
`

	request := IntlRateV2Request{
		USERID: username,
	}

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
		request *IntlRateV2Request
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
			want:    "1ST",
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
			got, err := U.InternationalRates(tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("USPS.InternationalRates() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.Package[0].ID, tt.want) {
				t.Errorf("USPS.InternationalRates() = %v, want %v", got.Package[0].ID, tt.want)
			}
		})
	}
}
