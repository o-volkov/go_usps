package go_usps

import "encoding/xml"

// eVS Cancel Request
// https://www.usps.com/business/web-tools-apis/evs-label-api.htm#_Toc487532701

type EVSCancelRequest struct {
	XMLName       xml.Name `xml:"eVSResponse"`
	USERID        string   `xml:"USERID,attr"`
	PASSWORD      string   `xml:"PASSWORD,attr,omitempty"`
	BarcodeNumber string   `xml:"BarcodeNumber"`
}

func (r *EVSCancelRequest) toHTTPRequestStr(bool) (string, error) {
	return createUSPSApiRequestStr("eVSCancel", r)
}

type EVSCancelResponse struct {
	XMLName       xml.Name `xml:"eVSCancelResponse"`
	BarcodeNumber string   `xml:"BarcodeNumber"`
	Status        string   `xml:"Status"`
	Reason        string   `xml:"Reason"`
}

func (U *USPS) EVSDomesticCancel(request *EVSCancelRequest) (EVSCancelResponse, error) {
	request.USERID = U.Username

	if U.Password != "" {
		request.PASSWORD = U.Password
	}

	result := new(EVSCancelResponse)
	err := U.Client.Execute(request, result)

	return *result, err
}
