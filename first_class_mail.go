package go_usps

// USPS First Class Mail API
// https://www.usps.com/business/web-tools-apis/domestic-mail-service-standards-api.htm#_Toc527965984

type FirstClassMailRequest struct {
	USERID          string `xml:"USERID,attr"`
	OriginZip       string `xml:"OriginZip"`
	DestinationZip  string `xml:"DestinationZip"`
	DestinationType string `xml:"DestinationType,omitempty"`
	ClientType      string `xml:"ClientType,omitempty"`
}

func (r *FirstClassMailRequest) toHTTPRequestStr(bool) (string, error) {
	return createUSPSApiRequestStr("FirstClassMail", r)
}

type FirstClassMailResponse struct {
	OriginZip               string `xml:"OriginZip"`
	DestinationZip          string `xml:"DestinationZip"`
	Days                    string `xml:"Days,omitempty"`
	Message                 string `xml:"Message,omitempty"`
	EffectiveAcceptanceDate string `xml:"EffectiveAcceptanceDate,omitempty"`
	ScheduledDeliveryDate   string `xml:"ScheduledDeliveryDate,omitempty"`
}

func (U *USPS) FirstClassMail(request *FirstClassMailRequest) (FirstClassMailResponse, error) {
	request.USERID = U.Username

	result := new(FirstClassMailResponse)
	err := U.Client.Execute(request, result)

	return *result, err
}
