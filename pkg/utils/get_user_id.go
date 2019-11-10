package utils

import (
	"errors"
	"github.com/gin-gonic/gin"
)

// 获取用户ID
func GetUserID(ctx *gin.Context) (uint, error) {

	if userId, exist := ctx.Get("userId"); exist {
		return userId.(uint), nil
	} else {
		return 0, errors.New("授权中间件异常")
	}
}
