package go_usps

import (
	"bytes"
	"encoding/xml"
	"net/url"
	"strings"
)

func urlEncode(urlToEncode string) string {
	return url.QueryEscape(urlToEncode)
}

func createUSPSApiRequestStr(api string, r interface{}) (string, error) {
	xmlOut, err := xml.Marshal(r)
	if err != nil {
		return "", err
	}

	var requestURL bytes.Buffer
	requestURL.WriteString(api + "&XML=")
	requestURL.WriteString(urlEncode(string(xmlOut)))

	return requestURL.String(), nil
}

func parseUSPSXml(xmlBytes []byte, s interface{}) error {
	bodyHl := strings.Replace(string(xmlBytes), xml.Header, "", 1)
	e := new(Error)
	err := xml.Unmarshal([]byte(bodyHl), &e)
	if err != nil {
		return err
	}
	if e != nil && e.Number != "" {
		return e
	}

	return xml.Unmarshal([]byte(bodyHl), &s)
}
