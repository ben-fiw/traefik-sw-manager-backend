package responses

import (
	"encoding/xml"
)

type ErrorResponse struct {
	XMLName xml.Name `json:"-" xml:"response" yaml:"-"`
	Status  int      `json:"status" xml:"status" yaml:"status"`
	Code    string   `json:"code" xml:"code" yaml:"code"`
	Error   string   `json:"error" xml:"error" yaml:"error"`
}

func NewErrorResponse(status int, code string, error string) *ErrorResponse {
	return &ErrorResponse{
		Status: status,
		Code:   code,
		Error:  error,
	}
}

type MultiErrorResponse struct {
	XMLName xml.Name `json:"-" xml:"response" yaml:"-"`
	Status  int      `json:"status" xml:"status" yaml:"status"`
	Errors  []string `json:"errors" xml:"errors>error" yaml:"errors"`
}

func NewMultiErrorResponse(status int, errors ...string) *MultiErrorResponse {
	return &MultiErrorResponse{
		Status: status,
		Errors: errors,
	}
}

type HealthCheckResponse struct {
	XMLName xml.Name `json:"-" xml:"response" yaml:"-"`
	Status  string   `json:"status" xml:"status" yaml:"status"`
}

func NewHealthCheckResponse(status string) *HealthCheckResponse {
	return &HealthCheckResponse{
		Status: status,
	}
}

type DocumentIndexResponse struct {
	XMLName xml.Name `json:"-" xml:"response" yaml:"-"`
	Meta    struct {
		Total        int    `json:"total" xml:"total" yaml:"total"`
		Page         int    `json:"page" xml:"page" yaml:"page"`
		PageSize     int    `json:"pageSize" xml:"pageSize" yaml:"pageSize"`
		DocumentName string `json:"documentName" xml:"documentName" yaml:"documentName"`
	} `json:"meta" xml:"meta" yaml:"meta"`
	Documents []interface{} `json:"documents" xml:"documents>document" yaml:"documents"`
}

func NewDocumentIndexResponse(total int, page int, pageSize int, documentName string, documents ...interface{}) *DocumentIndexResponse {
	return &DocumentIndexResponse{
		Meta: struct {
			Total        int    `json:"total" xml:"total" yaml:"total"`
			Page         int    `json:"page" xml:"page" yaml:"page"`
			PageSize     int    `json:"pageSize" xml:"pageSize" yaml:"pageSize"`
			DocumentName string `json:"documentName" xml:"documentName" yaml:"documentName"`
		}{
			Total:        total,
			Page:         page,
			PageSize:     pageSize,
			DocumentName: documentName,
		},
		Documents: documents,
	}
}
