package announcement

import (
	"gorm.io/gorm"
	"nurul-iman-blok-m/model"
)

type AnnouncementRepository interface {
	AddAnnouncement(announcement model.Announcement) (model.Announcement, error)
	GetUserName(announcement model.Announcement, userId uint) (model.Announcement, error)
	GetRoleForException(user model.User) (model.User, error)
}

type announcementRepository struct {
	database *gorm.DB
}

func NewRepositoryAnnouncement(db *gorm.DB) *announcementRepository {
	return &announcementRepository{db}
}

func (r *announcementRepository) AddAnnouncement(announcement model.Announcement) (model.Announcement, error) {
	err := r.database.Create(&announcement).Error

	if err != nil {
		return announcement, err
	}

	return announcement, nil
}

func (r *announcementRepository) GetUserName(announcement model.Announcement, userId uint) (model.Announcement, error) {
	err := r.database.Preload("User").Where("id = ?", userId).Find(&announcement).Error
	if err != nil {
		return announcement, err
	}

	return announcement, nil
}

func (r *announcementRepository) GetRoleForException(user model.User) (model.User, error) {
	userRole := user
	err := r.database.Preload("Role").Find(&userRole).Error
	if err != nil {
		return userRole, err
	}

	return userRole, nil
}
