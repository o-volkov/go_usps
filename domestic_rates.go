package go_usps

// Domestic Rates
// https://www.usps.com/business/web-tools-apis/rate-calculator-api.htm#_Toc525907183

type PackageSpecialServiceRateV4Request struct {
	SpecialService string `xml:"SpecialService,omitempty"`
}

type PackageContentRateV4Request struct {
	ContentType        string `xml:"ContentType,omitempty"`
	ContentDescription string `xml:"ContentDescription,omitempty"`
}

type PackageShipDateRateV4Request struct {
	Option string `xml:"Option,attr,omitempty"`
}

type PackageRateV4Request struct {
	ID                           string                               `xml:"ID,attr"`
	SortationLevel               string                               `xml:"SortationLevel,omitempty"`
	DestinationEntryFacilityType string                               `xml:"DestinationEntryFacilityType,omitempty"`
	Nonprofit                    string                               `xml:"Nonprofit,omitempty"`
	Service                      string                               `xml:"Service"`
	FirstClassMailType           string                               `xml:"FirstClassMailType,omitempty"`
	ZipOrigination               string                               `xml:"ZipOrigination"`
	ZipDestination               string                               `xml:"ZipDestination"`
	Pounds                       string                               `xml:"Pounds"`
	Ounces                       string                               `xml:"Ounces"`
	Container                    string                               `xml:"Container"`
	Size                         string                               `xml:"Size"`
	Width                        string                               `xml:"Width,omitempty"`
	Length                       string                               `xml:"Length,omitempty"`
	Height                       string                               `xml:"Height,omitempty"`
	Girth                        string                               `xml:"Girth,omitempty"`
	Value                        string                               `xml:"Value,omitempty"`
	AmountToCollect              string                               `xml:"AmountToCollect,omitempty"`
	SpecialServices              []PackageSpecialServiceRateV4Request `xml:"SpecialServices,omitempty"`
	Content                      *PackageContentRateV4Request         `xml:"Content,omitempty"`
	GroundOnly                   string                               `xml:"GroundOnly,omitempty"`
	SortBy                       string                               `xml:"SortBy,omitempty"`
	Machinable                   string                               `xml:"Machinable,omitempty"`
	ReturnLocations              string                               `xml:"ReturnLocations,omitempty"`
	ReturnServiceInfo            string                               `xml:"ReturnServiceInfo,omitempty"`
	DropOffTime                  string                               `xml:"DropOffTime,omitempty"`
	ShipDate                     *PackageShipDateRateV4Request        `xml:"ShipDate,omitempty"`
}

type RateV4Request struct {
	USERID   string                 `xml:"USERID,attr"`
	Revision string                 `xml:"Revision,omitempty"`
	Package  []PackageRateV4Request `xml:"Package,omitempty"`
}

func (r *RateV4Request) toHTTPRequestStr(bool) (string, error) {
	return createUSPSApiRequestStr("RateV4", r)
}

type RateV4Response struct {
	Package []struct {
		ID                 string `xml:"ID,attr"`
		ZipOrigination     string `xml:"ZipOrigination"`
		ZipDestination     string `xml:"ZipDestination"`
		Pounds             string `xml:"Pounds"`
		Ounces             string `xml:"Ounces"`
		FirstClassMailType string `xml:"FirstClassMailType,omitempty"`
		Container          string `xml:"Container,omitempty"`
		Size               string `xml:"Size"`
		Width              string `xml:"Width,omitempty"`
		Length             string `xml:"Length,omitempty"`
		Height             string `xml:"Height,omitempty"`
		Girth              string `xml:"Girth,omitempty"`
		Machinable         string `xml:"Machinable,omitempty"`
		Zone               string `xml:"Zone,omitempty"`
		Postage            struct {
			CLASSID            string `xml:"CLASSID,attr"`
			MailService        string `xml:"MailService"`
			Rate               string `xml:"Rate"`
			CommercialRate     string `xml:"CommercialRate,omitempty"`
			CommercialPlusRate string `xml:"CommercialPlusRate,omitempty"`
			MaxDimensions      string `xml:"MaxDimensions,omitempty"`
			ServiceInformation string `xml:"ServiceInformation,omitempty"`
			SpecialServices    struct {
				SpecialService struct {
					ServiceID             string `xml:"ServiceID"`
					ServiceName           string `xml:"ServiceName"`
					Available             string `xml:"Available"`
					AvailableOnline       string `xml:"AvailableOnline"`
					AvailableCPP          string `xml:"AvailableCPP"`
					Price                 string `xml:"Price"`
					PriceOnline           string `xml:"PriceOnline"`
					PriceCPP              string `xml:"PriceCPP"`
					DeclaredValueRequired string `xml:"DeclaredValueRequired,omitempty"`
					DueSenderRequired     string `xml:"DueSenderRequired,omitempty"`
				} `xml:"SpecialService,omitempty"`
			} `xml:"SpecialServices,omitempty"`
			Zone string `xml:"Zone,omitempty"`
		} `xml:"Postage"`
		Restriction struct {
			Restrictions string `xml:"Restrictions"`
		} `xml:"Restriction"`
		Error string `xml:"Error,omitempty"`
	} `xml:"Package"`
}

func (U *USPS) RateDomestic(request *RateV4Request) (RateV4Response, error) {
	request.USERID = U.Username

	result := new(RateV4Response)
	err := U.Client.Execute(request, result)

	return *result, err
}
