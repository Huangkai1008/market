package application

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"go.uber.org/zap"

	"market/internal/app/account"
	"market/internal/app/index"
	"market/internal/app/product"
	"market/internal/app/user"
	"market/internal/pkg/auth/jwtauth"
	"market/internal/pkg/config"
	gorm2 "market/internal/pkg/database/gorm"
	"market/internal/pkg/logging"
	"market/internal/pkg/middleware"
)

// Application 应用
type Application struct {
	Logger *zap.Logger
	Config *config.Config
	DB     *gorm.DB
	Auth   *jwtauth.JwtAuth
	Server *gin.Engine
}

// New 返回一个新的应用实例
func New() (*Application, error) {
	conf, err := config.New()
	if err != nil {
		return nil, errors.Wrap(err, "读取配置错误")
	}

	logger, err := logging.NewWithOptions(&logging.Options{Config: conf})
	if err != nil {
		return nil, errors.Wrap(err, "日志配置错误")
	}

	db, err := gorm2.New(&gorm2.Options{
		Database: &conf.Database,
		Gorm:     &conf.Gorm,
	})
	if err != nil {
		return nil, errors.Wrap(err, "数据库配置错误")
	}

	auth := jwtauth.New(&jwtauth.Options{
		Jwt: &conf.Jwt,
	})

	gin.SetMode(conf.RunMode)
	r := gin.New()

	// Recover
	r.Use(gin.Recovery())

	// Logger
	r.Use(middleware.LoggerMiddleware(logger))

	// Application
	application := &Application{
		Logger: logger,
		Config: conf,
		DB:     db,
		Auth:   auth,
		Server: r,
	}

	if application.configureApps() != nil {
		return nil, errors.Wrap(err, "App配置错误")
	}

	return application, nil

}

func (a *Application) configureApps() error {
	// Repository
	indexRepository := index.NewRepository(a.DB)
	userRepository := user.NewRepository(a.DB)
	accountRepository := account.NewRepository(a.DB)
	productRepository := product.NewRepository(a.DB)

	// Handler
	indexHandler := index.NewHandler(indexRepository, a.Auth)
	userHandler := user.NewHandler(userRepository, a.Auth)
	accountHandler := account.NewHandler(accountRepository, a.Auth)
	productHandler := product.NewHandler(productRepository, a.Auth)

	// Initial Router
	indexRouter := index.NewRouter(indexHandler)
	userRouter := user.NewRouter(userHandler)
	accountRouter := account.NewRouter(accountHandler)
	productRouter := product.NewRouter(productHandler)

	// Register Router
	indexRouter(a.Server)
	userRouter(a.Server)
	accountRouter(a.Server)
	productRouter(a.Server)
	return nil
}

// Start 启动应用
func (a *Application) Start() error {
	server := &http.Server{
		Addr:           fmt.Sprintf(":%d", a.Config.HttpPort),
		Handler:        a.Server,
		ReadTimeout:    a.Config.ReadTimeout * time.Second,
		WriteTimeout:   a.Config.WriteTimeout * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	err := server.ListenAndServe()
	return err
}
