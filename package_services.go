package go_usps

// USPS Package Services API
// https://www.usps.com/business/web-tools-apis/domestic-mail-service-standards-api.htm#_Toc527965977

type StandardBRequest struct {
	USERID          string `xml:"USERID,attr"`
	OriginZip       string `xml:"OriginZip"`
	DestinationZip  string `xml:"DestinationZip"`
	DestinationType string `xml:"DestinationType,omitempty"`
	ClientType      string `xml:"ClientType,omitempty"`
}

func (r *StandardBRequest) toHTTPRequestStr(bool) (string, error) {
	return createUSPSApiRequestStr("StandardB", r)
}

type StandardBResponse struct {
	OriginZip               string `xml:"OriginZip"`
	DestinationZip          string `xml:"DestinationZip"`
	Days                    string `xml:"Days,omitempty"`
	Message                 string `xml:"Message,omitempty"`
	EffectiveAcceptanceDate string `xml:"EffectiveAcceptanceDate,omitempty"`
	ScheduledDeliveryDate   string `xml:"ScheduledDeliveryDate,omitempty"`
	Error                   struct {
		ErrorDescription string `xml:"ErrorDescription,omitempty"`
		ReturnCode       string `xml:"ReturnCode,omitempty"`
	} `xml:"Error,omitempty"`
}

func (U *USPS) PackageServices(request *StandardBRequest) (StandardBResponse, error) {
	request.USERID = U.Username

	result := new(StandardBResponse)
	err := U.Client.Execute(request, result)

	return *result, err
}
