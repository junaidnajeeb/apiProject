package modal

type apiError struct {
	ErrorMessage string `json:"ErrorMessage"`
	ErrorCode    int    `json:"ErrorCode"`
}

const defaultErrorCode = 0

// APIError export error struct
//type APIError apiError

// APIErrorMessageInstance returns error object with just message
func APIErrorMessageInstance(message string) apiError {
	return apiError{ErrorMessage: message, ErrorCode: defaultErrorCode}
}

// APIErrorMessageCodeInstance returns error object with just message
func APIErrorMessageCodeInstance(message string, code int) apiError {
	return apiError{ErrorMessage: message, ErrorCode: code}
}
