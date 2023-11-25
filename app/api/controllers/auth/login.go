package auth

import (
	"github.com/gin-gonic/gin"
	"go-api/app/api/requests"
	"go-api/app/api/services"
	"go-api/common/response"
)

func Login(c *gin.Context) {
	var form requests.Login
	if err := c.ShouldBindJSON(&form); err != nil {
		response.ValidateFail(c, requests.GetErrorMsg(form, err))
		return
	}

	if err, user := services.UserService.Login(form); err != nil {
		response.BusinessFail(c, err.Error())
	} else {
		tokenData, err, _ := services.JwtService.CreateToken(services.AppGuardName, user)
		if err != nil {
			response.BusinessFail(c, err.Error())
			return
		}
		response.Success(c, tokenData)
	}
}
