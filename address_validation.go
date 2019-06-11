package go_usps

import "errors"

// Address Web Tool
// https://www.usps.com/business/web-tools-apis/address-information-api.htm#_Toc532390339

type AddressValidateRequest struct {
	USERID   string `xml:"USERID,attr"`
	PASSWORD string `xml:"PASSWORD,attr,omitempty"`
	APPID    string `xml:"APPID,attr,omitempty"`
	Revision string `xml:"Revision,omitempty"`
	Address  struct {
		ID           string `xml:"ID,attr"`
		FirmName     string `xml:"FirmName"`
		Address1     string `xml:"Address1"`
		Address2     string `xml:"Address2"`
		City         string `xml:"City"`
		State        string `xml:"State"`
		Urbanization string `xml:"Urbanization"`
		Zip5         string `xml:"Zip5"`
		Zip4         string `xml:"Zip4"`
	} `xml:"Address"`
}

func (r *AddressValidateRequest) toHTTPRequestStr(bool) (string, error) {
	return createUSPSApiRequestStr("Verify", r)
}

type Address struct {
	FirmName     string `xml:"FirmName,omitempty"`
	Address1     string `xml:"Address1,omitempty"`
	Address2     string `xml:"Address2,omitempty"`
	City         string `xml:"City,omitempty"`
	State        string `xml:"State,omitempty"`
	Urbanization string `xml:"Urbanization,omitempty"`
	Zip5         string `xml:"Zip5,omitempty"`
	Zip4         string `xml:"Zip4,omitempty"`
}

type AddressValidateResponse struct {
	Address struct {
		ID            string `xml:"ID,omitempty"`
		FirmName      string `xml:"FirmName,omitempty"`
		Address1      string `xml:"Address1,omitempty"`
		Address2      string `xml:"Address2,omitempty"`
		City          string `xml:"City,omitempty"`
		State         string `xml:"State,omitempty"`
		Urbanization  string `xml:"Urbanization,omitempty"`
		Zip5          string `xml:"Zip5,omitempty"`
		Zip4          string `xml:"Zip4,omitempty"`
		DeliveryPoint string `xml:"DeliveryPoint,omitempty"`
		CarrierRoute  string `xml:"CarrierRoute,omitempty"`
		Error         *Error `xml:"Error,omitempty"`
	} `xml:"Address"`
}

func (U *USPS) AddressVerification(request *AddressValidateRequest) (AddressValidateResponse, error) {
	request.USERID = U.Username

	if U.Password != "" {
		request.PASSWORD = U.Password
	}
	if U.AppId != "" {
		request.APPID = U.AppId
	}

	result := new(AddressValidateResponse)
	err := U.Client.Execute(request, result)

	if result.Address.Error != nil {
		return *result, errors.New(result.Address.Error.Description)
	}

	return *result, err
}
