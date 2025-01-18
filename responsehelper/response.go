package responsehelper

type ResultSuccess struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ResultError struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Err     interface{} `json:"err"`
}

func ResponseSuccess(message string, data interface{}) ResultSuccess {
	return ResultSuccess{
		Status:  1,
		Message: message,
		Data:    data,
	}
}

func ResponseError(message string, err interface{}) ResultError {
	return ResultError{
		Status:  0,
		Message: message,
		Err:     err,
	}
}
