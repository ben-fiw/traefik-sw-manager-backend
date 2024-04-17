package requests

import "gopkg.in/yaml.v2"

type YamlRequestParser struct{}

func (y YamlRequestParser) CanParse(contentType string) bool {
	return contentType == "application/yaml" || contentType == "application/x-yaml" || contentType == "text/yaml" || contentType == "text/x-yaml"
}

func (y YamlRequestParser) Parse(body []byte, v interface{}) error {
	return yaml.Unmarshal(body, v)
}

func init() {
	GlobalRequestParserRegistry["application/yaml"] = YamlRequestParser{}
}
