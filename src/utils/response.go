package utils

const (
	ERR = "error"
	MSG = "message"
)

type response struct {
	MessageType string      `json:"message_type" xml:"message_type"`
	Message     string      `json:"message" xml:"message"`
	Error       bool        `json:"error" xml:"error"`
	Data        interface{} `json:"data" xml:"data"`
}

func Response(
	MessageType string,
	Message string,
	Error bool,
	Data interface{},
) *response {
	return &response{
		MessageType,
		Message,
		Error,
		Data,
	}
}

func ResponseErr(
	Message string,
	Data interface{},
) *response {
	return &response{
		ERR,
		Message,
		true,
		Data,
	}
}

func ResponseSatisfactory(
	Message string,
	Data interface{},
) *response {
	return &response{
		MSG,
		Message,
		false,
		Data,
	}
}
