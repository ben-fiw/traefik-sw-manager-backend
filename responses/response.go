package responses

import (
	"demo-shop-manager/internal"
	"errors"
	"net/http"
	"os"
	"strings"
)

// ##############################
// #      Global Responders     #
// ##############################

func SendErrorCode(r *http.Request, w http.ResponseWriter, err *internal.ErrorCode) {
	response := &ErrorResponse{
		Status: err.StatusCode,
		Code:   err.BuildCode(),
		Error:  err.ErrorDescription,
	}

	castableResponse := &CastableResponse{
		StatusCode: err.StatusCode,
		Body:       response,
	}

	castableResponse.CastAndSend(r, w)
}

func SendError(r *http.Request, w http.ResponseWriter, statusCode int, errorMessage string) {
	errorCode := internal.ErrorCodes.GetErrorCode(statusCode)()

	response := &ErrorResponse{
		Status: statusCode,
		Code:   errorCode.BuildCode(),
		Error:  errorMessage,
	}

	castableResponse := &CastableResponse{
		StatusCode: statusCode,
		Body:       response,
	}

	castableResponse.CastAndSend(r, w)
}

func SendResponse(r *http.Request, w http.ResponseWriter, statusCode int, body interface{}) {
	response := &CastableResponse{
		StatusCode: statusCode,
		Body:       body,
	}

	response.CastAndSend(r, w)
}

// ##############################
// #         Responses          #
// ##############################

type CastableResponse struct {
	StatusCode int
	Body       interface{}
}

type CastedResponse struct {
	StatusCode int
	Body       string
}

func (c *CastableResponse) Cast(contentType string) (CastedResponse, error) {
	for _, caster := range GlobalResponseCasterRegistry {
		if caster.CanCastTo(contentType) {
			return caster.Cast(*c)
		}
	}

	// return an error if no caster found
	return CastedResponse{}, errors.New("unable to cast response to " + contentType)
}

func (c *CastableResponse) CastAndSend(r *http.Request, w http.ResponseWriter) {
	// Get accepted content types
	contentTypes := strings.Split(r.Header.Get("Accept"), ",")
	if len(contentTypes) == 0 || contentTypes[0] == "" || contentTypes[0] == "*/*" {
		contentTypes = []string{os.Getenv("DEFAULT_CONTENT_TYPE")}
	}

	// Try to cast the response to the accepted content types
	for _, contentType := range contentTypes {
		castedResponse, err := c.Cast(contentType)
		if err == nil {
			w.Header().Set("Content-Type", contentType)
			w.WriteHeader(castedResponse.StatusCode)
			w.Write([]byte(castedResponse.Body))
			return
		}
	}

	// If no casters were found, return a 406
	w.WriteHeader(http.StatusNotAcceptable)
	w.Write([]byte("406 Not Acceptable"))
}

// ##############################
// #       Response Casters     #
// ##############################

type ResponseCaster interface {
	CanCastTo(contentType string) bool
	Cast(response CastableResponse) (CastedResponse, error)
}

var GlobalResponseCasterRegistry = make(map[string]ResponseCaster)
