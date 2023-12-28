package rt_user

import (
	"git.legchelife.ru/root/template/internal/repo/db/user"
	"git.legchelife.ru/root/template/pkg/models/context"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// Get Method Get
// @Description Метод получения пользователя по id
// @Tags User
// @Accept  json
// @Produce  json
// @Param id query int true "ID пользователя"
// @Success 200 {object} db_user.User{id=int}
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /user/get [get]
func (r *Router) Get(c *gin.Context) {
	ctx := context.Base{}.Create(c)

	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err = r.vGet(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	re, err := r.usecase.Module.User.Get(ctx, id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	} else {
		c.JSON(http.StatusOK, re)
	}
}

// GetList Method GetList
// @Description Метод получения списка пользователей
// @Tags User
// @Accept  json
// @Produce  json
// @Param filter body db_user.UserFilter false "Фильтр"
// @Success 200 {array} []db_user.User{id=int}
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /user/get.list [get]
func (r *Router) GetList(c *gin.Context) {
	ctx := context.Base{}.Create(c)

	filter := db_user.UserFilter{}
	if err := c.ShouldBindJSON(&filter); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err := r.vGetList(filter)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	re, err := r.usecase.Module.User.GetList(ctx, filter)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	} else {
		c.JSON(http.StatusOK, re)
	}
}

// Post Method Post
// @Description Метод создания пользователя
// @Tags User
// @Accept  json
// @Produce  json
// @Param user body db_user.User true "Пользователь"
// @Success 200 {object} db_user.User{id=int}
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /user/post [post]
func (r *Router) Post(c *gin.Context) {
	ctx := context.Base{}.Create(c)

	user := db_user.User{}
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err := r.vPost(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	re, err := r.usecase.Module.User.Post(ctx, user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	} else {
		c.JSON(http.StatusOK, re)
	}
}

// Put Method Put
// @Description Метод обновления пользователя
// @Tags User
// @Accept  json
// @Produce  json
// @Param user body db_user.User{id=int} true "Пользователь"
// @Success 200 {boolean} boolean
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /user/put [put]
func (r *Router) Put(c *gin.Context) {
	ctx := context.Base{}.Create(c)

	user := db_user.User{}
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err := r.vPut(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	re, err := r.usecase.Module.User.Put(ctx, &user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	} else {
		c.JSON(http.StatusOK, re)
	}
}

// Delete Method Delete
// @Description Метод удаления пользователя
// @Tags User
// @Accept  json
// @Produce  json
// @Param id query int true "ID пользователя"
// @Success 200 {boolean} boolean
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /user/delete [delete]
func (r *Router) Delete(c *gin.Context) {
	ctx := context.Base{}.Create(c)

	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err = r.vDelete(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	re, err := r.usecase.Module.User.Delete(ctx, id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	} else {
		c.JSON(http.StatusOK, re)
	}
}
