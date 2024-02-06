package rt_user

import (
	"git.legchelife.ru/root/template/pkg/logger"
	"git.legchelife.ru/root/template/pkg/models/context"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// GetParamList Method Get
// @Description Метод получения параметров пользователя по id
// @Tags User/Param
// @Accept  json
// @Produce  json
// @Param id query string true "ID пользователя"
// @Success 200 {object} user.User{id=int}
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /user/param/get.list [get]
func (r *Router) GetParamList(c *gin.Context) {
	ctx := context.Base{}.Create(c)

	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		r.l.Error(err, logger.Data{Module: "User", Method: "Get", Action: logger.Parse, Params: map[string]interface{}{"id": id}})
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err = r.vGet(id)
	if err != nil {
		r.l.Error(err, logger.Data{Module: "User", Method: "Get", Action: logger.Validate, Params: map[string]interface{}{"id": id}})
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	re, err := r.usecase.UserRepo.Param.GetList(&ctx, id)
	if err != nil {
		r.l.Error(err, logger.Data{Module: "User", Method: "Get", Action: logger.Usecase, Params: map[string]interface{}{"id": id}})
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	} else {
		r.l.Info("Get - OK", logger.Data{Module: "User", Method: "Get", Action: logger.Usecase, Params: map[string]interface{}{"id": id}})
		c.JSON(http.StatusOK, re)
	}
}
