package requests

type Register struct {
	Username string `form:"username" json:"username" binding:"required"`
	Email    string `from:"email" json:"email" binding:"required,email"`
	Password string `form:"password" json:"password" binding:"required"`
}

type Login struct {
	Email    string `from:"email" json:"email" binding:"required,email"`
	Password string `form:"password" json:"password" binding:"required"`
}

func (register Register) GetMessages() ValidatorMessages {
	return ValidatorMessages{
		"username.required": "用户名称不能为空",
		"email.required":    "邮箱不能为空",
		"email.email":       "邮箱格式错误",
		"password.required": "用户密码不能为空",
	}
}

func (login Login) GetMessages() ValidatorMessages {
	return ValidatorMessages{
		"email.required":    "邮箱不能为空",
		"email.email":       "邮箱格式错误",
		"password.required": "用户密码不能为空",
	}
}
