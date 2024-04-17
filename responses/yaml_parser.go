package responses

import "gopkg.in/yaml.v2"

type YamlResponseCaster struct {
	Middlewares []ResponseCasterMiddleware
}

func (y YamlResponseCaster) CanCastTo(contentType string) bool {
	return contentType == "text/yaml" || contentType == "application/yaml" || contentType == "application/x-yaml" || contentType == "text/x-yaml"
}

func (y YamlResponseCaster) Cast(response CastableResponse) (CastedResponse, error) {
	for _, middleware := range y.Middlewares {
		response = middleware.BeforeCast(response)
	}

	body, err := yaml.Marshal(response.Body)
	if err != nil {
		return CastedResponse{}, err
	}
	castedResponse := CastedResponse{StatusCode: response.StatusCode, Body: string(body)}

	for _, middleware := range y.Middlewares {
		castedResponse = middleware.AfterCast(castedResponse)
	}

	return castedResponse, nil
}

func init() {
	GlobalResponseCasterRegistry["application/yaml"] = YamlResponseCaster{
		Middlewares: []ResponseCasterMiddleware{
			MultiLineErrorMiddleware{},
		},
	}
}
