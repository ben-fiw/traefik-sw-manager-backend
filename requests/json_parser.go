package requests

import "encoding/json"

type JsonRequestParser struct{}

func (j JsonRequestParser) CanParse(contentType string) bool {
	return contentType == "application/json"
}

func (j JsonRequestParser) Parse(body []byte, v interface{}) error {
	return json.Unmarshal(body, v)
}

func init() {
	GlobalRequestParserRegistry["application/json"] = JsonRequestParser{}
}
