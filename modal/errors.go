package modal

type errors struct {
	ErrorMessage string `json:"ErrorMessage"`
	ErrorCode    int    `json:"ErrorCode"`
}

const defaultErrorCode = 0

// ErrorMessageInstance returns error object with just message
func ErrorMessageInstance(message string) errors {
	return errors{ErrorMessage: message, ErrorCode: defaultErrorCode}
}

// ErrorMessageCodeInstance returns error object with just message
func ErrorMessageCodeInstance(message string, code int) errors {
	return errors{ErrorMessage: message, ErrorCode: code}
}
