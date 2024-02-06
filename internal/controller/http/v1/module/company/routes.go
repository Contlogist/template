package rt_company

import (
	uc "git.legchelife.ru/root/template/internal/usecase"
	"git.legchelife.ru/root/template/pkg/logger"
	"git.legchelife.ru/root/template/pkg/security/jwt"
	"github.com/gin-gonic/gin"
)

const (
	section = 1
)

type Router struct {
	usecase uc.Repo
	l       logger.Interface
}

func Routes(handler *gin.RouterGroup, usecase uc.Repo, l logger.Interface) (h *gin.RouterGroup) {
	r := &Router{usecase, l}

	h = handler.Group("/company", jwt.SecurityJWT(section))

	{
		h.POST("/post", r.Post)
		//h.GET("/get", r.Get)
		//h.POST("/get.list", r.GetList)
		//h.PUT("/put", r.Put)
		//h.DELETE("/delete", r.Delete)

	}

	return h
}
