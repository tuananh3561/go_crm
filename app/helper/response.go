package helper

import "strings"

const (
	STATUS_SUCCESS = "success"
	STATUS_FAILURE = "fail"
)

//Response is used for static shape json return
type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors"`
	Data    interface{} `json:"data"`
}

//EmptyObj value object is used when data doesnt want to be null on json
type EmptyObj struct{}

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
	splitError := strings.Split(err, "\n")
	res := Response{
		Status:  STATUS_FAILURE,
		Message: message,
		Errors:  splitError,
		Data:    data,
	}
	return res
}

//BuildErrorMessageResponse method is to inject data value to dynamic failed response
func BuildErrorMessageResponse(message string, err error, data interface{}) Response {
	res := Response{
		Status:  STATUS_FAILURE,
		Message: message,
		Errors:  err,
		Data:    data,
	}
	return res
}

// BuildResponseSuccess method is to inject data value to dynamic response success
func BuildResponseSuccess(message string, data interface{}) Response {
	res := Response{
		Status:  STATUS_SUCCESS,
		Message: message,
		Errors:  nil,
		Data:    data,
	}
	return res
}

// BuildSuccessResponse method is to inject data value to dynamic response success
func BuildSuccessResponse(message string, data interface{}) Response {
	res := Response{
		Status:  STATUS_SUCCESS,
		Message: message,
		Errors:  nil,
		Data:    data,
	}
	return res
}
