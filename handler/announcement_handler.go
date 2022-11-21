package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"nurul-iman-blok-m/announcement"
	"nurul-iman-blok-m/helper"
	"nurul-iman-blok-m/model"
	"strconv"
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

func (h *announcementHandler) GetAllAnnouncement(c *gin.Context) {
	page := c.Request.URL.Query().Get("page")
	perPage := c.Request.URL.Query().Get("per_page")

	paginate := paginateList(page, perPage)

	announcements, count, err := h.service.GetListAnnouncement(paginate)
	if err != nil {
		response := helper.ApiResponse("Error to get announcements", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	pageString, _ := strconv.Atoi(page)
	pageSizeString, _ := strconv.Atoi(perPage)

	response := helper.ApiResponseList("List of roles", http.StatusOK, "success", pageString, pageSizeString, count, announcement.AnnouncementsFormat(announcements))
	c.JSON(http.StatusOK, response)
}

func paginateList(page string, perPage string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		fmt.Println("page :", page)
		fmt.Println("per_page :", perPage)
		pageValue, _ := strconv.Atoi(page)
		if pageValue == 0 {
			pageValue = 1
		}

		perPageValue, _ := strconv.Atoi(perPage)
		switch {
		case perPageValue > 100:
			perPageValue = 100
		case perPageValue <= 0:
			perPageValue = 10
		}

		offset := (pageValue - 1) * perPageValue // (1 - 1) * 5
		return db.Offset(offset).Limit(perPageValue)
	}
}
