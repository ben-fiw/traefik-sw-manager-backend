package requests

import (
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type ParsedRequest struct {
	ContentType string
	Body        interface{}
}

type RequestParser interface {
	CanParse(contentType string) bool
	Parse(body []byte, v interface{}) error
}

var GlobalRequestParserRegistry = make(map[string]RequestParser)

func ParseRequest(r *http.Request, v interface{}) error {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	contentType := r.Header.Get("Content-Type")
	for _, parser := range GlobalRequestParserRegistry {
		if parser.CanParse(contentType) {
			err := parser.Parse(body, v)
			if err != nil {
				return err
			}

			validate := validator.New()
			return validate.Struct(v)
		}
	}

	return errors.New("unsupported content type")
}
