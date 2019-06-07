package go_usps

// Return Receipt Electronic API
// https://www.usps.com/business/web-tools-apis/track-and-confirm-api.htm#_Toc536704374

type PTSRreRequest struct {
	USERID    string `xml:"USERID,attr"`
	TrackID   string `xml:"TrackID"`
	ClientIp  string `xml:"ClientIp,omitempty"`
	MpSuffix  string `xml:"MpSuffix"`
	MpDate    string `xml:"MpDate"`
	FirstName string `xml:"FirstName"`
	LastName  string `xml:"LastName"`
	Email1    string `xml:"Email1"`
	Email2    string `xml:"Email2,omitempty"`
	Email3    string `xml:"Email3,omitempty"`
	TableCode string `xml:"TableCode"`
	CustRegID string `xml:"CustRegID,omitempty"`
}

func (r *PTSRreRequest) toHTTPRequestStr(bool) (string, error) {
	return createUSPSApiRequestStr("PTSRre", r)
}

type PTSRreResult struct {
	ResultText string `xml:"ResultText"`
	ReturnCode string `xml:"ReturnCode"`
}

func (U *USPS) PTSRre(request *PTSRreRequest) (PTSRreResult, error) {
	request.USERID = U.Username

	result := new(PTSRreResult)
	err := U.Client.Execute(request, result)

	return *result, err
}
