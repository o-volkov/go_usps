package go_usps

type USPSHTTPRequest interface {
	toHTTPRequestStr(isProduction bool) (string, error)
}
