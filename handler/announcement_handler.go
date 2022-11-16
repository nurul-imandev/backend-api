package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"nurul-iman-blok-m/announcement"
	"nurul-iman-blok-m/helper"
	"nurul-iman-blok-m/model"
	"strings"
	"time"
)

type announcementHandler struct {
	service announcement.AnnouncementService
}

func NewHandlerAnnouncement(service announcement.AnnouncementService) *announcementHandler {
	return &announcementHandler{service}
}

func (h *announcementHandler) AddAnnouncement(c *gin.Context) {
	var input announcement.AnnouncementInput
	err := c.ShouldBind(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errMessage := gin.H{"errors": errors}

		response := helper.ApiResponse("You must completed field", http.StatusUnprocessableEntity, "error", errMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	currentUser := c.MustGet("currentUser").(model.User)
	input.UserID = currentUser.ID
	input.User = currentUser

	fileImage, errBanner := c.FormFile("banner")
	if errBanner != nil {
		fmt.Println("banner-error")
		response := helper.ApiResponse("Failed to upload banner image", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	extenstionFile := ""
	fileName := strings.Split(fileImage.Filename, ".")

	if len(fileName) == 2 {
		extenstionFile = fileName[1]
	}
	path := fmt.Sprintf("images/announcement-%s-%s.%s", input.Slug, time.Now().Format("2006-02-01"), extenstionFile)
	errUploadBanner := c.SaveUploadedFile(fileImage, path)

	if errUploadBanner != nil {
		fmt.Println("path-error")
		response := helper.ApiResponse("Failed to upload banner image", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	responseAddAnnouncement, createdBy, errAdd := h.service.AddAnnouncement(input, path)
	if errAdd != nil {
		response := helper.ApiResponse("Failed to add announcement or", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := announcement.AnnouncementFormat(responseAddAnnouncement, createdBy)

	response := helper.ApiResponse("Success to add announcement", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}
