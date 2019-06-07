package go_usps

const (
	defaultHttpsDev  string = "https://secure.shippingapis.com/ShippingAPITest.dll?API="
	defaultHttpsProd string = "https://secure.shippingapis.com/ShippingAPI.dll?API="
)

type USPSClient interface {
	Execute(request USPSHTTPRequest, result interface{}) error
	callRequest(requestURL string) ([]byte, error)
}

type USPS struct {
	Username   string
	Password   string
	AppId      string
	Production bool `default:"false"`
	Client     USPSClient
}

func InitUSPS(username, password, appId string, isProduction bool) *USPS {
	usps := new(USPS)
	usps.Username = username
	usps.Password = password
	usps.AppId = appId
	if isProduction {
		usps.Production = true
	}

	usps.Client = &USPSHttpClient{
		HttpClient: getHttpClient(),
		Production: isProduction,
	}

	return usps
}
