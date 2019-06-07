package go_usps

import "encoding/xml"

// USPS Merchandise Return Bulk Label API
// https://www.usps.com/business/web-tools-apis/merchandise-return-services-bulk-api.htm#_Toc527615003

type ImageParametersEMRSV40BulkRequest struct {
	ImageParameter string `xml:"ImageParameter,omitempty"`
}

type EMRSV40BulkRequest struct {
	XMLName                  xml.Name                            `xml:"EMRSV4.0BulkRequest"`
	USERID                   string                              `xml:"USERID,attr"`
	Option                   string                              `xml:"Option"`
	Revision                 string                              `xml:"Revision,omitempty"`
	LabelCount               string                              `xml:"LabelCount"`
	ImageParameters          []ImageParametersEMRSV40BulkRequest `xml:"ImageParameters,omitempty"`
	RetailerName             string                              `xml:"RetailerName"`
	RetailerAddress          string                              `xml:"RetailerAddress"`
	PermitNumber             string                              `xml:"PermitNumber"`
	PermitIssuingPOCity      string                              `xml:"PermitIssuingPOCity"`
	PermitIssuingPOState     string                              `xml:"PermitIssuingPOState"`
	PermitIssuingPOZip5      string                              `xml:"PermitIssuingPOZip5"`
	PDUFirmName              string                              `xml:"PDUFirmName,omitempty"`
	PDUPOBox                 string                              `xml:"PDUPOBox"`
	PDUCity                  string                              `xml:"PDUCity"`
	PDUState                 string                              `xml:"PDUState"`
	PDUZip5                  string                              `xml:"PDUZip5"`
	PDUZip4                  string                              `xml:"PDUZip4"`
	ServiceType              string                              `xml:"ServiceType"`
	DeliveryConfirmation     string                              `xml:"DeliveryConfirmation"`
	InsuranceValue           string                              `xml:"InsuranceValue"`
	WeightInPounds           string                              `xml:"WeightInPounds"`
	WeightInOunces           string                              `xml:"WeightInOunces"`
	ImageType                string                              `xml:"ImageType"`
	SenderName               string                              `xml:"SenderName,omitempty"`
	SenderEMail              string                              `xml:"SenderEMail,omitempty"`
	RecipientName            string                              `xml:"RecipientName,omitempty"`
	RecipientEMail           string                              `xml:"RecipientEMail,omitempty"`
	AllowNonCleansedDestAddr string                              `xml:"AllowNonCleansedDestAddr,omitempty"`
	NineDigitRoutingZip      string                              `xml:"NineDigitRoutingZip,omitempty"`
}

func (r *EMRSV40BulkRequest) toHTTPRequestStr(isProduction bool) (string, error) {
	api := "MerchReturnV4BulkCertify"
	if isProduction {
		api = "MerchReturnV4Bulk"
	}

	return createUSPSApiRequestStr(api, r)
}

// TODO: Implement real response
type EMRSV40BulkResponse struct {
}

func (U *USPS) MerchandiseReturnServiceBulkLabels(request *EMRSV40BulkRequest) (EMRSV40BulkResponse, error) {
	request.USERID = U.Username

	result := new(EMRSV40BulkResponse)
	err := U.Client.Execute(request, result)

	return *result, err
}
