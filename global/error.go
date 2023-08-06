package global

type CustomError struct {
	ErrorCode int
	ErrorMsg  string
}

type CustomErrors struct {
	BusinessError CustomError
	ValidateError CustomError
	AuthError     CustomError
}

var Errors = CustomErrors{
	BusinessError: CustomError{10001, "业务错误"},
	ValidateError: CustomError{11001, "请求参数错误"},
	AuthError:     CustomError{40001, "登录授权失效"},
}
