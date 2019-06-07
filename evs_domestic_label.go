package go_usps

import "encoding/xml"

// eVS Label API
// https://www.usps.com/business/web-tools-apis/evs-label-api.htm#_Toc487532684

type ImageParametersEVSRequest struct {
	ImageParameter string `xml:"ImageParameter,omitempty"`
	XCoordinate    string `xml:"XCoordinate,omitempty"`
	YCoordinate    string `xml:"YCoordinate,omitempty"`
	LabelSequence  struct {
		PackageNumber string `xml:"PackageNumber,omitempty"`
		TotalPackages string `xml:"TotalPackages,omitempty"`
	} `xml:"LabelSequence,omitempty"`
}

type ExpressMailOptionsEVSRequest struct {
	DeliveryOption    string `xml:"DeliveryOption,omitempty"`
	WaiverOfSignature string `xml:"WaiverOfSignature,omitempty"`
	ESOFAllowed       string `xml:"eSOFAllowed,omitempty"`
}

type ExtraServicesEVSRequest struct {
	ExtraService []string `xml:"ExtraService,omitempty"`
}

type ContentEVSRequest struct {
	ContentType        string `xml:"ContentType,omitempty"`
	ContentDescription string `xml:"ContentDescription,omitempty"`
}

type EVSRequest struct {
	XMLName                      xml.Name                     `xml:"eVSRequest"`
	USERID                       string                       `xml:"USERID,attr"`
	PASSWORD                     string                       `xml:"PASSWORD,attr,omitempty"`
	Option                       string                       `xml:"Option,omitempty"`
	Revision                     string                       `xml:"Revision,omitempty"`
	ImageParameters              []ImageParametersEVSRequest  `xml:"ImageParameters,omitempty"`
	FromName                     string                       `xml:"FromName"`
	FromFirm                     string                       `xml:"FromFirm"`
	FromAddress1                 string                       `xml:"FromAddress1"`
	FromAddress2                 string                       `xml:"FromAddress2"`
	FromCity                     string                       `xml:"FromCity"`
	FromState                    string                       `xml:"FromState"`
	FromZip5                     string                       `xml:"FromZip5"`
	FromZip4                     string                       `xml:"FromZip4"`
	FromPhone                    string                       `xml:"FromPhone"`
	POZipCode                    string                       `xml:"POZipCode,omitempty"`
	AllowNonCleansedOriginAddr   string                       `xml:"AllowNonCleansedOriginAddr,omitempty"`
	ToName                       string                       `xml:"ToName"`
	ToFirm                       string                       `xml:"ToFirm"`
	ToAddress1                   string                       `xml:"ToAddress1"`
	ToAddress2                   string                       `xml:"ToAddress2"`
	ToCity                       string                       `xml:"ToCity"`
	ToState                      string                       `xml:"ToState"`
	ToZip5                       string                       `xml:"ToZip5"`
	ToZip4                       string                       `xml:"ToZip4"`
	ToPhone                      string                       `xml:"ToPhone"`
	POBox                        string                       `xml:"POBox,omitempty"`
	ToContactPreference          string                       `xml:"ToContactPreference,omitempty"`
	ToContactMessaging           string                       `xml:"ToContactMessaging,omitempty"`
	ToContactEmail               string                       `xml:"ToContactEmail,omitempty"`
	AllowNonCleansedDestAddr     string                       `xml:"AllowNonCleansedDestAddr,omitempty"`
	WeightInOunces               string                       `xml:"WeightInOunces"`
	ServiceType                  string                       `xml:"ServiceType"`
	Container                    string                       `xml:"Container,omitempty"`
	Width                        string                       `xml:"Width,omitempty"`
	Length                       string                       `xml:"Length,omitempty"`
	Height                       string                       `xml:"Height,omitempty"`
	Machinable                   string                       `xml:"Machinable,omitempty"`
	ProcessingCategory           string                       `xml:"ProcessingCategory,omitempty"`
	PriceOptions                 string                       `xml:"PriceOptions,omitempty"`
	InsuredAmount                string                       `xml:"InsuredAmount,omitempty"`
	AddressServiceRequested      string                       `xml:"AddressServiceRequested,omitempty"`
	ExpressMailOptions           ExpressMailOptionsEVSRequest `xml:"ExpressMailOptions,omitempty"`
	ShipDate                     string                       `xml:"ShipDate,omitempty"`
	CustomerRefNo                string                       `xml:"CustomerRefNo,omitempty"`
	CustomerRefNo2               string                       `xml:"CustomerRefNo2,omitempty"`
	ExtraServices                ExtraServicesEVSRequest      `xml:"ExtraServices,omitempty"`
	HoldForPickup                string                       `xml:"HoldForPickup,omitempty"`
	OpenDistribute               string                       `xml:"OpenDistribute,omitempty"`
	PermitNumber                 string                       `xml:"PermitNumber,omitempty"`
	PermitZIPCode                string                       `xml:"PermitZIPCode,omitempty"`
	PermitHolderName             string                       `xml:"PermitHolderName,omitempty"`
	CRID                         string                       `xml:"CRID,omitempty"`
	MID                          string                       `xml:"MID,omitempty"`
	LogisticsManagerMID          string                       `xml:"LogisticsManagerMID,omitempty"`
	VendorCode                   string                       `xml:"VendorCode,omitempty"`
	VendorProductVersionNumber   string                       `xml:"VendorProductVersionNumber,omitempty"`
	SenderName                   string                       `xml:"SenderName,omitempty"`
	SenderEMail                  string                       `xml:"SenderEMail,omitempty"`
	RecipientName                string                       `xml:"RecipientName,omitempty"`
	RecipientEMail               string                       `xml:"RecipientEMail,omitempty"`
	ReceiptOption                string                       `xml:"ReceiptOption,omitempty"`
	ImageType                    string                       `xml:"ImageType,omitempty"`
	HoldForManifest              string                       `xml:"HoldForManifest,omitempty"`
	NineDigitRoutingZip          string                       `xml:"NineDigitRoutingZip,omitempty"`
	ShipInfo                     string                       `xml:"ShipInfo,omitempty"`
	CarrierRelease               string                       `xml:"CarrierRelease,omitempty"`
	DropOffTime                  string                       `xml:"DropOffTime,omitempty"`
	ReturnCommitments            string                       `xml:"ReturnCommitments,omitempty"`
	PrintCustomerRefNo           string                       `xml:"PrintCustomerRefNo,omitempty"`
	PrintCustomerRefNo2          string                       `xml:"PrintCustomerRefNo2,omitempty"`
	Content                      ContentEVSRequest            `xml:"Content,omitempty"`
	ActionCode                   string                       `xml:"ActionCode,omitempty"`
	OptOutOfSPE                  string                       `xml:"OptOutOfSPE,omitempty"`
	SortationLevel               string                       `xml:"SortationLevel,omitempty"`
	DestinationEntryFacilityType string                       `xml:"DestinationEntryFacilityType,omitempty"`
}

