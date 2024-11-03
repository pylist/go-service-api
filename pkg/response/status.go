package response

const (
	StatusSuccess          = 0
	StatusError            = 10000
	StatusErrData          = 50001
	StatusErrNot           = 50002
	StatusErrAuth          = 40001
	StatusErrTokenNotFound = 40002
	StatusErrTokenExpired  = 40003
	StatusErrTokenInvalid  = 40004
)

var msgMap = map[int]string{
	StatusSuccess:          "成功",
	StatusError:            "失败",
	StatusErrData:          "请求数据错误",
	StatusErrNot:           "请求失败",
	StatusErrAuth:          "认证失败",
	StatusErrTokenNotFound: "token未找到",
	StatusErrTokenInvalid:  "token无效",
	StatusErrTokenExpired:  "token过期",
}

func GetMsg(code int) string {
	if msg, ok := msgMap[code]; ok {
		return msg
	}
	return "未知错误"
}
