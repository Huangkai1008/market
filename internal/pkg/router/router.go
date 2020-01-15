package router

import "github.com/gin-gonic/gin"

type Router func(*gin.Engine)

type Group func(*gin.RouterGroup)
