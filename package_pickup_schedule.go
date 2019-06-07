package go_usps

// Package Pickup Schedule Web Tools
// https://www.usps.com/business/web-tools-apis/package-pickup-api.htm#_Toc450550252

type PackageCarrierPickupScheduleRequest struct {
	ServiceType string `xml:"ServiceType"`
	Count       string `xml:"Count"`
}

type CarrierPickupScheduleRequest struct {
	USERID              string                              `xml:"USERID,attr"`
	FirstName           string                              `xml:"FirstName"`
	LastName            string                              `xml:"LastName"`
	FirmName            string                              `xml:"FirmName,omitempty"`
	SuiteOrApt          string                              `xml:"SuiteOrApt"`
	Address2            string                              `xml:"Address2"`
	Urbanization        string                              `xml:"Urbanization"`
	City                string                              `xml:"City"`
	State               string                              `xml:"State"`
	ZIP5                string                              `xml:"ZIP5"`
	ZIP4                string                              `xml:"ZIP4"`
	Phone               string                              `xml:"Phone"`
	Extension           string                              `xml:"Extension,omitempty"`
	Package             PackageCarrierPickupScheduleRequest `xml:"Package"`
	EstimatedWeight     string                              `xml:"EstimatedWeight"`
	PackageLocation     string                              `xml:"PackageLocation"`
	SpecialInstructions string                              `xml:"SpecialInstructions,omitempty"`
	EmailAddress        string                              `xml:"EmailAddress,omitempty"`
}

func (r *CarrierPickupScheduleRequest) toHTTPRequestStr(bool) (string, error) {
	return createUSPSApiRequestStr("CarrierPickupSchedule", r)
}

type CarrierPickupScheduleResponse struct {
	FirstName    string `xml:"FirstName"`
	LastName     string `xml:"LastName"`
	FirmName     string `xml:"FirmName,omitempty"`
	SuiteOrApt   string `xml:"SuiteOrApt"`
	Address2     string `xml:"Address2"`
	Urbanization string `xml:"Urbanization"`
	City         string `xml:"City"`
	State        string `xml:"State"`
	ZIP5         string `xml:"ZIP5"`
	ZIP4         string `xml:"ZIP4"`
	Phone        string `xml:"Phone"`
	Extension    string `xml:"Extension,omitempty"`
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
	CarrierRoute        string `xml:"CarrierRoute,omitempty"`
	EmailAddress        string `xml:"EmailAddress,omitempty"`
}

func (U *USPS) PackagePickupSchedule(request *CarrierPickupScheduleRequest) (CarrierPickupScheduleResponse, error) {
	request.USERID = U.Username

	result := new(CarrierPickupScheduleResponse)
	err := U.Client.Execute(request, result)

	return *result, err
}
