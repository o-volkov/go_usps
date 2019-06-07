package go_usps

// Track Request API
// https://www.usps.com/business/web-tools-apis/track-and-confirm-api.htm#_Toc536704352

type TrackRequest struct {
	USERID  string `xml:"USERID,attr"`
	TrackID []struct {
		ID string `xml:"ID,attr"`
	} `xml:"TrackID"`
}

func (r *TrackRequest) toHTTPRequestStr(bool) (string, error) {
	return createUSPSApiRequestStr("TrackV2", r)
}

type TrackResults struct {
	RequestSeqNumber string `xml:"RequestSeqNumber"`
	TrackInfo        []struct {
		ID                       string   `xml:"ID,attr"`
		DeliveryNotificationDate string   `xml:"DeliveryNotificationDate"`
		ExpectedDeliveryDate     string   `xml:"ExpectedDeliveryDate,omitempty"`
		ExpectedDeliveryTime     string   `xml:"ExpectedDeliveryTime,omitempty"`
		GuaranteedDeliveryDate   string   `xml:"GuaranteedDeliveryDate,omitempty"`
		TrackSummary             string   `xml:"TrackSummary,omitempty"`
		TrackDetail              []string `xml:"TrackDetail,omitempty"`
	}
}

func (U *USPS) Track(request *TrackRequest) (TrackResults, error) {
	request.USERID = U.Username

	result := new(TrackResults)
	err := U.Client.Execute(request, result)

	return *result, err
}
