package utils

const (
	MESSAGE = "message"
	ERR = "error"
)

type response struct {
	TypeMessage string `bson:"type_message" json:"type_message" xml:"type_message"`
	Message string `bson:"message" json:"message" xml:"message"`
	Error bool `bson:"error" json:"error" xml:"error"`
	Data interface{} `bson:"data" json:"data" xml:"data"`
}

func ResponseError(err error) *response {
	return &response{
		TypeMessage: ERR,
		Message: err.Error(),
		Error: true,
		Data: nil,
	}
}

func ResponseSatisfactory(message string, data interface{}) *response {
	return &response{
		TypeMessage: MESSAGE,
		Message: message,
		Error: false,
		Data: data,
	}
}