package application

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"go.uber.org/zap"

	"market/internal/app/v1/index"
	"market/internal/app/v1/product"
	"market/internal/app/v1/user"
	"market/internal/pkg/auth/jwtauth"
	"market/internal/pkg/config"
	gormApi "market/internal/pkg/database/gorm"
	"market/internal/pkg/logging"
	"market/internal/pkg/middleware"
)

// Application 应用
type Application struct {
	logger     *zap.Logger
	config     *config.Config
	db         *gorm.DB
	auth       *jwtauth.JwtAuth
	router     *gin.Engine
	httpServer *http.Server
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

	db, err := gormApi.New(&gormApi.Options{
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

	// logger
	r.Use(middleware.LoggerMiddleware(logger))

	// Application
	application := &Application{
		logger: logger.With(zap.String("type", "Application")),
		config: conf,
		db:     db,
		auth:   auth,
		router: r,
	}

	if application.configureApps() != nil {
		return nil, errors.Wrap(err, "App配置错误")
	}

	return application, nil

}

func (a *Application) configureApps() error {
	// Repository
	indexRepository := index.NewRepository(a.db)
	userRepository := user.NewRepository(a.db)
	productRepository := product.NewRepository(a.db)

	// Handler
	indexHandler := index.NewHandler(indexRepository, a.auth)
	userHandler := user.NewHandler(userRepository, a.auth)
	productHandler := product.NewHandler(productRepository, a.auth)

	// Initial router
	indexRouter := index.NewRouter(indexHandler)
	userRouter := user.NewRouter(userHandler)
	productRouter := product.NewRouter(productHandler)

	// API router Group
	apiGroup := a.router.Group("/api")
	v1Group := apiGroup.Group("/v1")
	{
		indexRouter(v1Group)
		userRouter(v1Group)
		productRouter(v1Group)
	}
	return nil
}

// Start 启动应用
func (a *Application) Start() error {
	a.httpServer = &http.Server{
		Addr:           fmt.Sprintf("%s:%d", a.config.HttpHost, a.config.HttpPort),
		Handler:        a.router,
		ReadTimeout:    a.config.ReadTimeout * time.Second,
		WriteTimeout:   a.config.WriteTimeout * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	a.logger.Info("Http server starting ...", zap.String("addr", a.config.Addr()))

	go func() {
		if err := a.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			a.logger.Fatal("Start http server err", zap.Error(err))
		}
		return
	}()
	return nil
}

// Stop 停止应用
func (a *Application) Stop() error {
	a.logger.Info("Stopping http server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second) // 平滑关闭, 等待5秒钟处理
	defer cancel()
	if err := a.httpServer.Shutdown(ctx); err != nil {
		return errors.Wrap(err, "Shutdown http server error")
	}
	a.logger.Info("Server exiting ...")
	return nil
}

// AwaitSignal 等待信号
func (a *Application) AwaitSignal() {
	c := make(chan os.Signal, 1)
	signal.Reset(syscall.SIGTERM, syscall.SIGINT)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)
	select {
	case s := <-c:
		a.logger.Info("Receive a signal", zap.String("signal", s.String()))
		if a.httpServer != nil {
			if err := a.Stop(); err != nil {
				a.logger.Warn("Stop http server error", zap.Error(err))
			}
		}
		os.Exit(0)
	}
}
