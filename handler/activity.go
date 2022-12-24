package handler

import (
	"gotodo/activity"
	"gotodo/utils"
	"net/http"

	"github.com/gin-gonic/gin"
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
			http.StatusText(http.StatusBadRequest),
			nil,
			myErr,
			true,
		)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	// call service
	newActivity, err := h.activityService.CreateActivity(input)
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

	response := utils.ResponseAPI(
		"Success",
		newActivity,
		"Success",
		false,
	)

	c.JSON(http.StatusOK, response)
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
