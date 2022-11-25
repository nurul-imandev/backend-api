package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"nurul-iman-blok-m/helper"
	"nurul-iman-blok-m/model"
	"nurul-iman-blok-m/study_rundown"
)

type StudyRundownHandler struct {
	service study_rundown.StudyService
}

func NewHandlerStudyRundown(service study_rundown.StudyService) *StudyRundownHandler {
	return &StudyRundownHandler{service}
}

func (h *StudyRundownHandler) AddStudy(c *gin.Context) {
	var input study_rundown.StudyRundownInput

	err := c.ShouldBind(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errMessage := gin.H{"errors": errors}

		response := helper.ApiResponse("you must complete field", http.StatusBadRequest, "error", errMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	currentUser := c.MustGet("currentUser").(model.User)

	if currentUser.Role.RoleName == "user" {
		response := helper.ApiResponse("You not have access for add", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	study, errAdd := h.service.AddStudy(input)
	if errAdd != nil {
		response := helper.ApiResponse("Failed to add rundown", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := study_rundown.StudyResponseFormat(study)

	response := helper.ApiResponse("Success to add rundown", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}

func (h *StudyRundownHandler) GetListUstadzName(c *gin.Context) {
	name, err := h.service.GetListUstadName()
	if err != nil {
		response := helper.ApiResponse("Error to get ustadz name", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.ApiResponse("List ustadz", http.StatusOK, "success", study_rundown.ListUstadzJsonFormatter(name))
	c.JSON(http.StatusOK, response)
}
