package routers

import (
	"github.com/gin-gonic/gin"
	"market/api/v1"
	"market/middleware"
	"market/pkg/logging"
	"market/pkg/setting"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.GinZap(logging.GinLogger))
	gin.SetMode(setting.RunMode)

	pingApi := r.Group("/ping")
	{
		pingApi.GET("", v1.Ping)
	}

	userApi := r.Group("/users")
	{
		userApi.POST("register", v1.Register)
	}
	return r
}
