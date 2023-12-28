package v1

import (
	"git.legchelife.ru/root/template/internal/controller/http/v1/module/user"
	"git.legchelife.ru/root/template/internal/usecase"
	"git.legchelife.ru/root/template/pkg/logger"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"net/http"
	"time"

	// Swagger docs.
	_ "git.legchelife.ru/root/template/docs"
)

type EmptyBody struct{}

// NewRouter -.
// Swagger spec:
// @title       TEMPLATE
// @description Шаблон для создания нового сервиса
// @version     $(VERSION)

// @host        localhost:1000
// @BasePath    /v1

// @securityDefinitions.apikey Token-A
// @in header
// @name Token-A

func NewRouter(
	handler *gin.Engine,
	l logger.Interface,
	usecase uc.Repo,
) {
	// Options
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())
	handler.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowHeaders:     []string{"Origin", "Token-A", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Token-A", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Swagger
	swaggerHandler := ginSwagger.DisablingWrapHandler(swaggerFiles.Handler, "DISABLE_SWAGGER_HTTP_HANDLER")
	handler.GET("/swagger/*any", swaggerHandler)

	// K8s probe
	handler.GET("/healthz", func(c *gin.Context) { c.Status(http.StatusOK) })

	// Prometheus metrics
	handler.GET("/metrics", gin.WrapH(promhttp.Handler()))

	// Routers
	h := handler.Group("/v1")
	{

		//User
		rt_user.Routes(h, usecase, l)
	}
}
