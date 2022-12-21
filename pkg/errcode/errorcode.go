package errorcode

import (
	"fmt"
	"net/http"
)

// ErrorInfo 錯誤格式
type ErrorInfo struct {
	code    int      `json:"code"`
	message string   `json:"message"`
	details []string `json:"details"`
}

var codes = map[int]string{}

// NewError 定義新錯誤
func NewError(code int, msg string) *ErrorInfo {
	if _, ok := codes[code]; ok {
		panic(fmt.Sprintf("This error has been defined, code: %d, message: %v", code, msg))
	}
	codes[code] = msg
	return &ErrorInfo{code: code, message: msg}
}

// Error 取得錯誤的錯誤代碼與錯誤訊息
func (e *ErrorInfo) Error() string {
	return fmt.Sprintf("Error occured, error code: %d, message: %v", e.Code(), e.Message())
}

// Code 取得錯誤的錯誤代碼
func (e *ErrorInfo) Code() int {
	return e.code
}

// Ｍessage 取得錯誤的錯誤訊息
func (e *ErrorInfo) Message() string {
	return e.message
}

// ＭessageF 取得錯誤的錯誤訊息可自定義後面的錯誤資訊
func (e *ErrorInfo) MessageF(args []interface{}) string {
	return fmt.Sprintf(e.message, args...)
}

// Details 取得錯誤的錯誤詳細資訊
func (e *ErrorInfo) Details() []string {
	return e.details
}

func (e *ErrorInfo) WithDetails(details ...string) *ErrorInfo {
	newError := *e
	newError.details = details
	return &newError
}

func (e *ErrorInfo) StatusCode() int {
	switch e.Code() {
	case Success.Code():
		return http.StatusOK
	case ServerError.Code():
		return http.StatusInternalServerError
	case InvaildParams.Code():
		return http.StatusBadRequest
	case UnauthorizedAuthNotExist.Code():
		fallthrough
	case UnauthorizedToken.Code():
		fallthrough
	case UnauthorizedTokenGenerate.Code():
		fallthrough
	case TooManyRequest.Code():
		return http.StatusTooManyRequests
	}

	return http.StatusInternalServerError
}
