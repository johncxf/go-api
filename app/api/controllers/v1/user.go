package v1

import (
	"github.com/gin-gonic/gin"
	"go-api/app/api/services"
	"go-api/common/response"
)

func Info(c *gin.Context) {
	err, user := services.UserService.GetUserInfo(c.Keys["id"].(string))
	if err != nil {
		response.BusinessFail(c, err.Error())
		return
	}
	response.Success(c, user)
}
