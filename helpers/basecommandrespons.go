package helpers

type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type Meta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
}

func BaseCommandResponse(message string, code int, status string, data interface{}) Response {
	metax := Meta{
		Message: message,
		Code:    code,
		Status:  status,
	}

	baseresponse := Response{
		Meta: metax,
		Data: data,
	}

	return baseresponse

}
