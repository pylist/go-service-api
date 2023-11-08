package response

import (
	"go-service-api/pkg/response/code"
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
		Code:    code.OK,
		Message: "成功",
		Data:    map[string]string{},
	})
}

func SuccessMsg(ctx *gin.Context, msg string) {
	ctx.JSON(http.StatusOK, Response{
		Code:    code.OK,
		Message: msg,
		Data:    map[string]string{},
	})
}

func SuccessData(ctx *gin.Context, data any) {
	ctx.JSON(http.StatusOK, Response{
		Code:    code.OK,
		Message: "成功",
		Data:    data,
	})
}

func Failed(ctx *gin.Context, msg string) {
	ctx.JSON(http.StatusOK, Response{
		Code:    code.ErrNot,
		Message: msg,
		Data:    map[string]string{},
	})
	ctx.Abort()
}

func FailedData(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, Response{
		Code:    code.ErrData,
		Message: "请求数据错误",
		Data:    map[string]string{},
	})
}
