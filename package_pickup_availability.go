package go_usps

// Package Pickup Availability API
// https://www.usps.com/business/web-tools-apis/package-pickup-api.htm#_Toc450550246

type CarrierPickupAvailabilityRequest struct {
	USERID       string `xml:"USERID,attr"`
	FirmName     string `xml:"FirmName,omitempty"`
	SuiteOrApt   string `xml:"SuiteOrApt"`
	Address2     string `xml:"Address2"`
	Urbanization string `xml:"Urbanization"`
	City         string `xml:"City"`
	State        string `xml:"State"`
	ZIP5         string `xml:"ZIP5"`
	ZIP4         string `xml:"ZIP4"`
	Date         string `xml:"Date,omitempty"`
}

func (r *CarrierPickupAvailabilityRequest) toHTTPRequestStr(bool) (string, error) {
	return createUSPSApiRequestStr("CarrierPickupAvailability", r)
}

type CarrierPickupAvailabilityResponse struct {
	FirmName     string `xml:"FirmName,omitempty"`
	SuiteOrApt   string `xml:"SuiteOrApt"`
	Address2     string `xml:"Address2"`
	Urbanization string `xml:"Urbanization"`
	City         string `xml:"City"`
	State        string `xml:"State"`
	ZIP5         string `xml:"ZIP5"`
	ZIP4         string `xml:"ZIP4"`
	DayOfWeek    string `xml:"DayOfWeek"`
	Date         string `xml:"Date"`
	CarrierRoute string `xml:"CarrierRoute,omitempty"`
}

func (U *USPS) PackagePickupAvailability(request *CarrierPickupAvailabilityRequest) (CarrierPickupAvailabilityResponse, error) {
	request.USERID = U.Username

	result := new(CarrierPickupAvailabilityResponse)
	err := U.Client.Execute(request, result)

	return *result, err
}
