package entities

// Response struct is built for dynamic response. key status, message, data are mandatory
type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Error   interface{} `json:"error"`
	Data    interface{} `json:"data"`
}

// BuildResponse method is to inject data value to dynamic response
func BuildResponse(status bool, message string, data interface{}) Response {
	res := Response{Status: status, Message: message, Error: nil, Data: data}
	return res
}

// BuildErrorResponse method is to show response failed of the request.
func BuildErrorResponse(message string, error interface{}, data interface{}) Response {
	res := Response{Status: false, Message: message, Error: error, Data: data}
	return res
}
