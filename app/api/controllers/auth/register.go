package auth

import (
	"github.com/gin-gonic/gin"
	"go-api/app/api/requests"
	"go-api/app/api/services"
	"go-api/common/response"
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
