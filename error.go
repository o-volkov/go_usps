package go_usps

type Error struct {
	Number      string `xml:"Number"`
	Source      string `xml:"Source,omitempty"`
	Description string `xml:"Description,omitempty"`
	HelpFile    string `xml:"HelpFile,omitempty"`
	HelpContext string `xml:"HelpContext,omitempty"`
}

func (e *Error) Error() string {
	return e.Description
}
