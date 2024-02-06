package rt_section

import (
	"git.legchelife.ru/root/template/pkg/logger"
	"git.legchelife.ru/root/template/pkg/models/context"
	"git.legchelife.ru/root/template/pkg/models/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetList Method GetList
// @Summary  Получение списка секций
// @Description Метод получает список секций
// @Description
// @Tags Section
// @Accept json
// @Produce json
// @Security Authorization
// @Success 200 {object} response.Base{data=rp_section.Section{id=int},error=interface{}}
// @Failure 400 {object} response.Base{data=interface{},error=string}
// @Failure 500 {object} response.Base{data=interface{},error=string}
// @Router /section/get.list [get]
func (r *Router) GetList(c *gin.Context) {
	ctx := context.Base{}.Create(c)

	re, err := r.usecase.SectionRepo.GetList(&ctx)
	if err != nil {
		r.l.Error(err, logger.Data{Module: "Section", Method: "GetList", Action: logger.Usecase, Params: map[string]interface{}{"id": ctx.Payload.ID}})
		c.JSON(http.StatusInternalServerError, response.Base{Data: nil, Error: err.Error()})
		return
	} else {
		r.l.Info("Get - OK", logger.Data{Module: "Section", Method: "GetList", Action: logger.Usecase, Params: map[string]interface{}{"id": ctx.Payload.ID}})
		c.JSON(http.StatusOK, response.Base{Data: re, Error: nil})
	}
}
