package errors

type RestErr struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Error   string `json:"error"`
}

func BadRequestError(msg string) *RestErr {
	return &RestErr{
		Message: msg,
		Code:    400,
		Error:   "bad_request",
	}
}

func NotFoundError(msg string) *RestErr {
	return &RestErr{
		Message: msg,
		Code:    404,
		Error:   "not_found",
	}
}

func InternalServerError(msg string) *RestErr {
	return &RestErr{
		Message: msg,
		Code:    500,
		Error:   "internal_server_error",
	}
}
