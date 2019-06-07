package go_usps

// USPS Express Mail Service Commitments API
// https://www.usps.com/business/web-tools-apis/domestic-mail-service-standards-api.htm#_Toc527965991

type ExpressMailCommitmentRequest struct {
	USERID         string `xml:"USERID,attr"`
	OriginZip      string `xml:"OriginZip"`
	DestinationZIP string `xml:"DestinationZIP"`
	Date           string `xml:"Date"`
	DropOffTime    string `xml:"DropOffTime,omitempty"`
	ReturnDates    string `xml:"ReturnDates,omitempty"`
	PMGuarantee    string `xml:"PMGuarantee,omitempty"`
}

func (r *ExpressMailCommitmentRequest) toHTTPRequestStr(bool) (string, error) {
	return createUSPSApiRequestStr("ExpressMailCommitment", r)
}

type ExpressMailCommitmentResponse struct {
	OriginZip               string `xml:"OriginZip"`
	OriginCity              string `xml:"OriginCity"`
	OriginState             string `xml:"OriginState"`
	DestinationZip          string `xml:"DestinationZip"`
	DestinationCity         string `xml:"DestinationCity"`
	DestinationState        string `xml:"DestinationState"`
	Date                    string `xml:"Date"`
	Time                    string `xml:"Time"`
	ExpeditedTransMessage   string `xml:"ExpeditedTransMessage,omitempty"`
	MsgCode                 string `xml:"MsgCode,omitempty"`
	Msg                     string `xml:"Msg,omitempty"`
	EffectiveAcceptanceDate string `xml:"EffectiveAcceptanceDate,omitempty"`
	Commitment              []struct {
		Name     string `xml:"Name,omitempty"`
		Time     string `xml:"Time,omitempty"`
		Sequence string `xml:"Sequence,omitempty"`
		Location []struct {
			ScheduledDeliveryDate string `xml:"ScheduledDeliveryDate,omitempty"`
			CutOff                string `xml:"CutOff,omitempty"`
			Facility              string `xml:"Facility,omitempty"`
			Street                string `xml:"Street,omitempty"`
			City                  string `xml:"City,omitempty"`
			State                 string `xml:"State,omitempty"`
			Zip                   string `xml:"Zip,omitempty"`
			IsGuaranteed          string `xml:"IsGuaranteed,omitempty"`
		} `xml:"Location,omitempty"`
	} `xml:"Commitment,omitempty"`
	Message string `xml:"Message,omitempty"`
}

func (U *USPS) ExpressMailServiceCommitments(request *ExpressMailCommitmentRequest) (ExpressMailCommitmentResponse, error) {
	request.USERID = U.Username

	result := new(ExpressMailCommitmentResponse)
	err := U.Client.Execute(request, result)

	return *result, err
}