func (r *EVSRequest) toHTTPRequestStr(isProduction bool) (string, error) {
	api := "eVSCertify"
	if isProduction {
		api = "eVS"
	}

	return createUSPSApiRequestStr(api, r)
}

type EVSResponse struct {
	XMLName                xml.Name `xml:"eVSResponse"`
	BarcodeNumber          string   `xml:"BarcodeNumber"`
	LabelImage             string   `xml:"LabelImage,omitempty"`
	ReceiptImage           string   `xml:"ReceiptImage,omitempty"`
	ToName                 string   `xml:"ToName"`
	ToFirm                 string   `xml:"ToFirm"`
	ToAddress1             string   `xml:"ToAddress1"`
	ToAddress2             string   `xml:"ToAddress2"`
	ToAddress2Abbreviation string   `xml:"ToAddress2Abbreviation,omitempty"`
	ToCity                 string   `xml:"ToCity"`
	ToCityAbbreviation     string   `xml:"ToCityAbbreviation,omitempty"`
	ToState                string   `xml:"ToState"`
	ToZip5                 string   `xml:"ToZip5"`
	ToZip4                 string   `xml:"ToZip4"`
	Postnet                string   `xml:"Postnet"`
	RDC                    string   `xml:"RDC"`
	Postage                string   `xml:"Postage"`
	ExtraServices          struct {
		ExtraService []struct {
			ServiceID   string `xml:"ServiceID"`
			ServiceName string `xml:"ServiceName"`
			Price       string `xml:"Price"`
		} `xml:"ExtraService,omitempty"`
	} `xml:"ExtraServices,omitempty"`
	HoldForPickup     string `xml:"HoldForPickup,omitempty"`
	Zone              string `xml:"Zone"`
	DimensionalWeight string `xml:"DimensionalWeight,omitempty"`
	CarrierRoute      string `xml:"CarrierRoute"`
	PermitHolderName  string `xml:"PermitHolderName"`
	InductionType     string `xml:"InductionType"`
	LogMessage        string `xml:"LogMessage"`
	Commitment        struct {
		CommitmentName        string `xml:"CommitmentName"`
		ScheduledDeliveryDate string `xml:"ScheduledDeliveryDate"`
	} `xml:"Commitment,omitempty"`
}

func (U *USPS) EVSDomesticLabel(request *EVSRequest) (EVSResponse, error) {
	request.USERID = U.Username

	if U.Password != "" {
		request.PASSWORD = U.Password
	}

	result := new(EVSResponse)
	err := U.Client.Execute(request, result)

	return *result, err
}
