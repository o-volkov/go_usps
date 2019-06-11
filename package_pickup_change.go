package go_usps

// Package Pickup Change Web Tools
// https://www.usps.com/business/web-tools-apis/package-pickup-api.htm#_Toc450550264

type PackageCarrierPickupChangeRequest struct {
	ServiceType string `xml:"ServiceType"`
	Count       string `xml:"Count"`
}

type CarrierPickupChangeRequest struct {
	USERID              string                              `xml:"USERID,attr"`
	FirstName           string                              `xml:"FirstName"`
	LastName            string                              `xml:"LastName"`
	FirmName            string                              `xml:"FirmName,omitempty"`
	SuiteOrApt          string                              `xml:"SuiteOrApt,omitempty"`
	Address2            string                              `xml:"Address2"`
	Urbanization        string                              `xml:"Urbanization"`
	City                string                              `xml:"City"`
	State               string                              `xml:"State"`
	ZIP5                string                              `xml:"ZIP5"`
	ZIP4                string                              `xml:"ZIP4"`
	Phone               string                              `xml:"Phone"`
	Extension           string                              `xml:"Extension"`
	Package             []PackageCarrierPickupChangeRequest `xml:"Package"`
	EstimatedWeight     string                              `xml:"EstimatedWeight"`
	PackageLocation     string                              `xml:"PackageLocation"`
	SpecialInstructions string                              `xml:"SpecialInstructions"`
	ConfirmationNumber  string                              `xml:"ConfirmationNumber,omitempty"`
	EmailAddress        string                              `xml:"EmailAddress,omitempty"`
}

func (r *CarrierPickupChangeRequest) toHTTPRequestStr(bool) (string, error) {
	return createUSPSApiRequestStr("CarrierPickupChange", r)
}

type CarrierPickupChangeResponse struct {
	FirstName    string `xml:"FirstName"`
	LastName     string `xml:"LastName"`
	FirmName     string `xml:"FirmName,omitempty"`
	SuiteOrApt   string `xml:"SuiteOrApt,omitempty"`
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
	ConfirmationNumber  string `xml:"ConfirmationNumber,omitempty"`
	DayOfWeek           string `xml:"DayOfWeek"`
	Date                string `xml:"Date"`
	Status              string `xml:"Status,omitempty"`
	EmailAddress        string `xml:"EmailAddress,omitempty"`
}

func (U *USPS) PackagePickupChange(request *CarrierPickupChangeRequest) (CarrierPickupChangeResponse, error) {
	request.USERID = U.Username

	result := new(CarrierPickupChangeResponse)
	err := U.Client.Execute(request, result)

	return *result, err
}
