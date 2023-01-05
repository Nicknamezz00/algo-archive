package errcode

var (
	Success                    = NewError(0, "成功")
	ServerError                = NewError(10000, "内部错误")
	InvalidParameters          = NewError(10001, "入参错误")
	NotFound                   = NewError(10002, "找不到")
	UnauthorizedAuthNotExist   = NewError(10003, "账号不存在")
	UnauthorizedAuthFailed     = NewError(10004, "账号密码错误")
	UnauthorizedTokenError     = NewError(10005, "鉴权失败, Token 错误或丢失")
	UnauthorizedTokenTimeout   = NewError(10006, "鉴权失败, Token 超时")
	UnauthorizedTokenGenFailed = NewError(10007, "鉴权失败, Token 生成失败")
	TooManyRequests            = NewError(10008, "请求过多")

	GatewayMethodsLimit = NewError(10009, "仅接受GET/POST")
)
