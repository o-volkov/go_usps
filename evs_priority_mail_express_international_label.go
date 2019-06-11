package go_usps

import "encoding/xml"

// eVS Priority Mail Express International Label API
// https://www.usps.com/business/web-tools-apis/evs-international-label-api.htm#_Toc529432214
// TODO: Not tested (API Authorization failure. User XXXXXXXXXXXX is not authorized to use API eVSExpressMailIntl.)
type ImageParametersEVSExpressMailIntlRequest struct {
	ImageParameter string `xml:"ImageParameter,omitempty"`
}

type ShippingContentsEVSExpressMailIntlRequest struct {
	ItemDetail []ItemDetailShippingContentsEVSExpressMailIntlRequest `xml:"ItemDetail"`
}

type ItemDetailShippingContentsEVSExpressMailIntlRequest struct {
	Description     string `xml:"Description"`
	Quantity        string `xml:"Quantity"`
	Value           string `xml:"Value"`
	NetPounds       string `xml:"NetPounds"`
	NetOunces       string `xml:"NetOunces"`
	HSTariffNumber  string `xml:"HSTariffNumber"`
	CountryOfOrigin string `xml:"CountryOfOrigin"`
}

type EVSExpressMailIntlRequest struct {
	XMLName                    xml.Name                                   `xml:"eVSExpressMailIntlRequest"`
	USERID                     string                                     `xml:"USERID,attr"`
	PASSWORD                   string                                     `xml:"PASSWORD,attr,omitempty"`
	APPID                      string                                     `xml:"APPID,attr,omitempty"`
	VERSION                    string                                     `xml:"VERSION,attr,omitempty"`
	Option                     string                                     `xml:"Option,omitempty"`
	Revision                   string                                     `xml:"Revision,omitempty"`
	ImageParameters            []ImageParametersEVSExpressMailIntlRequest `xml:"ImageParameters,omitempty"`
	FromFirstName              string                                     `xml:"FromFirstName,omitempty"`
	FromMiddleInitial          string                                     `xml:"FromMiddleInitial,omitempty"`
	FromLastName               string                                     `xml:"FromLastName,omitempty"`
	FromFirm                   string                                     `xml:"FromFirm,omitempty"`
	FromAddress1               string                                     `xml:"FromAddress1,omitempty"`
	FromAddress2               string                                     `xml:"FromAddress2"`
	FromUrbanization           string                                     `xml:"FromUrbanization,omitempty"`
	FromCity                   string                                     `xml:"FromCity"`
	FromState                  string                                     `xml:"FromState"`
	FromZip5                   string                                     `xml:"FromZip5"`
	FromZip4                   string                                     `xml:"FromZip4,omitempty"`
	FromPhone                  string                                     `xml:"FromPhone"`
	FromCustomsReference       string                                     `xml:"FromCustomsReference,omitempty"`
	ToName                     string                                     `xml:"ToName,omitempty"`
	ToFirstName                string                                     `xml:"ToFirstName,omitempty"`
	ToLastName                 string                                     `xml:"ToLastName,omitempty"`
	ToFirm                     string                                     `xml:"ToFirm,omitempty"`
	ToAddress1                 string                                     `xml:"ToAddress1"`
	ToAddress2                 string                                     `xml:"ToAddress2,omitempty"`
	ToAddress3                 string                                     `xml:"ToAddress3,omitempty"`
	ToCity                     string                                     `xml:"ToCity"`
	ToProvince                 string                                     `xml:"ToProvince,omitempty"`
	ToCountry                  string                                     `xml:"ToCountry"`
	ToPostalCode               string                                     `xml:"ToPostalCode"`
	ToPOBoxFlag                string                                     `xml:"ToPOBoxFlag"`
	ToPhone                    string                                     `xml:"ToPhone,omitempty"`
	ToFax                      string                                     `xml:"ToFax,omitempty"`
	ToEmail                    string                                     `xml:"ToEmail,omitempty"`
	ImportersReferenceNumber   string                                     `xml:"ImportersReferenceNumber,omitempty"`
	NonDeliveryOption          string                                     `xml:"NonDeliveryOption,omitempty"`
	RedirectName               string                                     `xml:"RedirectName,omitempty"`
	RedirectEmail              string                                     `xml:"RedirectEmail,omitempty"`
	RedirectSMS                string                                     `xml:"RedirectSMS,omitempty"`
	RedirectAddress            string                                     `xml:"RedirectAddress,omitempty"`
	RedirectCity               string                                     `xml:"RedirectCity,omitempty"`
	RedirectState              string                                     `xml:"RedirectState,omitempty"`
	RedirectZipCode            string                                     `xml:"RedirectZipCode,omitempty"`
	RedirectZip4               string                                     `xml:"RedirectZip4,omitempty"`
	Container                  string                                     `xml:"Container,omitempty"`
	ShippingContents           ShippingContentsEVSExpressMailIntlRequest  `xml:"ShippingContents"`
	InsuredNumber              string                                     `xml:"InsuredNumber,omitempty"`
	InsuredAmount              string                                     `xml:"InsuredAmount,omitempty"`
	Postage                    string                                     `xml:"Postage,omitempty"`
	GrossPounds                string                                     `xml:"GrossPounds"`
	GrossOunces                string                                     `xml:"GrossOunces"`
	ContentType                string                                     `xml:"ContentType"`
	ContentTypeOther           string                                     `xml:"ContentTypeOther,omitempty"`
	Agreement                  string                                     `xml:"Agreement"`
	Comments                   string                                     `xml:"Comments,omitempty"`
	LicenseNumber              string                                     `xml:"LicenseNumber,omitempty"`
	CertificateNumber          string                                     `xml:"CertificateNumber,omitempty"`
	InvoiceNumber              string                                     `xml:"InvoiceNumber,omitempty"`
	ImageType                  string                                     `xml:"ImageType"`
	ImageLayout                string                                     `xml:"ImageLayout,omitempty"`
	CustomerRefNo              string                                     `xml:"CustomerRefNo,omitempty"`
	CustomerRefNo2             string                                     `xml:"CustomerRefNo2,omitempty"`
	POZipCode                  string                                     `xml:"POZipCode,omitempty"`
	LabelDate                  string                                     `xml:"LabelDate,omitempty"`
	EMCAAccount                string                                     `xml:"EMCAAccount,omitempty"`
	HoldForManifest            string                                     `xml:"HoldForManifest,omitempty"`
	EELPFC                     string                                     `xml:"EELPFC,omitempty"`
	PriceOptions               string                                     `xml:"PriceOptions,omitempty"`
	Size                       string                                     `xml:"Size,omitempty"`
	Length                     string                                     `xml:"Length,omitempty"`
	Width                      string                                     `xml:"Width,omitempty"`
	Height                     string                                     `xml:"Height,omitempty"`
	Girth                      string                                     `xml:"Girth,omitempty"`
	LabelTime                  string                                     `xml:"LabelTime,omitempty"`
	MeterPaymentFlag           string                                     `xml:"MeterPaymentFlag,omitempty"`
	ActionCode                 string                                     `xml:"ActionCode,omitempty"`
	OptOutOfSPE                string                                     `xml:"OptOutOfSPE,omitempty"`
	PermitNumber               string                                     `xml:"PermitNumber,omitempty"`
	ImportersReferenceType     string                                     `xml:"ImportersReferenceType,omitempty"`
	ImportersTelephoneNumber   string                                     `xml:"ImportersTelephoneNumber,omitempty"`
	ImportersFaxNumber         string                                     `xml:"ImportersFaxNumber,omitempty"`
	ImportersEmail             string                                     `xml:"ImportersEmail,omitempty"`
	Machinable                 string                                     `xml:"Machinable,omitempty"`
	DestinationRateIndicator   string                                     `xml:"DestinationRateIndicator,omitempty"`
	MID                        string                                     `xml:"MID,omitempty"`
	LogisticsManagerMID        string                                     `xml:"LogisticsManagerMID,omitempty"`
	CRID                       string                                     `xml:"CRID,omitempty"`
	VendorCode                 string                                     `xml:"VendorCode,omitempty"`
	VendorProductVersionNumber string                                     `xml:"VendorProductVersionNumber,omitempty"`
	EPostageMailerReporting    string                                     `xml:"ePostageMailerReporting,omitempty"`
	SenderFirstName            string                                     `xml:"SenderFirstName,omitempty"`
	SenderLastName             string                                     `xml:"SenderLastName,omitempty"`
	SenderBusinessName         string                                     `xml:"SenderBusinessName,omitempty"`
	SenderAddress1             string                                     `xml:"SenderAddress1,omitempty"`
	SenderCity                 string                                     `xml:"SenderCity,omitempty"`
	SenderState                string                                     `xml:"SenderState,omitempty"`
	SenderZip5                 string                                     `xml:"SenderZip5,omitempty"`
	SenderPhone                string                                     `xml:"SenderPhone,omitempty"`
	SenderEmail                string                                     `xml:"SenderEmail,omitempty"`
}

