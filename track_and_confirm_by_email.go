package go_usps

// Track and Confirm by Email API
// https://www.usps.com/business/web-tools-apis/track-and-confirm-api.htm#_Toc536704365
// TODO: Not tested (API Authorization failure. User XXXXXXXXXXXX is not authorized to use API PTSEmail.)
type PTSEmailRequest struct {
	USERID      string `xml:"USERID,attr"`
	TrackID     string `xml:"TrackID"`
	ClientIp    string `xml:"ClientIp,omitempty"`
	MpSuffix    string `xml:"MpSuffix"`
	MpDate      string `xml:"MpDate"`
	RequestType string `xml:"RequestType"`
	FirstName   string `xml:"FirstName,omitempty"`
	LastName    string `xml:"LastName,omitempty"`
	Email1      string `xml:"Email1"`
	Email2      string `xml:"Email2,omitempty"`
	Email3      string `xml:"Email3,omitempty"`
}

func (r *PTSEmailRequest) toHTTPRequestStr(bool) (string, error) {
	return createUSPSApiRequestStr("PTSEmail", r)
}

type PTSEmailResult struct {
	ResultText string `xml:"ResultText"`
	ReturnCode string `xml:"ReturnCode"`
}

func (U *USPS) PTSEmail(request *PTSEmailRequest) (PTSEmailResult, error) {
	request.USERID = U.Username

	result := new(PTSEmailResult)
	err := U.Client.Execute(request, result)

	return *result, err
}
