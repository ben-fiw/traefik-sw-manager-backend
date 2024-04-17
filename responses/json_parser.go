package responses

import (
	"encoding/json"
)

// ##############################
// #    JSON Response Caster    #
// ##############################

type JsonResponseCaster struct {
	MiddleWares []ResponseCasterMiddleware
}

func (j JsonResponseCaster) CanCastTo(contentType string) bool {
	return contentType == "application/json"
}

func (j JsonResponseCaster) Cast(response CastableResponse) (CastedResponse, error) {
	for _, middleware := range j.MiddleWares {
		response = middleware.BeforeCast(response)
	}

	body, err := json.Marshal(response.Body)
	if err != nil {
		return CastedResponse{}, err
	}
	castedResponse := CastedResponse{StatusCode: response.StatusCode, Body: string(body)}

	for _, middleware := range j.MiddleWares {
		castedResponse = middleware.AfterCast(castedResponse)
	}

	return castedResponse, nil
}

func init() {
	GlobalResponseCasterRegistry["application/json"] = JsonResponseCaster{
		MiddleWares: []ResponseCasterMiddleware{
			MultiLineErrorMiddleware{},
		},
	}
}
