package study_rundown

import (
	"gorm.io/gorm"
	"nurul-iman-blok-m/model"
)

type StudyRepository interface {
	AddStudy(study model.StudyRundown) (model.StudyRundown, error)
	GetListUstadName() ([]model.User, error)
	//GetListStudy(list func(db *gorm.DB) *gorm.DB)([]model.StudyRundown, error)
	//DetailStudy(ID uint)(model.StudyRundown, error)
	//DeleteStudy(ID uint) error
	//UpdateStudy(study model.StudyRundown)(model.StudyRundown, error)
}

type StudyRepositoryImpl struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *StudyRepositoryImpl {
	return &StudyRepositoryImpl{db}
}

func (s *StudyRepositoryImpl) AddStudy(study model.StudyRundown) (model.StudyRundown, error) {
	var user model.User
	err := s.db.Create(&study).Error

	s.db.Where("id = ?", study.UserID).Find(&user)
	study.User = model.User{Name: user.Name}

	if err != nil {
		return study, err
	}
	return study, nil
}

func (s *StudyRepositoryImpl) GetListUstadName() ([]model.User, error) {
	var users []model.User
	var role model.Role
	ustadzName := []model.User{}

	err := s.db.Find(&users).Error
	if err != nil {
		return users, err
	}
	for _, user := range users {
		s.db.Where("id = ?", user.RoleID).Find(&role)
		itemUstadzname := model.User{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
			Role:  model.Role{RoleName: role.RoleName},
		}
		if itemUstadzname.Role.RoleName == "ustadz" {
			ustadzName = append(ustadzName, itemUstadzname)
		}
		role = model.Role{}
	}

	return ustadzName, nil

}
