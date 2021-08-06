package response

import (
	"github.com/labstack/echo/v4"
)

type response struct {
	MessageType string      `json:"message_type,omitempty" xml:"message_type,omitempty"`
	Message     string      `json:"message,omitempty" xml:"message,omitempty"`
	Error       bool        `json:"error,omitempty" xml:"error,omitempty"`
	Data        interface{} `json:"data,omitempty" xml:"data,omitempty"`
}

func New(
	ctx echo.Context,
	codeHttp int,
	message string,
	err bool,
	data interface{},
) error {
	var messageType string
	if err {
		messageType = "error"
	} else {
		messageType = "message"
	}

	res := &response{
		MessageType: messageType,
		Message:     message,
		Error:       err,
		Data:        data,
	}

	switch ctx.Request().Header.Get("Content-Type") {
	case "application/json":
		return ctx.JSON(codeHttp, res)
	case "application/xml":
		return ctx.XML(codeHttp, res)
	default:
		return ctx.JSON(codeHttp, res)
	}
}
