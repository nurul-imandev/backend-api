package study_rundown

import (
	"fmt"
	"gorm.io/gorm"
	"nurul-iman-blok-m/model"
)

type StudyService interface {
	AddStudy(input StudyRundownInput) (model.StudyRundown, error)
	GetListUstadName() ([]model.User, error)
	GetListStudy(list func(db *gorm.DB) *gorm.DB) ([]model.StudyRundown, int, error)
}

type StudyServiceImpl struct {
	repository StudyRepository
}

func NewService(repository StudyRepository) *StudyServiceImpl {
	return &StudyServiceImpl{repository}
}

func (s *StudyServiceImpl) AddStudy(input StudyRundownInput) (model.StudyRundown, error) {
	study := model.StudyRundown{}
	study.Title = input.Title
	study.UserID = input.UserID
	study.Time = input.Time
	study.ScheduleDate = input.ScheduleDate
	if input.OnScheduled {
		study.OnScheduled = 1
	} else {
		study.OnScheduled = 0
	}

	addStudy, err := s.repository.AddStudy(study)
	if err != nil {
		return model.StudyRundown{}, err
	}
	return addStudy, nil
}

func (s *StudyServiceImpl) GetListUstadName() ([]model.User, error) {
	ustadName, err := s.repository.GetListUstadName()
	if err != nil {
		return ustadName, err
	}

	return ustadName, nil
}

func (s *StudyServiceImpl) GetListStudy(list func(db *gorm.DB) *gorm.DB) ([]model.StudyRundown, int, error) {
	rundowns, count, err := s.repository.GetListStudies(list)
	for _, studyRundown := range rundowns {
		fmt.Println("namanya adalah", studyRundown.User)
	}
	if err != nil {
		return rundowns, 0, err
	}
	return rundowns, count, err
}
