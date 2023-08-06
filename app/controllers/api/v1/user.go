package v1

import (
	request2 "gin-practice/app/requests"
	"gin-practice/app/services"
	"gin-practice/common/response"
	"github.com/gin-gonic/gin"
)

// Register 用户注册
func Register(c *gin.Context) {
	var form request2.Register
	if err := c.ShouldBindJSON(&form); err != nil {
		response.ValidateFail(c, request2.GetErrorMsg(form, err))
		return
	}

	if err, user := services.UserService.Register(form); err != nil {
		response.BusinessFail(c, err.Error())
	} else {
		response.Success(c, user)
	}
}
