package requests

import "encoding/xml"

type XmlRequestParser struct{}

func (x XmlRequestParser) CanParse(contentType string) bool {
	return contentType == "application/xml"
}

func (x XmlRequestParser) Parse(body []byte, v interface{}) error {
	return xml.Unmarshal(body, v)
}

func init() {
	GlobalRequestParserRegistry["application/xml"] = XmlRequestParser{}
}
