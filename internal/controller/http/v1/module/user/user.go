package rt_user

import (
	user "git.legchelife.ru/root/template/internal/repo/user"
	"git.legchelife.ru/root/template/pkg/logger"
	"git.legchelife.ru/root/template/pkg/models/context"
	"git.legchelife.ru/root/template/pkg/models/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Get Method Get
// @Summary  Получение пользователя
// @Description Метод получения пользователя - (payload)
// @Description ID пользователя берется из payload
// @Tags User
// @Accept  json
// @Produce  json
// @Security 	Authorization
// @Success 200 {object} response.Base{data=user.User{id=int},error=interface{}}
// @Failure 400 {object} response.Base{data=interface{},error=string}
// @Failure 500 {object} response.Base{data=interface{},error=string}
// @Router /user/get [get]
func (r *Router) Get(c *gin.Context) {
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

// GetList Method GetList
// @Summary Получение списка пользователей
// @Description Метод получения списка пользователей - (filter + payload)
// @Description Фильтр основывается на payload(cid)
// @Tags User
// @Accept json
// @Produce json
// @Security Authorization
// @Param filter body user.UserFilter false "Фильтр"
// @Success 200 {array} response.Base{data=[]user.User,error=interface{}}
// @Failure 400 {object} response.Base{data=interface{},error=string}
// @Failure 500 {object} response.Base{data=interface{},error=string}
// @Router /user/get.list [post]
func (r *Router) GetList(c *gin.Context) {
	ctx := context.Base{}.Create(c)

	filter := user.UserFilter{}
	if err := c.ShouldBindJSON(&filter); err != nil {
		r.l.Error(err, logger.Data{Module: "User", Method: "GetList", Action: logger.Parse, Params: map[string]interface{}{"filter": filter}})
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err := r.vGetList(filter)
	if err != nil {
		r.l.Error(err, logger.Data{Module: "User", Method: "GetList", Action: logger.Validate, Params: map[string]interface{}{"filter": filter}})
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	re, err := r.usecase.UserRepo.User.GetList(&ctx, filter)

	if err != nil {
		r.l.Error(err, logger.Data{Module: "User", Method: "GetList", Action: logger.Usecase, Params: map[string]interface{}{"filter": filter}})
		c.JSON(http.StatusInternalServerError, err.Error())
	} else {
		r.l.Info("GetList - OK", logger.Data{Module: "User", Method: "GetList", Action: logger.Usecase, Params: map[string]interface{}{"filter": filter}})
		c.JSON(http.StatusOK, response.Base{Data: re, Error: nil})
	}
}

// Post Method Post
// @Summary Создание пользователя
// @Description Метод создания пользователя, используется при регистрации компании (создается администратор)),
// @Description а так же когда администратор создает пользователя в своей компании
// @Description CID (ID компании) берется из payload (из токена администратора)
// @Tags User
// @Accept  json
// @Produce  json
// @Security Authorization
// @Param user body user.User true "Пользователь"
// @Success 200 {object} response.Base{data=int,error=interface{}}
// @Failure 400 {object} response.Base{data=interface{},error=string}
// @Failure 500 {object} response.Base{data=interface{},error=string}
// @Router /user/post [post]
func (r *Router) Post(c *gin.Context) {
	ctx := context.Base{}.Create(c)

	user := user.User{}
	if err := c.ShouldBindJSON(&user); err != nil {
		r.l.Error(err, logger.Data{Module: "User", Method: "Post", Action: logger.Parse, Params: map[string]interface{}{"user": user}})
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err := r.vPost(user)
	if err != nil {
		r.l.Error(err, logger.Data{Module: "User", Method: "Post", Action: logger.Validate, Params: map[string]interface{}{"user": user}})
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	re, err := r.usecase.UserRepo.User.Post(&ctx, user)

	if err != nil {
		r.l.Error(err, logger.Data{Module: "User", Method: "Post", Action: logger.Usecase, Params: map[string]interface{}{"user": user}})
		c.JSON(http.StatusInternalServerError, err.Error())
	} else {
		c.JSON(http.StatusOK, re)
	}
}

// Put Method Put
// @Summary Обновление пользователя
// @Description Метод обновления пользователя
// @Description ID пользователя берется из payload
// @Tags User
// @Accept  json
// @Produce  json
// @Security Authorization
// @Param user body user.User true "Пользователь"
// @Success 200 {object} response.Base{data=boolean,error=interface{}}
// @Failure 400 {object} response.Base{data=interface{},error=string}
// @Failure 500 {object} response.Base{data=interface{},error=string}
// @Router /user/put [put]
func (r *Router) Put(c *gin.Context) {
	ctx := context.Base{}.Create(c)

	user := user.User{}
	if err := c.ShouldBindJSON(&user); err != nil {
		r.l.Error(err, logger.Data{Module: "User", Method: "Put", Action: logger.Parse, Params: map[string]interface{}{"user": user}})
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err := r.vPut(user)
	if err != nil {
		r.l.Error(err, logger.Data{Module: "User", Method: "Put", Action: logger.Validate, Params: map[string]interface{}{"user": user}})
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	re, err := r.usecase.UserRepo.User.Put(&ctx, &user)

	if err != nil {
		r.l.Error(err, logger.Data{Module: "User", Method: "Put", Action: logger.Usecase, Params: map[string]interface{}{"user": user}})
		c.JSON(http.StatusInternalServerError, err.Error())
	} else {
		r.l.Info("Put - OK", logger.Data{Module: "User", Method: "Put", Action: logger.Usecase, Params: map[string]interface{}{"user": user}})
		c.JSON(http.StatusOK, re)
	}
}

// Delete Method Delete
// @Summary Обновление пользователя
// @Description Метод удаления пользователя
// @Description ID пользователя берется из payload
// @Tags User
// @Accept  json
// @Produce  json
// @Security Authorization
// @Success 200 {object} response.Base{data=boolean,error=interface{}}
// @Failure 400 {object} response.Base{data=interface{},error=string}
// @Failure 500 {object} response.Base{data=interface{},error=string}
// @Router /user/delete [delete]
func (r *Router) Delete(c *gin.Context) {
	ctx := context.Base{}.Create(c)

	var id int
	err := c.ShouldBindQuery(&id)
	if err != nil {
		r.l.Error(err, logger.Data{Module: "User", Method: "Delete", Action: logger.Parse, Params: map[string]interface{}{"id": id}})
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err = r.vDelete(id)
	if err != nil {
		r.l.Error(err, logger.Data{Module: "User", Method: "Delete", Action: logger.Validate, Params: map[string]interface{}{"id": id}})
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	re, err := r.usecase.UserRepo.User.Delete(&ctx, id)

	if err != nil {
		r.l.Error(err, logger.Data{Module: "User", Method: "Delete", Action: logger.Usecase, Params: map[string]interface{}{"id": id}})
		c.JSON(http.StatusInternalServerError, err.Error())
	} else {
		r.l.Info("Delete - OK", logger.Data{Module: "User", Method: "Delete", Action: logger.Usecase, Params: map[string]interface{}{"id": id}})
		c.JSON(http.StatusOK, re)
	}
}
