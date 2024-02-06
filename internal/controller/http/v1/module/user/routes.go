package rt_user

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

	//USER
	h = handler.Group("/user", jwt.SecurityJWT(section))
	{
		h.GET("/get", r.Get)
		h.POST("/get.list", r.GetList)
		h.POST("/post", r.Post)
		h.PUT("/put", r.Put)
		h.DELETE("/delete", r.Delete)
	}
	//USER PARAM
	h = handler.Group("/user/param", jwt.SecurityJWT(section))
	{
		h.POST("/get.list", r.GetParamList)
	}
	//USER TOKEN
	h = handler.Group("/user/token")
	{
		h.GET("/get", r.TokenGet)
	}

	return h
}
