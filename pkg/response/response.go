package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func Success(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, Response{
		Code:    StatusSuccess,
		Message: "成功",
		Data:    map[string]string{},
	})
}

func SuccessMsg(ctx *gin.Context, msg string) {
	ctx.JSON(http.StatusOK, Response{
		Code:    StatusSuccess,
		Message: msg,
		Data:    map[string]string{},
	})
}

func SuccessData(ctx *gin.Context, data any) {
	ctx.JSON(http.StatusOK, Response{
		Code:    StatusSuccess,
		Message: "成功",
		Data:    data,
	})
}

func Error(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, Response{
		Code:    StatusError,
		Message: "失败",
		Data:    map[string]string{},
	})
}

func ErrorMsg(ctx *gin.Context, msg string) {
	ctx.JSON(http.StatusOK, Response{
		Code:    StatusError,
		Message: msg,
		Data:    map[string]string{},
	})
}

func ErrorCode(ctx *gin.Context, code int) {
	ctx.JSON(http.StatusOK, Response{
		Code:    code,
		Message: GetMsg(code),
		Data:    map[string]string{},
	})
}

func ErrorCodeMsg(ctx *gin.Context, code int, msg string) {
	ctx.JSON(http.StatusOK, Response{
		Code:    code,
		Message: msg,
		Data:    map[string]string{},
	})
}

func Failed(ctx *gin.Context, msg string) {
	ctx.JSON(http.StatusOK, Response{
		Code:    StatusErrNot,
		Message: msg,
		Data:    map[string]string{},
	})
	ctx.Abort()
}

func FailedData(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, Response{
		Code:    StatusErrData,
		Message: "请求数据错误",
		Data:    map[string]string{},
	})
}
