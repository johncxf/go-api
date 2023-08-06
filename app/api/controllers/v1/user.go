package v1

import (
	"gin-practice/app/api/services"
	"gin-practice/common/response"
	"github.com/gin-gonic/gin"
)

func Info(c *gin.Context) {
	err, user := services.UserService.GetUserInfo(c.Keys["id"].(string))
	if err != nil {
		response.BusinessFail(c, err.Error())
		return
	}
	response.Success(c, user)
}
