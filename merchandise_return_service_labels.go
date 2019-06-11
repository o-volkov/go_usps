package go_usps

import "encoding/xml"

// USPS Merchandise Return V4.0 Label API
// https://www.usps.com/business/web-tools-apis/merchandise-return-service-label-api.htm#_Toc527616657

type ContentEMRSV40Request struct {
	ContentType string `xml:"ContentType"`
}

type EMRSV40Request struct {
	XMLName                  xml.Name              `xml:"EMRSV4.0Request"`
	USERID                   string                `xml:"USERID,attr"`
	Option                   string                `xml:"Option"`
	Revision                 string                `xml:"Revision,omitempty"`
	CustomerName             string                `xml:"CustomerName"`
	CustomerAddress1         string                `xml:"CustomerAddress1"`
	CustomerAddress2         string                `xml:"CustomerAddress2"`
	CustomerCity             string                `xml:"CustomerCity"`
	CustomerState            string                `xml:"CustomerState"`
	CustomerZip5             string                `xml:"CustomerZip5"`
	CustomerZip4             string                `xml:"CustomerZip4"`
	RetailerName             string                `xml:"RetailerName"`
	RetailerAddress          string                `xml:"RetailerAddress"`
	PermitNumber             string                `xml:"PermitNumber"`
	PermitIssuingPOCity      string                `xml:"PermitIssuingPOCity"`
	PermitIssuingPOState     string                `xml:"PermitIssuingPOState"`
	PermitIssuingPOZip5      string                `xml:"PermitIssuingPOZip5"`
	PDUFirmName              string                `xml:"PDUFirmName,omitempty"`
	PDUPOBox                 string                `xml:"PDUPOBox"`
	PDUCity                  string                `xml:"PDUCity"`
	PDUState                 string                `xml:"PDUState"`
	PDUZip5                  string                `xml:"PDUZip5"`
	PDUZip4                  string                `xml:"PDUZip4"`
	ServiceType              string                `xml:"ServiceType"`
	DeliveryConfirmation     string                `xml:"DeliveryConfirmation"`
	InsuranceValue           string                `xml:"InsuranceValue"`
	MailingAckPackageID      string                `xml:"MailingAckPackageID,omitempty"`
	WeightInPounds           string                `xml:"WeightInPounds"`
	WeightInOunces           string                `xml:"WeightInOunces"`
	RMA                      string                `xml:"RMA"`
	RMAPICFlag               string                `xml:"RMAPICFlag"`
	ImageType                string                `xml:"ImageType"`
	SenderName               string                `xml:"SenderName,omitempty"`
	SenderEMail              string                `xml:"SenderEMail,omitempty"`
	RecipientName            string                `xml:"RecipientName,omitempty"`
	RecipientEMail           string                `xml:"RecipientEMail,omitempty"`
	RMABarcode               string                `xml:"RMABarcode,omitempty"`
	AllowNonCleansedDestAddr string                `xml:"AllowNonCleansedDestAddr,omitempty"`
	AllowNonCleansedCustAddr string                `xml:"AllowNonCleansedCustAddr,omitempty"`
	NineDigitRoutingZip      string                `xml:"NineDigitRoutingZip,omitempty"`
	Content                  ContentEMRSV40Request `xml:"Content,omitempty"`
	GroundOnly               string                `xml:"GroundOnly,omitempty"`
	Oversized                string                `xml:"Oversized,omitempty"`
}

func (r *EMRSV40Request) toHTTPRequestStr(isProduction bool) (string, error) {
	api := "MerchReturnCertifyV4"
	if isProduction {
		api = "MerchandiseReturnV4"
	}

	return createUSPSApiRequestStr(api, r)
}

type EMRSV40Response struct {
	XMLName                xml.Name `xml:"EMRSV4.0Response"`
	Zone                   string   `xml:"Zone"`
	MerchandiseReturnLabel string   `xml:"MerchandiseReturnLabel"`
	InsuranceCost          string   `xml:"InsuranceCost"`
	PDUFirmName            string   `xml:"PDUFirmName"`
	PDUPOBox               string   `xml:"PDUPOBox"`
	PDUPOBoxAbbreviation   string   `xml:"PDUPOBoxAbbreviation,omitempty"`
	PDUCity                string   `xml:"PDUCity"`
	PDUCityAbbreviation    string   `xml:"PDUCityAbbreviation,omitempty"`
	PDUState               string   `xml:"PDUState"`
	PDUZip5                string   `xml:"PDUZip5"`
	PDUZip4                string   `xml:"PDUZip4"`
	Postnet                string   `xml:"Postnet"`
	CustomerAddress1       string   `xml:"CustomerAddress1"`
	CustomerAddress2       string   `xml:"CustomerAddress2"`
	CustomerCity           string   `xml:"CustomerCity"`
	CustomerState          string   `xml:"CustomerState"`
	CustomerZip5           string   `xml:"CustomerZip5"`
	CustomerZip4           string   `xml:"CustomerZip4"`
	CustomerPostnet        string   `xml:"CustomerPostnet"`
	RDC                    string   `xml:"RDC,omitempty"`
	LogMessage             string   `xml:"LogMessage,omitempty"`
}

func (U *USPS) MerchandiseReturnServiceLabels(request *EMRSV40Request) (EMRSV40Response, error) {
	request.USERID = U.Username

	result := new(EMRSV40Response)
	err := U.Client.Execute(request, result)

	return *result, err
}
