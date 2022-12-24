package handler

import (
	"gotodo/todo"
	"gotodo/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type todoHandler struct {
	serviceTodo todo.ITodoService
}

// new handler
func NewTodoHandler(serviceTodo todo.ITodoService) *todoHandler {
	return &todoHandler{serviceTodo}
}

func (h *todoHandler) CreateTodo(c *gin.Context) {
	var input todo.TodoCreateInput

	// bind
	if err := c.BindJSON(&input); err != nil {
		myErr := utils.ErrorBinding(err)
		response := utils.ResponseAPI(
			http.StatusText(http.StatusBadRequest),
			nil,
			myErr,
			true,
		)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	// call service
	newTodo, err := h.serviceTodo.CreateTodo(input)
	if err != nil {
		response := utils.ResponseAPI(
			http.StatusText(http.StatusBadRequest),
			nil,
			err.Error(),
			true,
		)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	todoFormatted := todo.FormatterTodo(newTodo)

	response := utils.ResponseAPI(
		"Success",
		todoFormatted,
		"Success",
		false,
	)

	c.JSON(http.StatusOK, response)
}
