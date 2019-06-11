package go_usps

// Proof of Delivery API
// https://www.usps.com/business/web-tools-apis/track-and-confirm-api.htm#_Toc536704369
// TODO: Not tested (API Authorization failure. User XXXXXXXXXXXX is not authorized to use API PTSPod.)
type PTSPodRequest struct {
	USERID        string `xml:"USERID,attr"`
	TrackID       string `xml:"TrackID"`
	ClientIp      string `xml:"ClientIp,omitempty"`
	MpSuffix      string `xml:"MpSuffix"`
	MpDate        string `xml:"MpDate"`
	RequestType   string `xml:"RequestType"`
	FirstName     string `xml:"FirstName,omitempty"`
	LastName      string `xml:"LastName,omitempty"`
	Email1        string `xml:"Email1"`
	Email2        string `xml:"Email2,omitempty"`
	Email3        string `xml:"Email3,omitempty"`
	FaxNumber     string `xml:"FaxNumber,omitempty"`
	AddressLine1  string `xml:"AddressLine1,omitempty"`
	AddressLine2  string `xml:"AddressLine2,omitempty"`
	City          string `xml:"City,omitempty"`
	State         string `xml:"State,omitempty"`
	Zip           string `xml:"Zip,omitempty"`
	VerifyAddress string `xml:"VerifyAddress,omitempty"`
	TableCode     string `xml:"TableCode"`
	CustRegID     string `xml:"CustRegID,omitempty"`
}

func (r *PTSPodRequest) toHTTPRequestStr(bool) (string, error) {
	return createUSPSApiRequestStr("PTSPod", r)
}

type PTSPodResult struct {
	ResultText string `xml:"ResultText"`
	ReturnCode string `xml:"ReturnCode"`
}

func (U *USPS) PTSPod(request *PTSPodRequest) (PTSPodResult, error) {
	request.USERID = U.Username

	result := new(PTSPodResult)
	err := U.Client.Execute(request, result)

	return *result, err
}
