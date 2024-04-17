package responses

import "strings"

type ResponseCasterMiddleware interface {
	BeforeCast(response CastableResponse) CastableResponse
	AfterCast(response CastedResponse) CastedResponse
}

// ###############################
// # Multi Line Error Middleware #
// ###############################

type MultiLineErrorMiddleware struct{}

func (m MultiLineErrorMiddleware) BeforeCast(response CastableResponse) CastableResponse {
	// check if the response is an `ErrorResponse`
	if _, ok := response.Body.(*ErrorResponse); ok {
		// check if the error message is a single line
		if !strings.Contains(response.Body.(*ErrorResponse).Error, "\n") {
			return response
		}

		// split the error message into lines and remove empty lines
		lines := strings.Split(response.Body.(*ErrorResponse).Error, "\n")
		var newLines []string
		for _, line := range lines {
			if line != "" {
				newLines = append(newLines, line)
			}
		}

		// create a new `MultiErrorResponse`
		newResponse := &MultiErrorResponse{
			Status: response.Body.(*ErrorResponse).Status,
			Errors: newLines,
		}

		// return the new response
		return CastableResponse{
			StatusCode: response.StatusCode,
			Body:       newResponse,
		}
	}

	return response
}

func (m MultiLineErrorMiddleware) AfterCast(response CastedResponse) CastedResponse {
	return response
}
