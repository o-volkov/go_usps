package go_usps

// Package Pickup Cancel Web Tools
// https://www.usps.com/business/web-tools-apis/package-pickup-api.htm#_Toc450550258

type CarrierPickupCancelRequest struct {
	USERID             string `xml:"USERID,attr"`
	FirmName           string `xml:"FirmName,omitempty"`
	SuiteOrApt         string `xml:"SuiteOrApt"`
	Address2           string `xml:"Address2"`
	Urbanization       string `xml:"Urbanization"`
	City               string `xml:"City"`
	State              string `xml:"State"`
	ZIP5               string `xml:"ZIP5"`
	ZIP4               string `xml:"ZIP4"`
	ConfirmationNumber string `xml:"ConfirmationNumber"`
}

func (r *CarrierPickupCancelRequest) toHTTPRequestStr(bool) (string, error) {
	return createUSPSApiRequestStr("CarrierPickupCancel", r)
}

type CarrierPickupCancelResponse struct {
	FirmName           string `xml:"FirmName"`
	SuiteOrApt         string `xml:"SuiteOrApt"`
	Address2           string `xml:"Address2"`
	Urbanization       string `xml:"Urbanization"`
	City               string `xml:"City"`
	State              string `xml:"State"`
	ZIP5               string `xml:"ZIP5"`
	ZIP4               string `xml:"ZIP4"`
	ConfirmationNumber string `xml:"ConfirmationNumber"`
	Status             string `xml:"Status"`
}

func (U *USPS) PackagePickupCancel(request *CarrierPickupCancelRequest) (CarrierPickupCancelResponse, error) {
	request.USERID = U.Username

	result := new(CarrierPickupCancelResponse)
	err := U.Client.Execute(request, result)

	return *result, err
}
