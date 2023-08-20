package auth

import (
	"go-api/app/api/services"
	"go-api/common/response"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Logout(c *gin.Context) {
	err := services.JwtService.JoinBlackList(c.Keys["token"].(*jwt.Token))
	if err != nil {
		response.BusinessFail(c, "登出失败")
		return
	}
	response.Success(c, nil)
}
