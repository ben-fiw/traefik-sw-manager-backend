package responses

import "encoding/xml"

type XmlResponseCaster struct {
	Middlewares []ResponseCasterMiddleware
}

func (x XmlResponseCaster) CanCastTo(contentType string) bool {
	return contentType == "application/xml"
}

func (x XmlResponseCaster) Cast(response CastableResponse) (CastedResponse, error) {
	for _, middleware := range x.Middlewares {
		response = middleware.BeforeCast(response)
	}

	body, err := xml.Marshal(response.Body)
	if err != nil {
		return CastedResponse{}, err
	}
	castedResponse := CastedResponse{StatusCode: response.StatusCode, Body: string(body)}

	for _, middleware := range x.Middlewares {
		castedResponse = middleware.AfterCast(castedResponse)
	}

	return castedResponse, nil
}

func init() {
	GlobalResponseCasterRegistry["application/xml"] = XmlResponseCaster{
		Middlewares: []ResponseCasterMiddleware{
			MultiLineErrorMiddleware{},
		},
	}
}
