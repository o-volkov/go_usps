package go_usps

// USPS SCAN API
// https://www.usps.com/business/web-tools-apis/scan-api.htm#_Toc514661497

const (
	Form3152 Form = "3152"
	Form5630 Form = "5630"
)

const (
	ImageTypeTif  ImageType = "TIF"
	ImageTypePdf  ImageType = "PDF"
	ImageTypeNone ImageType = "NONE"
)

type Form string
type ImageType string

type SCANRequest struct {
	USERID   string `xml:"USERID,attr"`
	Revision string `xml:"Revision,omitempty"`
	Option   struct {
		Form Form `default:"",xml:"Form,omitempty"`
	} `xml:"Option,omitempty"`
	FromName     string `xml:"FromName"`
	FromFirm     string `xml:"FromFirm,omitempty"`
	FromAddress1 string `xml:"FromAddress1"`
	FromAddress2 string `xml:"FromAddress2"`
	FromCity     string `xml:"FromCity"`
	FromState    string `xml:"FromState"`
	FromZip5     string `xml:"FromZip5"`
	FromZip4     string `xml:"FromZip4"`
	Shipment     struct {
		PackageDetail struct {
			PkgBarcode     string `xml:"PkgBarcode"`
			SpecialService struct {
				SpcServCode string `xml:"SpcServCode"`
				SpcServFee  string `xml:"SpcServFee"`
			} `xml:"SpecialService"`
			EMail string `default:"",xml:"EMail"`
		} `xml:"PackageDetail"`
	} `xml:"Shipment"`
	CloseManifest string    `xml:"CloseManifest,omitempty"`
	MailDate      string    `xml:"MailDate"`
	MailTime      string    `xml:"MailTime"`
	EntryFacility string    `xml:"EntryFacility"`
	ImageType     ImageType `xml:"ImageType"`
	CustomerRefNo string    `xml:"CustomerRefNo,omitempty"`
	CarrierPickup bool      `xml:"CarrierPickup,omitempty"`
}

func (r *SCANRequest) toHTTPRequestStr(isProduction bool) (string, error) {
	api := "SCANCertify"
	if isProduction {
		api = "SCAN"
	}

	return createUSPSApiRequestStr(api, r)
}

type SCANResponse struct {
	SCANFormNumber string `xml:"SCANFormNumber"`
	SCANFormImage  string `xml:"SCANFormImage"`
}

func (U *USPS) Scan(request *SCANRequest) (SCANResponse, error) {
	request.USERID = U.Username

	result := new(SCANResponse)
	err := U.Client.Execute(request, result)

	return *result, err
}
