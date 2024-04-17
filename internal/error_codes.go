package internal

type ErrorCode struct {
	CausingParty     string
	ErrorCode        string
	ErrorDescription string
	StatusCode       int
}

type ErrorGenerator func(args ...interface{}) *ErrorCode

func (e *ErrorCode) BuildCode() string {
	return "E-" + e.CausingParty + "-" + e.ErrorCode
}

func (e *ErrorCode) Describe(d string) *ErrorCode {
	e.ErrorDescription = d
	return e
}

type ErrorCodesType struct {
	InternalError      ErrorGenerator
	InvalidServiceAuth ErrorGenerator
	InvalidRequest     ErrorGenerator
	NotFound           ErrorGenerator
	ServiceNotFound    ErrorGenerator
	DuplicateValue     ErrorGenerator
}

func (e *ErrorCodesType) GetErrorCode(statusCode int) ErrorGenerator {
	switch statusCode {
	case 500:
		return e.InternalError
	case 401:
		return e.InvalidServiceAuth
	case 400:
		return e.InvalidRequest
	case 404:
		return e.ServiceNotFound
	default:
		return e.InternalError
	}
}

var ErrorCodes = ErrorCodesType{
	InternalError: func(args ...interface{}) *ErrorCode {
		return &ErrorCode{
			CausingParty:     "S",
			ErrorCode:        "INTERNAL_ERROR",
			ErrorDescription: "An internal error occurred",
			StatusCode:       500,
		}
	},
	InvalidServiceAuth: func(args ...interface{}) *ErrorCode {
		return &ErrorCode{
			CausingParty:     "S",
			ErrorCode:        "INVALID_SERVICE_AUTH",
			ErrorDescription: "Invalid service authentication",
			StatusCode:       401,
		}
	},
	InvalidRequest: func(args ...interface{}) *ErrorCode {
		return &ErrorCode{
			CausingParty:     "C",
			ErrorCode:        "INVALID_REQUEST",
			ErrorDescription: "Invalid request",
			StatusCode:       400,
		}
	},
	NotFound: func(args ...interface{}) *ErrorCode {
		return &ErrorCode{
			CausingParty:     "C",
			ErrorCode:        "NOT_FOUND",
			ErrorDescription: "Not found",
			StatusCode:       404,
		}
	},
	DuplicateValue: func(args ...interface{}) *ErrorCode {
		errorDescription := "Duplicate value"
		if len(args) > 0 {
			errorDescription += " for " + args[0].(string)
		}

		return &ErrorCode{
			CausingParty:     "S",
			ErrorCode:        "DUPLICATE_VALUE",
			ErrorDescription: errorDescription,
			StatusCode:       400,
		}
	},
}
