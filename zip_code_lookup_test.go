package go_usps

import (
	"errors"
	"reflect"
	"strings"
	"testing"
)

func TestUSPS_ZipCodeLookup(t *testing.T) {
	username := "USPSUsername"
	password := "USPSPassword"

	successRespXmlStr := `
<?xml version="1.0" encoding="UTF-8"?>
<ZipCodeLookupResponse>
  <Address ID="0">
    <FirmName>XYZ CORP.</FirmName>
    <Address2>6406 IVY LN</Address2>
    <City>GREENBELT</City>
    <State>MD</State>
    <Zip5>90013</Zip5>
    <Zip4>1441</Zip4>
  </Address>
  <Address ID="1">
    <FirmName>ABC COMPANY</FirmName>
    <Address1>Apt/Suite 2</Address1>
    <Address2>435 S MAIN ST</Address2>
    <City>LOS ANGELES</City>
    <State>CA</State>
    <Zip5>90013</Zip5>
    <Zip4>1310</Zip4>
  </Address>
</ZipCodeLookupResponse>
`

	failureRespXmlStr := `
<?xml version="1.0" encoding="UTF-8"?>
<Error>
  <Number>123456</Number>
  <Source></Source>
  <Description>error 123456</Description>
  <HelpFile></HelpFile>
  <HelpContext></HelpContext>
</Error>`
	request := ZipCodeLookupRequest{
		USERID:   username,
		PASSWORD: password,
	}

	addr := AddressZipCodeLookupRequest{
		Address2: "8 Wildwood Drive",
		City:     "Old Lyme",
		State:    "CT",
	}

	request.Address = append(request.Address, addr)

	requestStr, _ := request.toHTTPRequestStr(false)

	failureRequest := ZipCodeLookupRequest{
		USERID:   username,
		PASSWORD: password,
	}

	failureAddr := AddressZipCodeLookupRequest{
		Address2: "",
		City:     "",
		State:    "NY",
	}

	failureRequest.Address = append(failureRequest.Address, failureAddr)

	failureRequestStr, _ := failureRequest.toHTTPRequestStr(false)

	requestResponseMap := map[string][]byte{
		requestStr:        []byte(strings.TrimSuffix(successRespXmlStr, "\n")),
		failureRequestStr: []byte(strings.TrimSuffix(failureRespXmlStr, "\n")),
	}

	type fields struct {
		Username string
		Password string
		AppId    string
		Client   USPSClient
	}
	type args struct {
		request *ZipCodeLookupRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
		error   error
	}{
		{name: "Success flow",
			fields:  fields{Username: username, Password: password, Client: &TestClient{RequestResponseMap: requestResponseMap}},
			args:    args{request: &request},
			want:    "90013",
			wantErr: false},
		{name: "Failure parsing flow",
			fields:  fields{Username: username, Password: password, Client: &TestClient{RequestResponseMap: requestResponseMap}},
			args:    args{request: new(ZipCodeLookupRequest)},
			want:    "",
			wantErr: true,
			error:   errors.New("error on request")},
		{name: "Failure USPS flow",
			fields:  fields{Username: username, Password: password, Client: &TestClient{RequestResponseMap: requestResponseMap}},
			args:    args{request: &failureRequest},
			want:    "",
			wantErr: true,
			error:   errors.New("error 123456")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			U := &USPS{
				Username:   tt.fields.Username,
				Password:   tt.fields.Password,
				AppId:      tt.fields.AppId,
				Production: false,
				Client:     tt.fields.Client,
			}
			got, err := U.ZipCodeLookup(tt.args.request)

			if tt.wantErr == true {
				if err == nil {
					t.Errorf("USPS.ZipCodeLookup() error = %v, wantErr %v", nil, tt.wantErr)
				} else if tt.error != nil && tt.error.Error() != err.Error() {
					t.Errorf("USPS.ZipCodeLookup() error = %v, wantErr %v", err.Error(), tt.error.Error())
				}
			} else {
				if !reflect.DeepEqual(got.Address[0].Zip5, tt.want) {
					t.Errorf("USPS.ZipCodeLookup() = %v, want %v", got.Address[0].Zip5, tt.want)
				}
			}
		})
	}
}
