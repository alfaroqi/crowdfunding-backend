package helper

import (
	"fmt"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
)

type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type Meta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
}

func ApiResponse(message string, code int, status string, data interface{}) Response {
	meta := Meta{
		Message: message,
		Code:    code,
		Status:  status,
	}
	jsonResponse := Response{
		Meta: meta,
		Data: data,
	}
	return jsonResponse
}

// format error validasi
func FormatValitationError(err error) []string {
	var errors []string

	for _, err := range err.(validator.ValidationErrors) {
		errors = append(errors, err.Error())
	}

	return errors
}

func TimeNowMil() string {
	timeMil := time.Now().UnixNano() / int64(time.Millisecond)
	return strconv.Itoa(int(timeMil))
}

func GenCodeTransaction(userID int) string {
	input := fmt.Sprintf("TRX-%s-%s", strconv.Itoa(userID), TimeNowMil())
	return input
}