func (r *EVSExpressMailIntlRequest) toHTTPRequestStr(isProduction bool) (string, error) {
	api := "eVSExpressMailIntlCertify"
	if isProduction {
		api = "eVSExpressMailIntl"
	}

	return createUSPSApiRequestStr(api, r)
}

type EVSExpressMailIntlResponse struct {
	XMLName                  xml.Name `xml:"eVSExpressMailIntlResponse"`
	Postage                  string   `xml:"Postage"`
	TotalValue               string   `xml:"TotalValue"`
	SDRValue                 string   `xml:"SDRValue"`
	BarcodeNumber            string   `xml:"BarcodeNumber"`
	LabelImage               string   `xml:"LabelImage"`
	Page2Image               string   `xml:"Page2Image"`
	Page3Image               string   `xml:"Page3Image"`
	Page4Image               string   `xml:"Page4Image"`
	Page5Image               string   `xml:"Page5Image"`
	Page6Image               string   `xml:"Page6Image"`
	Prohibitions             string   `xml:"Prohibitions"`
	Restrictions             string   `xml:"Restrictions"`
	Observations             string   `xml:"Observations"`
	Regulations              string   `xml:"Regulations"`
	AdditionalRestrictions   string   `xml:"AdditionalRestrictions"`
	InsuranceFee             string   `xml:"InsuranceFee,omitempty"`
	DestinationBarcodeNumber string   `xml:"DestinationBarcodeNumber,omitempty"`
	GuaranteeAvailability    string   `xml:"GuaranteeAvailability,omitempty"`
	RemainingBarcodes        string   `xml:"RemainingBarcodes"`
	Warning                  string   `xml:"Warning,omitempty"`
}

func (U *USPS) EVSPriorityMailExpressInternationalLabel(request *EVSExpressMailIntlRequest) (EVSExpressMailIntlResponse, error) {
	request.USERID = U.Username

	if U.Password != "" {
		request.PASSWORD = U.Password
	}

	result := new(EVSExpressMailIntlResponse)
	err := U.Client.Execute(request, result)

	return *result, err
}
