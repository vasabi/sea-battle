package v1

// HTTPResponse - successful request response model.
type HTTPResponse struct {
	Success bool        `json:"success" example:"true"`
	Result  interface{} `json:"result,omitempty"`
	Errors  []HTTPError `json:"errors,omitempty"`
}

// HTTPError - failed request response model.
type HTTPError struct {
	Code    int    `json:"code" example:"500"`
	Message string `json:"message" example:"Cannot unmarshal json"`
}

// WrapResponse wraps the response body into HTTPResponse model.
func WrapResponse(success bool, result interface{}, errors ...HTTPError) HTTPResponse {
	return HTTPResponse{
		Success: success,
		Result:  result,
		Errors:  errors,
	}
}

// WrapError - wrap the error into HTTPError model.
func WrapError(code int, err error) HTTPError {
	return HTTPError{
		Code:    code,
		Message: err.Error(),
	}
}
