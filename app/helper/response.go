package helper

const StatusSuccess = "success"
const StatusFailure = "fail"

//Response is used for static shape json return
type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors"`
	Data    interface{} `json:"data"`
}

//BuildResponse method is to inject data value to dynamic success response
func BuildResponse(status string, message string, data interface{}) Response {
	res := Response{
		Status:  status,
		Message: message,
		Errors:  nil,
		Data:    data,
	}
	return res
}

// BuildErrorResponse method is to inject data value to dynamic failed response
func BuildErrorResponse(message string, err string, data interface{}) Response {
	res := Response{
		Status:  StatusFailure,
		Message: message,
		Errors:  err,
		Data:    data,
	}
	return res
}

// BuildSuccessResponse method is to inject data value to dynamic response success
func BuildSuccessResponse(message string, data interface{}) Response {
	res := Response{
		Status:  StatusSuccess,
		Message: message,
		Errors:  nil,
		Data:    data,
	}
	return res
}
