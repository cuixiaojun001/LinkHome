package response

import (
	"errors"
	"net/http"

	"github.com/cuixiaojun001/LinkHome/common/errdef"
)

const (
	success             = 0
	failed              = -1
	badRequest          = 100000
	unauthorized        = 1
	internalServerError = 300000 // internalServerError 接口异常 前端只显示内部错误
)

type JSONResponse struct {
	Code int         `json:"code"` // Code 100000直接展示message，200000为成功，300000为服务器错误，展示系统异常
	Msg  string      `json:"message"`
	Data interface{} `json:"data"`
}

func Success(data interface{}) *JSONResponse {
	return &JSONResponse{
		Code: success,
		Msg:  "SUCCESS",
		Data: data,
	}
}

func Failed(data interface{}) *JSONResponse {
	return &JSONResponse{
		Code: failed,
		Msg:  "FAILED",
		Data: data,
	}
}

func BusinessException(err error) *JSONResponse {
	var customErr *errdef.CustomError
	if errors.As(err, &customErr) {
		return &JSONResponse{
			Code: customErr.Code,
			Msg:  customErr.Message,
			Data: nil,
		}
	}
	return &JSONResponse{
		Code: -1,
		Msg:  err.Error(),
		Data: nil,
	}
}

func BadRequest(err error) *JSONResponse {
	return &JSONResponse{
		Code: badRequest,
		Msg:  http.StatusText(http.StatusBadRequest),
		Data: err.Error(),
	}
}

func Unauthorized(err error) *JSONResponse {
	return &JSONResponse{
		Code: unauthorized,
		Msg:  http.StatusText(http.StatusUnauthorized),
		Data: err.Error(),
	}
}

func InternalServerError(err error) *JSONResponse {
	return &JSONResponse{
		Code: internalServerError,
		Msg:  http.StatusText(http.StatusInternalServerError),
		Data: err.Error(),
	}
}
