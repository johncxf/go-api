package auth

import (
	"github.com/gin-gonic/gin"
	"go-api/app/api/requests"
	"go-api/app/api/services"
	"go-api/common/global"
	"go-api/common/response"
	"time"
)

func Login(c *gin.Context) {
	var form requests.Login
	if err := c.ShouldBindJSON(&form); err != nil {
		response.ValidateError(c, requests.GetErrorMsg(form, err))
		return
	}

	if err, user := services.UserService.Login(form); err != nil {
		response.BusinessError(c, err.Error())
	} else {
		tokenData, err, _ := services.JwtService.CreateToken(services.AppGuardName, user)
		if err != nil {
			response.BusinessError(c, err.Error())
			return
		}

		// 记录登陆日志
		loginIP := c.ClientIP()
		if loginIP == "::1" {
			loginIP = "127.0.0.1"
		}
		loginTime := time.Now()
		if err, _ := services.LoginLogService.AddLoginLog(user.GetId(), loginIP, loginTime); err != nil {
			global.Logger.Warn("add login log error: " + err.Error())
		}

		response.Success(c, tokenData)
	}
}
