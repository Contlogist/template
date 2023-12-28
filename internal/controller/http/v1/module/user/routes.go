package rt_user

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

func Routes(handler *gin.RouterGroup, usecase uc.Repo, l logger.Interface) (h *gin.RouterGroup) {
	r := &Router{usecase, l}

	h = handler.Group("/user", security.TokenA())
	{
		h.GET("/get", r.Get)
		h.GET("/get.list", r.GetList)
		h.POST("/post", r.Post)
		h.PUT("/put", r.Put)
		h.DELETE("/delete", r.Delete)
	}
	return h
}
