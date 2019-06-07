package go_usps

// PriorityMail
// https://www.usps.com/business/web-tools-apis/domestic-mail-service-standards-api.htm#_Toc527965972

type PriorityMailRequest struct {
	USERID          string `xml:"USERID,attr"`
	OriginZip       string `xml:"OriginZip"`
	DestinationZip  string `xml:"DestinationZip"`
	DestinationType string `xml:"DestinationType,omitempty"`
	PMGuarantee     string `xml:"PMGuarantee,omitempty"`
	ClientType      string `xml:"ClientType,omitempty"`
}

func (r *PriorityMailRequest) toHTTPRequestStr(bool) (string, error) {
	return createUSPSApiRequestStr("PriorityMail", r)
}

type PriorityMailResponse struct {
	OriginZip               string `xml:"OriginZip"`
	DestinationZip          string `xml:"DestinationZip"`
	Days                    string `xml:"Days,omitempty"`
	IsGuaranteed            string `xml:"IsGuaranteed,omitempty"`
	Message                 string `xml:"Message,omitempty"`
	EffectiveAcceptanceDate string `xml:"EffectiveAcceptanceDate,omitempty"`
	ScheduledDeliveryDate   string `xml:"ScheduledDeliveryDate,omitempty"`
	Error                   struct {
		ErrorDescription string `xml:"ErrorDescription,omitempty"`
		ReturnCode       string `xml:"ReturnCode,omitempty"`
	} `xml:"Error,omitempty"`
}

func (U *USPS) PriorityMail(request *PriorityMailRequest) (PriorityMailResponse, error) {
	request.USERID = U.Username

	result := new(PriorityMailResponse)
	err := U.Client.Execute(request, result)

	return *result, err
}
