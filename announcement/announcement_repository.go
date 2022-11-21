package announcement

import (
	"gorm.io/gorm"
	"nurul-iman-blok-m/model"
)

type AnnouncementRepository interface {
	AddAnnouncement(announcement model.Announcement) (model.Announcement, error)
	GetUserName(announcement model.Announcement, userId uint) (model.Announcement, error)
	GetRoleForException(user model.User) (model.User, error)
	GetListAnnouncement(list func(db *gorm.DB) *gorm.DB) ([]model.Announcement, int, error)
	DetailAnnouncement(ID uint) (model.Announcement, error)
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

func (r *announcementRepository) GetListAnnouncement(list func(db *gorm.DB) *gorm.DB) ([]model.Announcement, int, error) {
	var announcements []model.Announcement
	var user model.User
	listAnnouncement := []model.Announcement{}

	err := r.database.Scopes(list).Find(&announcements).Error
	for _, item := range announcements {
		r.database.Where("id = ?", item.UserID).Find(&user)
		itemAnnouncement := model.Announcement{
			ID:          item.ID,
			Title:       item.Title,
			Description: item.Description,
			Images:      item.Images,
			User:        model.User{Name: user.Name},
			UserID:      item.UserID,
			Slug:        item.Slug,
			CreatedAt:   item.CreatedAt,
			UpdatedAt:   item.UpdatedAt,
		}
		listAnnouncement = append(listAnnouncement, itemAnnouncement)
		user = model.User{}
	}

	if err != nil {
		return announcements, 0, err
	}
	totalCount := int64(0)
	r.database.Find(&announcements).Count(&totalCount)
	return listAnnouncement, int(totalCount), nil
}

func (r *announcementRepository) DetailAnnouncement(ID uint) (model.Announcement, error) {
	var announcement model.Announcement
	err := r.database.Preload("User").Where("id = ?", ID).Find(&announcement).Error
	if err != nil {
		return announcement, err
	}
	return announcement, nil
}
