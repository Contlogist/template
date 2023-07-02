package rt_module_component

import (
	"git.legchelife.ru/root/template/pkg/models/context"
	"git.legchelife.ru/root/template/pkg/models/response"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

// @Summary     getTest - получить категории
// @Description Метод для получения списка категорий
// @ID          module.block.get
// @Tags  	    Catalog/Category
// @Security 	Token-A
// @Produce     json
// @Success     200 {object} string "Возвращает список категорий"
// @Failure     500 {object} string "Возвращает ошибку"
// @Router      /module/block/get/test [get]
func (r *Router) getTest(c *gin.Context) {
	ctx := context.Base{}.Create(c)

	logrus.Info("getTest")

	res, err := r.usecase.Module.Component.GetTest(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Base{Error: err.Error()})
		return
	} else {
		c.JSON(http.StatusOK, response.Base{Data: res})
	}
}
