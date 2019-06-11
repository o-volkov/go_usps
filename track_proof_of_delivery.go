package go_usps

// Track Proof of Delivery
// https://www.usps.com/business/web-tools-apis/track-and-confirm-api.htm#_Toc536704379
// TODO: Not tested (API Authorization failure. User XXXXXXXXXXXX is not authorized to use API PTSPod.)
type PTSTPodRequest struct {
	USERID        string `xml:"USERID,attr"`
	TrackID       string `xml:"TrackID"`
	MpSuffix      string `xml:"MpSuffix"`
	MpDate        string `xml:"MpDate"`
	RequestType   string `xml:"RequestType"`
	FirstName     string `xml:"FirstName"`
	LastName      string `xml:"LastName"`
	Email1        string `xml:"Email1"`
	Email2        string `xml:"Email2,omitempty"`
	Email3        string `xml:"Email3,omitempty"`
	SignInID      string `xml:"SignInID,omitempty"`
	CustRegID     string `xml:"CustRegID"`
	VerifyAddress string `xml:"VerifyAddress,omitempty"`
	TableCode     string `xml:"TableCode"`
	ClientIp      string `xml:"ClientIp,omitempty"`
	SourceId      string `xml:"SourceId,omitempty"`
}

func (r *PTSTPodRequest) toHTTPRequestStr(bool) (string, error) {
	return createUSPSApiRequestStr("PTSTPod", r)
}

type PTSTPodResult struct {
	ResultText string `xml:"ResultText"`
	ReturnCode string `xml:"ReturnCode"`
}

func (U *USPS) PTSTPod(request *PTSPodRequest) (PTSTPodResult, error) {
	request.USERID = U.Username

	result := new(PTSTPodResult)
	err := U.Client.Execute(request, result)

	return *result, err
}
