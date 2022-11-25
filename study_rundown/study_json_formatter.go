package study_rundown

import "nurul-iman-blok-m/model"

type StudyRundownFormatResponse struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	OnScheduled bool   `json:"on_scheduled"`
	Date        string `json:"date"`
	Time        string `json:"time"`
	UstadzName  string `json:"ustadz_name"`
}

type UstadzFormatter struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func StudyResponseFormat(announcement model.StudyRundown) StudyRundownFormatResponse {
	onScheduled := false
	if announcement.OnScheduled == 1 {
		onScheduled = true
	}
	return StudyRundownFormatResponse{
		ID:          announcement.ID,
		Title:       announcement.Title,
		OnScheduled: onScheduled,
		Date:        announcement.ScheduleDate,
		Time:        announcement.Time,
		UstadzName:  announcement.User.Name,
	}
}

func ustadzJsonFormatter(user model.User) UstadzFormatter {
	return UstadzFormatter{
		ID:   user.ID,
		Name: user.Name,
	}
}

func ListUstadzJsonFormatter(users []model.User) []UstadzFormatter {
	listFormatter := []UstadzFormatter{}

	for _, user := range users {
		userFormatter := ustadzJsonFormatter(user)
		listFormatter = append(listFormatter, userFormatter)
	}

	return listFormatter
}
