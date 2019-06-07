package go_usps

// ZIP Code Lookup Web Tool
// https://www.usps.com/business/web-tools-apis/address-information-api.htm#_Toc532390349

type ZipCodeLookupRequest struct {
	USERID   string `xml:"USERID,attr"`
	PASSWORD string `xml:"PASSWORD,attr,omitempty"`
	Address  struct {
		ID           string `xml:"ID,attr"`
		FirmName     string `xml:"FirmName,omitempty"`
		Address1     string `xml:"Address1,omitempty"`
		Address2     string `xml:"Address2,omitempty"`
		City         string `xml:"City,omitempty"`
		State        string `xml:"State,omitempty"`
		Urbanization string `xml:"Urbanization,omitempty"`
	} `xml:"Address"`
}

func (r *ZipCodeLookupRequest) toHTTPRequestStr(bool) (string, error) {
	return createUSPSApiRequestStr("ZipCodeLookup", r)
}

type ZipCodeLookupResponse struct {
	Address []struct {
		FirmName     string `xml:"FirmName,omitempty"`
		Address1     string `xml:"Address1,omitempty"`
		Address2     string `xml:"Address2,omitempty"`
		City         string `xml:"City,omitempty"`
		State        string `xml:"State,omitempty"`
		Urbanization string `xml:"Urbanization,omitempty"`
		Zip5         string `xml:"Zip5,omitempty"`
		Zip4         string `xml:"Zip4,omitempty"`
		Error        string `xml:"Error,omitempty"`
	} `xml:"Address"`
}

func (U *USPS) ZipCodeLookup(request *ZipCodeLookupRequest) (ZipCodeLookupResponse, error) {
	request.USERID = U.Username

	if U.Password != "" {
		request.PASSWORD = U.Password
	}

	result := new(ZipCodeLookupResponse)
	err := U.Client.Execute(request, result)

	return *result, err
}
