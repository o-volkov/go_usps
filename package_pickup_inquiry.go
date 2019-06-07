package go_usps

// Package Pickup Inquiry Web Tools
// https://www.usps.com/business/web-tools-apis/package-pickup-api.htm#_Toc450550270

type CarrierPickupInquiryRequest struct {
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

func (r *CarrierPickupInquiryRequest) toHTTPRequestStr(bool) (string, error) {
	return createUSPSApiRequestStr("CarrierPickupInquiry", r)
}

type CarrierPickupInquiryResponse struct {
	FirstName    string `xml:"FirstName"`
	LastName     string `xml:"LastName"`
	FirmName     string `xml:"FirmName"`
	SuiteOrApt   string `xml:"SuiteOrApt"`
	Address2     string `xml:"Address2"`
	Urbanization string `xml:"Urbanization"`
	City         string `xml:"City"`
	State        string `xml:"State"`
	ZIP5         string `xml:"ZIP5"`
	ZIP4         string `xml:"ZIP4"`
	Phone        string `xml:"Phone"`
	Extension    string `xml:"Extension"`
	Package      struct {
		ServiceType string `xml:"ServiceType"`
		Count       string `xml:"Count"`
	} `xml:"Package"`
	EstimatedWeight     string `xml:"EstimatedWeight"`
	PackageLocation     string `xml:"PackageLocation"`
	SpecialInstructions string `xml:"SpecialInstructions"`
	ConfirmationNumber  string `xml:"ConfirmationNumber"`
	DayOfWeek           string `xml:"DayOfWeek"`
	Date                string `xml:"Date"`
	EmailAddress        string `xml:"EmailAddress,omitempty"`
}

func (U *USPS) PackagePickupInquiry(request *CarrierPickupInquiryRequest) (CarrierPickupInquiryResponse, error) {
	request.USERID = U.Username

	result := new(CarrierPickupInquiryResponse)
	err := U.Client.Execute(request, result)

	return *result, err
}
