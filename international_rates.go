package go_usps

// International Rates API – IntlRateV2
// https://www.usps.com/business/web-tools-apis/rate-calculator-api.htm#_Toc525907191

type GXGPackageIntlRateV2Request struct {
	POBoxFlag string `xml:"POBoxFlag"`
	GiftFlag  string `xml:"GiftFlag"`
}

type ExtraServicesPackageIntlRateV2Request struct {
	ExtraService []string `xml:"ExtraService,omitempty"`
}

type ContentPackageIntlRateV2Request struct {
	ContentType        string `xml:"ContentType,omitempty"`
	ContentDescription string `xml:"ContentDescription,omitempty"`
}

type PackageIntlRateV2Request struct {
	ID                    string                                 `xml:"ID,attr"`
	Pounds                string                                 `xml:"Pounds"`
	Ounces                string                                 `xml:"Ounces"`
	Machinable            string                                 `xml:"Machinable,omitempty"`
	MailType              string                                 `xml:"MailType"`
	GXG                   *GXGPackageIntlRateV2Request           `xml:"GXG,omitempty"`
	ValueOfContents       string                                 `xml:"ValueOfContents"`
	Country               string                                 `xml:"Country"`
	Container             string                                 `xml:"Container"`
	Size                  string                                 `xml:"Size"`
	Width                 string                                 `xml:"Width"`
	Length                string                                 `xml:"Length"`
	Height                string                                 `xml:"Height"`
	Girth                 string                                 `xml:"Girth"`
	OriginZip             string                                 `xml:"OriginZip,omitempty"`
	CommercialFlag        string                                 `xml:"CommercialFlag,omitempty"`
	CommercialPlusFlag    string                                 `xml:"CommercialPlusFlag,omitempty"`
	ExtraServices         *ExtraServicesPackageIntlRateV2Request `xml:"ExtraServices,omitempty"`
	AcceptanceDateTime    string                                 `xml:"AcceptanceDateTime,omitempty"`
	DestinationPostalCode string                                 `xml:"DestinationPostalCode,omitempty"`
	Content               *ContentPackageIntlRateV2Request       `xml:"Content,omitempty"`
}

type IntlRateV2Request struct {
	USERID   string                     `xml:"USERID,attr"`
	Revision string                     `xml:"Revision,omitempty"`
	Package  []PackageIntlRateV2Request `xml:"Package,omitempty"`
}

func (r *IntlRateV2Request) toHTTPRequestStr(bool) (string, error) {
	return createUSPSApiRequestStr("IntlRateV2", r)
}

type IntlRateV2Response struct {
	Package []struct {
		ID                     string `xml:"ID,attr"`
		Prohibitions           string `xml:"Prohibitions"`
		Restrictions           string `xml:"Restrictions"`
		Observations           string `xml:"Observations"`
		CustomsForms           string `xml:"CustomsForms"`
		ExpressMail            string `xml:"ExpressMail"`
		AreasServed            string `xml:"AreasServed"`
		AdditionalRestrictions string `xml:"AdditionalRestrictions"`
		Content                struct {
			ContentType        string `xml:"ContentType,omitempty"`
			ContentDescription string `xml:"ContentDescription,omitempty"`
		} `xml:"Content,omitempty"`
		Service struct {
			ID         string `xml:"ID,attr"`
			Pounds     string `xml:"Pounds"`
			Ounces     string `xml:"Ounces"`
			Machinable string `xml:"Machinable,omitempty"`
			MailType   string `xml:"MailType"`
			GXG        struct {
				POBoxFlag string `xml:"POBoxFlag"`
				GiftFlag  string `xml:"GiftFlag"`
			} `xml:"GXG,omitempty"`
			Container             string `xml:"Container"`
			Size                  string `xml:"Size"`
			Width                 string `xml:"Width"`
			Length                string `xml:"Length"`
			Height                string `xml:"Height"`
			Girth                 string `xml:"Girth"`
			Country               string `xml:"Country"`
			Postage               string `xml:"Postage"`
			CommercialPostage     string `xml:"CommercialPostage,omitempty"`
			CommercialPlusPostage string `xml:"CommercialPlusPostage,omitempty"`
			ExtraServices         struct {
				ExtraService []struct {
					ServiceID             string `xml:"ServiceID"`
					ServiceName           string `xml:"ServiceName"`
					Available             string `xml:"Available"`
					OnlineAvailable       string `xml:"OnlineAvailable,omitempty"`
					Price                 string `xml:"Price"`
					OnlinePrice           string `xml:"OnlinePrice,omitempty"`
					DeclaredValueRequired string `xml:"DeclaredValueRequired,omitempty"`
				} `xml:"ExtraService,omitempty"`
			} `xml:"ExtraServices,omitempty"`
			ValueOfContents         string `xml:"ValueOfContents,omitempty"`
			InsComment              string `xml:"InsComment,omitempty"`
			ParcelIndemnityCoverage string `xml:"ParcelIndemnityCoverage,omitempty"`
			SvcCommitments          string `xml:"SvcCommitments"`
			SvcDescription          string `xml:"SvcDescription"`
			MaxDimensions           string `xml:"MaxDimensions"`
			MaxWeight               string `xml:"MaxWeight"`
			GuaranteeAvailability   string `xml:"GuaranteeAvailability,omitempty"`
			GXGLocations            struct {
				PostOffice struct {
					Name                string `xml:"Name"`
					Address             string `xml:"Address"`
					City                string `xml:"City"`
					State               string `xml:"State"`
					ZipCode             string `xml:"ZipCode"`
					RetailGXGCutOffTime string `xml:"RetailGXGCutOffTime"`
					SaturdayCutOffTime  string `xml:"SaturdayCutOffTime"`
				} `xml:"PostOffice,omitempty"`
			} `xml:"GXGLocations,omitempty"`
		} `xml:"Service"`
		Error string `xml:"Error,omitempty"`
	} `xml:"Package"`
}

func (U *USPS) InternationalRates(request *IntlRateV2Request) (IntlRateV2Response, error) {
	request.USERID = U.Username

	result := new(IntlRateV2Response)
	err := U.Client.Execute(request, result)

	return *result, err
}
