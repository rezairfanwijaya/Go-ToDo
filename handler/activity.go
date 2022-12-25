package handler

import (
	"gotodo/activity"
	"gotodo/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

const (
	STATUS_BAD_REQUEST = "Bad Request"
	STATUS_NOT_FOUND   = "Not Found"
)

type activityHandler struct {
	activityService activity.IActivityService
}

// new handler
func NewActivityHandler(activityService activity.IActivityService) *activityHandler {
	return &activityHandler{activityService}
}

func (h *activityHandler) CreateActivity(c *gin.Context) {
	var input activity.ActivityCreateInput

	// bind
	if err := c.BindJSON(&input); err != nil {
		myErr := utils.ErrorBinding(err)
		response := utils.ResponseAPI(
			STATUS_BAD_REQUEST,
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

	// call service
	newActivity, err := h.activityService.CreateActivity(input)
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

	response := utils.ResponseAPI(
		"Success",
		newActivity,
		"Success",
		false,
	)

	c.JSON(http.StatusCreated, response)
}

func (h *activityHandler) GetAllActivity(c *gin.Context) {
	// call service
	activities := h.activityService.GetAllActivity()

	response := utils.ResponseAPI(
		"Success",
		activities,
		"Success",
		false,
	)

	c.JSON(http.StatusOK, response)
}

func (h *activityHandler) GetActivityByID(c *gin.Context) {
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
	activity, err := h.activityService.GetActivityByID(id)
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

	response := utils.ResponseAPI(
		"Success",
		activity,
		"Success",
		false,
	)

	c.JSON(http.StatusOK, response)
}

func (h *activityHandler) DeleteByID(c *gin.Context) {
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
	err = h.activityService.DeleteByID(id)
	if err != nil {
		response := utils.ResponseAPI(
			"Not Found",
			nil,
			err.Error(),
			true,
		)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := utils.ResponseAPI(
		"Success",
		activity.ActiviteAfterDelete{},
		"Success",
		false,
	)

	c.JSON(http.StatusOK, response)
}

func (h *activityHandler) UpdateByID(c *gin.Context) {
	// get id from param
	idParam := c.Param("id")

	// convert to int
	id, err := strconv.Atoi(idParam)
	if err != nil {
		response := utils.ResponseAPI(
			STATUS_BAD_REQUEST,
			nil,
			"id must be int",
			true,
		)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	// get body input
	var input activity.ActivityUpdateInput

	// bind
	if err := c.BindJSON(&input); err != nil {
		myErr := utils.ErrorBinding(err)
		response := utils.ResponseAPI(
			STATUS_BAD_REQUEST,
			nil,
			myErr,
			true,
		)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	// call service
	activityUpdated, err := h.activityService.UpdateByID(input, id)
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

	response := utils.ResponseAPI(
		"Success",
		activityUpdated,
		"Success",
		false,
	)

	c.JSON(http.StatusOK, response)
}
