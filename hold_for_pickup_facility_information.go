package go_usps

// USPS Hold for Pickup Facility Information API
// https://www.usps.com/business/web-tools-apis/hold-for-pickup-facilities-lookup-api.htm#_Toc504654015

type HFPFacilityInfoRequest struct {
	USERID      string `xml:"USERID,attr"`
	PASSWORD    string `xml:"PASSWORD,attr,omitempty"`
	PickupCity  string `xml:"PickupCity"`
	PickupState string `xml:"PickupState"`
	PickupZIP   string `xml:"PickupZIP"`
	PickupZIP4  string `xml:"PickupZIP4"`
	Service     string `xml:"Service,omitempty"`
}

func (r *HFPFacilityInfoRequest) toHTTPRequestStr(bool) (string, error) {
	return createUSPSApiRequestStr("HFPFacilityInfo", r)
}

type HFPFacilityInfoResponse struct {
	PickupCity  string `xml:"PickupCity"`
	PickupState string `xml:"PickupState"`
	PickupZIP   string `xml:"PickupZIP"`
	PickupZIP4  string `xml:"PickupZIP4"`
	Service     string `xml:"Service,omitempty"`
	Facility    []struct {
		FacilityID        string `xml:"FacilityID"`
		FacilityName      string `xml:"FacilityName"`
		FacilityAddress   string `xml:"FacilityAddress"`
		FacilityCity      string `xml:"FacilityCity"`
		FacilityState     string `xml:"FacilityState"`
		FacilityZIP       string `xml:"FacilityZIP"`
		FacilityZIP4      string `xml:"FacilityZIP4"`
		Has10amCommitment string `xml:"Has10amCommitment"`
	} `xml:"Facility"`
	Has10amCommitment string `xml:"LogMessage,omitempty"`
}

func (U *USPS) HoldForPickupFacilityInformation(request *HFPFacilityInfoRequest) (HFPFacilityInfoResponse, error) {
	request.USERID = U.Username

	if U.Password != "" {
		request.PASSWORD = U.Password
	}

	result := new(HFPFacilityInfoResponse)
	err := U.Client.Execute(request, result)

	return *result, err
}
