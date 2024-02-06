package rt_user

import (
	"git.legchelife.ru/root/template/pkg/logger"
	"git.legchelife.ru/root/template/pkg/models/context"
	"github.com/gin-gonic/gin"
	"net/http"
)

// TokenGet Method Get
// @Description Метод получает токен пользователя (аутентификация)
// @Tags User/Token
// @Accept  json
// @Produce  json
// @Param email query string true "email"
// @Param password query string true "password"
// @Success 200 {object} user.Tokens{}
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /user/token/get [get]
func (r *Router) TokenGet(c *gin.Context) {
	ctx := context.Base{}.Create(c)

	email := c.Query("email")
	password := c.Query("password")

	if email == "" || password == "" {
		c.JSON(http.StatusBadRequest, "email или password не заданы")
		return
	}

	re, err := r.usecase.UserRepo.Token.Get(&ctx, email, password)
	if err != nil {
		r.l.Error(err, logger.Data{Module: "User", Method: "TokenGet", Action: logger.Usecase, Params: map[string]interface{}{"email": email, "password": password}})
		c.JSON(http.StatusUnauthorized, err.Error())
		return
	} else {
		r.l.Info("TokenGet - OK", logger.Data{Module: "User", Method: "TokenGet", Action: logger.Usecase, Params: map[string]interface{}{"email": email, "password": password}})
		c.JSON(http.StatusOK, re)
	}
}
