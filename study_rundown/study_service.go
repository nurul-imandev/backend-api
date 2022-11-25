package study_rundown

import "nurul-iman-blok-m/model"

type StudyService interface {
	AddStudy(input StudyRundownInput) (model.StudyRundown, error)
	GetListUstadName() ([]model.User, error)
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
