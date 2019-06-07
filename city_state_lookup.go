package go_usps

// City/State Lookup Web Tool
// https://www.usps.com/business/web-tools-apis/address-information-api.htm#_Toc532390356

type CityStateLookupRequest struct {
	USERID   string `xml:"USERID,attr"`
	PASSWORD string `xml:"PASSWORD,attr,omitempty"`
	APPID    string `xml:"APPID,attr,omitempty"`
	Revision string `xml:"Revision,omitempty"`
	ZipCode  struct {
		ID   string `xml:"ID,attr"`
		Zip5 string `xml:"Zip5"`
	} `xml:"ZipCode"`
}

func (r *CityStateLookupRequest) toHTTPRequestStr(bool) (string, error) {
	return createUSPSApiRequestStr("CityStateLookup", r)
}

type CityStateLookupResponse struct {
	ZipCode []struct {
		Zip5               string `xml:"Zip5"`
		City               string `xml:"City"`
		State              string `xml:"State"`
		FinanceNumber      string `xml:"FinanceNumber,omitempty"`
		ClassificationCode string `xml:"ClassificationCode,omitempty"`
		Error              string `xml:"Error,omitempty"`
	} `xml:"ZipCode"`
}

func (U *USPS) CityStateLookup(request *CityStateLookupRequest) (CityStateLookupResponse, error) {
	request.USERID = U.Username

	if U.Password != "" {
		request.PASSWORD = U.Password
	}
	if U.AppId != "" {
		request.APPID = U.AppId
	}

	result := new(CityStateLookupResponse)
	err := U.Client.Execute(request, result)

	return *result, err
}
