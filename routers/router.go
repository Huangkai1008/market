package routers

import (
	"github.com/gin-gonic/gin"
	"market/api/v1"
	"market/pkg/setting"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(setting.RunMode)

	pingApi := r.Group("/ping")
	{
		pingApi.GET("", v1.Ping)
	}
	return r
}
