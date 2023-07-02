package rt_module_component

import (
	"git.legchelife.ru/gitlab-instance-7d441567/catalog_m/internal/usecase"
	"git.legchelife.ru/gitlab-instance-7d441567/catalog_m/pkg/logger"
	"git.legchelife.ru/gitlab-instance-7d441567/catalog_m/pkg/security"
	"github.com/gin-gonic/gin"
)

type Router struct {
	usecase uc.Repo
	l       logger.Interface
}

func NewRoutes(handler *gin.RouterGroup, usecase uc.Repo, l logger.Interface) (h *gin.RouterGroup) {
	r := &Router{usecase, l}

	h = handler.Group("/module", security.TokenA())
	{
		h.GET("/block/get/test", r.getTest)
	}
	return h
}
