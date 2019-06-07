package go_usps

import (
	"reflect"
	"strings"
	"testing"
)

func TestUSPS_EVSPriorityMailExpressInternationalLabel(t *testing.T) {
	username := "USPSUsername"
	password := "USPSPassword"

	successRespXmlStr := `
<?xml version="1.0"?>
<?xml version="1.0"?>
<eVSExpressMailIntlResponse>
  <Postage>66.21</Postage>
  <TotalValue>55.00</TotalValue>
  <SDRValue>65.80</SDRValue>
  <BarcodeNumber>EC502016320US</BarcodeNumber>
  <LabelImage>.....removed.....</LabelImage>
  <Page2Image>.....removed.....</Page2Image>
  <Page3Image>.....removed.....</Page3Image>
  <Page4Image>.....removed.....</Page4Image>
  <Page5Image>.....removed.....</Page5Image>
  <Page6Image></Page6Image>
  <Prohibitions>
    Coins; bank notes; currency notes (paper money); securities of any kind payable to bearer; traveler's checks; platinum, gold, and silver (except for jewelry items meeting the requirement in "Restrictions" below); precious stones (except when contained in jewelry items meeting the requirement in "Restrictions" below); and other valuable articles are prohibited.
    Fruit cartons (used or new).
    Goods bearing the name "Anzac."
    Goods produced wholly or partly in prisons or by convict labor.
    Most food, plant, and animal products, including the use of products such as straw and other plant material as packing materials.
    Perishable infectious biological substances.
    Radioactive materials.
    Registered philatelic articles with fictitious addresses.
    Seditious literature.
    Silencers for firearms.
    Used bedding.
  </Prohibitions>
  <Restrictions>
    Jewelry is permitted only when sent as an insured parcel using Priority Mail International service. In addition, Australian Customs regulations prohibit importation of jewelry that is made with ivory or from endangered species, such as snake, elephant, or crocodile, that does not have an accompanying Import/Export Permit in relation to the Convention on International Trade in Endangered Species of Wild Fauna and Flora (CITES).
    Meat and other animal products; powdered or concentrated milk; and other dairy products requires permission to import from the Australian quarantine authorities.
    Permission of the Australian Director-General of Health is required to import medicines.
  </Restrictions>
  <Observations>Duty may be levied on catalogs, price lists, circulars, and all advertising introduced into Australia through the mail, regardless of the class of mail used.</Observations>
  <Regulations>
    Country Code:
    AU
    
    Reciprocal Service Name:

    Express Post
    Required Customs Form/Endorsement
    1. Business and commercial papers.
    No form required. Endorse item clearly next to mailing label as BUSINESS PAPERS.
    2. Merchandise samples without commercial value microfilm, microfiche, and computer data.
    PS Form 2976-A, Customs Declaration and Dispatch Note CP 72, inside a PS Form 2976-E, Customs Declaration Envelope CP 91.
    3. Merchandise and all articles subject to customs duty.
    PS Form 2976-A, Customs Declaration and Dispatch Note CP 72, inside a PS Form 2976-E, Customs Declaration Envelope CP 91.
    Note:
    1. Coins; banknotes; currency notes, including paper money; securities of any kind payable to bearer; traveler's checks; platinum, gold, and silver; precious stones; jewelry; watches; and other valuable articles are prohibited in Priority Mail Express International shipments to Australia.
    2. Priority Mail Express International With Guarantee service - which offers a date-certain, postage-refund guarantee - is available to Australia.

    Areas Served:
    
    All except Lord Howe Island and the Australian Antarctic territories.
  </Regulations>
  <AdditionalRestrictions>No Additional Restrictions Data found.</AdditionalRestrictions>
  <InsuranceFee>0</InsuranceFee>
  <GuaranteeAvailability>03/11/2015</GuaranteeAvailability>
  <RemainingBarcodes>29</RemainingBarcodes>
</eVSExpressMailIntlResponse>
`
	request := EVSExpressMailIntlRequest{
		USERID:   username,
		PASSWORD: password,
	}

	request.FromAddress2 = "7 North Wilke-Barre Blvd"

	rStr, _ := request.toHTTPRequestStr(false)
	requestResponseMap := map[string][]byte{
		rStr: []byte(strings.TrimSuffix(successRespXmlStr, "\n")),
	}

	type fields struct {
		Username   string
		Password   string
		AppId      string
		Production bool
		Client     USPSClient
	}
	type args struct {
		request *EVSExpressMailIntlRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{name: "Success flow",
			fields:  fields{Username: username, Password: password, Client: &TestClient{RequestResponseMap: requestResponseMap}},
			args:    args{request: &request},
			want:    "EC502016320US",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			U := &USPS{
				Username:   tt.fields.Username,
				Password:   tt.fields.Password,
				AppId:      tt.fields.AppId,
				Production: tt.fields.Production,
				Client:     tt.fields.Client,
			}
			got, err := U.EVSPriorityMailExpressInternationalLabel(tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("USPS.EVSPriorityMailExpressInternationalLabel() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.BarcodeNumber, tt.want) {
				t.Errorf("USPS.EVSPriorityMailExpressInternationalLabel() = %v, want %v", got.BarcodeNumber, tt.want)
			}
		})
	}
}
