package go_usps

// Track/Confirm Fields API
// https://www.usps.com/business/web-tools-apis/track-and-confirm-api.htm#_Toc536704357

type TrackIDTrackFieldRequest struct {
	ID                 string `xml:"ID,attr"`
	DestinationZipCode string `xml:"DestinationZipCode,omitempty"`
	MailingDate        string `xml:"MailingDate,omitempty"`
}

type TrackFieldRequest struct {
	USERID      string                     `xml:"USERID,attr"`
	Revision    string                     `xml:"Revision,omitempty"`
	ClientIp    string                     `xml:"ClientIp,omitempty"`
	SourceId    string                     `xml:"SourceId,omitempty"`
	SourceIdZIP string                     `xml:"SourceIdZIP,omitempty"`
	TrackID     []TrackIDTrackFieldRequest `xml:"TrackID"`
}

func (r *TrackFieldRequest) toHTTPRequestStr(bool) (string, error) {
	return createUSPSApiRequestStr("TrackV2", r)
}

type TrackResponse struct {
	TrackInfo []struct {
		ID                     string `xml:"ID,attr"`
		GuaranteedDeliveryDate string `xml:"GuaranteedDeliveryDate,omitempty"`
		ESOFEligible           string `xml:"eSOFEligible,omitempty"`
		TrackSummary           struct {
			EventTime       string `xml:"EventTime"`
			EventDate       string `xml:"EventDate"`
			Event           string `xml:"Event"`
			EventCity       string `xml:"EventCity"`
			EventState      string `xml:"EventState"`
			EventZIPCode    string `xml:"EventZIPCode"`
			EventCountry    string `xml:"EventCountry,omitempty"`
			FirmName        string `xml:"FirmName,omitempty"`
			Name            string `xml:"Name,omitempty"`
			AuthorizedAgent string `xml:"AuthorizedAgent,omitempty"`
			EventCode       string `xml:"EventCode,omitempty"`
			ActionCode      string `xml:"ActionCode,omitempty"`
			ReasonCode      string `xml:"ReasonCode,omitempty"`
		} `xml:"TrackSummary"`
		TrackDetail struct {
			EventTime       string `xml:"EventTime"`
			EventDate       string `xml:"EventDate"`
			Event           string `xml:"Event"`
			EventCity       string `xml:"EventCity"`
			EventState      string `xml:"EventState"`
			EventZIPCode    string `xml:"EventZIPCode"`
			EventCountry    string `xml:"EventCountry,omitempty"`
			FirmName        string `xml:"FirmName,omitempty"`
			Name            string `xml:"Name,omitempty"`
			AuthorizedAgent string `xml:"AuthorizedAgent,omitempty"`
			EventCode       string `xml:"EventCode,omitempty"`
			ActionCode      string `xml:"ActionCode,omitempty"`
			ReasonCode      string `xml:"ReasonCode,omitempty"`
		} `xml:"TrackDetail"`
	} `xml:"TrackInfo"`
}

func (U *USPS) TrackField(request *TrackFieldRequest) (TrackResponse, error) {
	request.USERID = U.Username

	result := new(TrackResponse)
	err := U.Client.Execute(request, result)

	return *result, err
}
