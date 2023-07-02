package rt_module_component

import (
	uc "git.legchelife.ru/root/template/internal/usecase"
	"git.legchelife.ru/root/template/pkg/logger"
	"git.legchelife.ru/root/template/pkg/security"
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
