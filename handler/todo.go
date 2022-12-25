package handler

import (
	"gotodo/todo"
	"gotodo/utils"
	"net/http"
	"strconv"

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

	if input.Title == "" {
		response := utils.ResponseAPI(
			STATUS_BAD_REQUEST,
			nil,
			"title cannot be null",
			true,
		)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	if input.ActivityID <= 0 {
		response := utils.ResponseAPI(
			STATUS_BAD_REQUEST,
			nil,
			"activity_group_id cannot be null",
			true,
		)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	// call service
	newTodo, err := h.serviceTodo.CreateTodo(input)
	if err != nil {
		response := utils.ResponseAPI(
			STATUS_BAD_REQUEST,
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

	c.JSON(http.StatusCreated, response)
}

func (h *todoHandler) GetTodoByID(c *gin.Context) {
	// get id from param
	idParam := c.Param("id")

	// convert to int
	id, err := strconv.Atoi(idParam)
	if err != nil {
		response := utils.ResponseAPI(
			http.StatusText(http.StatusBadRequest),
			nil,
			"id must be int",
			true,
		)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	// call service
	todoByID, err := h.serviceTodo.GetTodoById(id)
	if err != nil {
		response := utils.ResponseAPI(
			STATUS_NOT_FOUND,
			nil,
			err.Error(),
			true,
		)

		c.JSON(http.StatusNotFound, response)
		return
	}

	todoFormatted := todo.FormatterTodo(todoByID)

	response := utils.ResponseAPI(
		"Success",
		todoFormatted,
		"Success",
		false,
	)

	c.JSON(http.StatusOK, response)
}

func (h *todoHandler) GetAllTodo(c *gin.Context) {
	idActivityQuery := c.Query("activity_group_id")

	if idActivityQuery == "" {
		// call service
		todos, err := h.serviceTodo.GetAllTodo(0, false)
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

		todoFormatted := todo.FormatterTodos(todos)

		response := utils.ResponseAPI(
			"Success",
			todoFormatted,
			"Success",
			false,
		)

		c.JSON(http.StatusOK, response)
		return
	}

	// convert to int
	id, err := strconv.Atoi(idActivityQuery)
	if err != nil {
		response := utils.ResponseAPI(
			http.StatusText(http.StatusBadRequest),
			nil,
			"id must be int",
			true,
		)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	// call service
	todos, err := h.serviceTodo.GetAllTodo(id, true)
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

	var resp interface{}
	todoFormatted := todo.FormatterTodos(todos)
	resp = todoFormatted

	if len(todoFormatted) == 0 {
		resp = []todo.Todo{}
	}

	response := utils.ResponseAPI(
		"Success",
		resp,
		"Success",
		false,
	)

	c.JSON(http.StatusOK, response)
}

func (h *todoHandler) DeleteTodoById(c *gin.Context) {
	// get id from param
	idParam := c.Param("id")

	// convert to int
	id, err := strconv.Atoi(idParam)
	if err != nil {
		response := utils.ResponseAPI(
			http.StatusText(http.StatusBadRequest),
			nil,
			"id must be int",
			true,
		)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	// call service
	if err := h.serviceTodo.DeleteByID(id); err != nil {
		response := utils.ResponseAPI(
			STATUS_NOT_FOUND,
			nil,
			err.Error(),
			true,
		)

		c.JSON(http.StatusNotFound, response)
		return
	}

	response := utils.ResponseAPI(
		"Success",
		todo.TodoAfterDelete{},
		"Success",
		false,
	)

	c.JSON(http.StatusOK, response)
}

func (h *todoHandler) UpdateTodoByID(c *gin.Context) {
	// get id from param
	idParam := c.Param("id")

	// convert to int
	id, err := strconv.Atoi(idParam)
	if err != nil {
		response := utils.ResponseAPI(
			http.StatusText(http.StatusBadRequest),
			nil,
			"id must be int",
			true,
		)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	// bind
	var input todo.TodoUpdateInput

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
	todoUpdated, err := h.serviceTodo.UpdateByID(input, id)
	if err != nil {
		response := utils.ResponseAPI(
			STATUS_NOT_FOUND,
			nil,
			err.Error(),
			true,
		)

		c.JSON(http.StatusNotFound, response)
		return
	}

	todoFormatted := todo.FormatterTodo(todoUpdated)
	response := utils.ResponseAPI(
		"Success",
		todoFormatted,
		"Success",
		false,
	)

	c.JSON(http.StatusOK, response)
}
