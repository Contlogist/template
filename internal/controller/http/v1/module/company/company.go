package rt_company

import (
	"git.legchelife.ru/root/template/pkg/logger"
	"git.legchelife.ru/root/template/pkg/models/context"
	"git.legchelife.ru/root/template/pkg/models/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Post Method POST
// @Summary  Получение компании
// @Description Метод получения компании - (payload)
// @Description ID компании берется из payload
// @Tags Company
// @Accept json
// @Produce json
// @Security Authorization
// @Success 200 {object} response.Base{data=user.User{id=int},error=interface{}}
// @Failure 400 {object} response.Base{data=interface{},error=string}
// @Failure 500 {object} response.Base{data=interface{},error=string}
// @Router /user/get [get]
func (r *Router) Post(c *gin.Context) {
	ctx := context.Base{}.Create(c)

	re, err := r.usecase.UserRepo.User.Get(&ctx, ctx.Payload.ID)
	if err != nil {
		r.l.Error(err, logger.Data{Module: "User", Method: "Get", Action: logger.Usecase, Params: map[string]interface{}{"id": ctx.Payload.ID}})
		c.JSON(http.StatusInternalServerError, response.Base{Data: nil, Error: err.Error()})
		return
	} else {
		r.l.Info("Get - OK", logger.Data{Module: "User", Method: "Get", Action: logger.Usecase, Params: map[string]interface{}{"id": ctx.Payload.ID}})
		c.JSON(http.StatusOK, response.Base{Data: re, Error: nil})
	}
}
