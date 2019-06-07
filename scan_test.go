package go_usps

import (
	"reflect"
	"strings"
	"testing"
)

func TestUSPS_Scan(t *testing.T) {
	username := "USPSUsername"
	password := "USPSPassword"

	successRespXmlStr := `
<?xml version="1.0"?>
<SCANResponse>
 <SCANFormNumber>CS201009141406094670660949S</SCANFormNumber>
 <SCANFormImage>SUkqAAgAAAASAP4ABAABAAAAAAAAAAABBAABAAAArgYAAAEBBA…
 <!--65255 skipped-->
 </SCANFormImage>
</SCANResponse>
`

	request := SCANRequest{}
	request.USERID = username
	request.FromName = "John Doe"
	request.FromFirm = "United States Postal Service"
	request.FromAddress1 = ""
	request.FromAddress2 = "475 L’Enfant Plaza SW, Room 1546"
	request.FromCity = "Washington"
	request.FromState = "DC"
	request.FromZip5 = "20260"
	request.FromZip4 = "1234"
	request.MailDate = "20120719"
	request.MailTime = "080501"
	request.EntryFacility = "20260"
	request.ImageType = ImageTypeTif
	request.CustomerRefNo = "123XYZ"
	request.Option.Form = Form5630
	request.Shipment.PackageDetail.PkgBarcode = "LJ019027504US"

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
		request *SCANRequest
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
			want:    "CS201009141406094670660949S",
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

			got, err := U.Scan(tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("USPS.Scan() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.SCANFormNumber, tt.want) {
				t.Errorf("USPS.Scan() = %v, want %v", got.SCANFormNumber, tt.want)
			}
		})
	}
}
