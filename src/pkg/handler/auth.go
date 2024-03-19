package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mauiwaui024/todo-golang"
)

func (h *Handler) signUp(c *gin.Context) {
	var input todo.User
	//1.парсим тело запроса
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	//2.передаем данные на слой ниже в сервис
	id, err := h.services.Autorization.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{"id": id})
}

// Для аутентификации структура юзера не подойдет
type signInInput struct {
	Username string `json:"username" binding: "required"`
	Password string `json:"password" binding: "required"`
}

func (h *Handler) signIn(c *gin.Context) {
	var input signInInput
	//1.парсим тело запроса
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	//2.передаем данные на слой ниже в сервис
	token, err := h.services.Autorization.GenerateToken(input.Username, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{"token": token})
}
