package routers

import (
	"github.com/gin-gonic/gin"

	v1 "market/api/v1"
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
		userApi.POST("tokens", v1.GetToken)
	}

	accountApi := r.Group("/account", middleware.Jwt())
	{
		accountApi.GET("addresses", v1.GetAddresses)
		accountApi.POST("addresses", v1.CreateAddress)
		accountApi.PUT("addresses/:address_id", v1.UpdateAddress)
		accountApi.DELETE("addresses/:address_id", v1.DeleteAddress)
	}

	categoryApi := r.Group("/categories")
	{
		categoryApi.GET("", v1.GetCategories)
		categoryApi.GET("/:cat_id/specs", v1.GetCategorySpecs)
	}

	return r
}
