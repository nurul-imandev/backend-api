package announcement

import (
	"errors"
	"nurul-iman-blok-m/model"
)

type AnnouncementService interface {
	AddAnnouncement(input AnnouncementInput, imageLocation string) (model.Announcement, string, error)
}

type announcementService struct {
	repository AnnouncementRepository
}

func NewServiceAnnouncement(repository AnnouncementRepository) *announcementService {
	return &announcementService{repository}
}

func (s *announcementService) AddAnnouncement(input AnnouncementInput, imageLocation string) (model.Announcement, string, error) {
	userRole, errRole := s.repository.GetRoleForException(input.User)
	if errRole != nil {
		return model.Announcement{}, "", errRole
	}

	if userRole.Role.RoleName == "user" {
		return model.Announcement{}, "", errors.New("access denied for create announcement")
	}
	announcement := model.Announcement{}
	announcement.Title = input.Title
	announcement.Description = input.Description
	announcement.Images = imageLocation
	announcement.Slug = input.Slug
	announcement.UserID = input.UserID

	announcementCreate, err := s.repository.AddAnnouncement(announcement)

	if err != nil {
		return announcementCreate, "", err
	}
	user, _ := s.repository.GetUserName(announcement, announcement.UserID)

	return announcementCreate, user.User.Name, nil
}