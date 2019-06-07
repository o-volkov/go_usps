package go_usps

import (
	"reflect"
	"strings"
	"testing"
)

func TestUSPS_EVSDomesticLabel(t *testing.T) {
	username := "USPSUsername"
	password := "USPSPassword"

	successRespXmlStr := `
<?xml version="1.0"?>
<eVSResponse>
 <BarcodeNumber>420063719270110101010XXXXXXXXX</BarcodeNumber>
  <LabelImage>SUkqAAgAAAASAP4ABAAB
    <!-- over 115000 suppressed -->
  </LabelImage>
  <ToName>TOM COLLINS</ToName>
  <ToFirm>XYZ CORP.</ToFirm>
  <ToAddress1>APT 1</ToAddress1>
<ToAddress2>970 DR MARTIN LUTHER KING JR BLVD</ToAddress2>
<ToAddress2Abbreviation>970 MLK BLVD</ToAddress2Abbreviation>
<ToCity>RIVIERA BEACH</ToCity>
<ToState>FL</ToState>
<ToZip5>33404</ToZip5>
<ToZip4>7400</ToZip4>
 <Postnet>063711844088</Postnet>
  <RDC>0007</RDC>
  <Postage>18.11</Postage>
  <Zone>00</Zone>
  <CarrierRoute>R###</CarrierRoute>
  <PermitHolderName>TEST - DO NOT MAIL</PermitHolderName>
  <InductionType>ePostage</InductionType>
  <LogMessage></LogMessage>
</eVSResponse>`

	request := EVSRequest{
		USERID:   username,
		PASSWORD: password,
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
		request *EVSRequest
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
			want:    "420063719270110101010XXXXXXXXX",
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
			got, err := U.EVSDomesticLabel(tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("USPS.EVSDomesticLabel() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.BarcodeNumber, tt.want) {
				t.Errorf("USPS.EVSDomesticLabel() = %v, want %v", got.BarcodeNumber, tt.want)
			}
		})
	}
}
