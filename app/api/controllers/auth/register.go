package auth

import (
	"gin-practice/app/api/requests"
	"gin-practice/app/api/services"
	"gin-practice/common/response"
	"github.com/gin-gonic/gin"
)

// Register 用户注册
func Register(c *gin.Context) {
	var form requests.Register
	if err := c.ShouldBindJSON(&form); err != nil {
		response.ValidateFail(c, requests.GetErrorMsg(form, err))
		return
	}

	if err, user := services.UserService.Register(form); err != nil {
		response.BusinessFail(c, err.Error())
	} else {
		response.Success(c, user)
	}
}
