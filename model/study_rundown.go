package model

import "time"

type StudyRundown struct {
	ID           uint   `gorm:"primaryKey;autoIncrement;not null"`
	Title        string `gorm:"size:100;not null"`
	OnScheduled  int    `gorm:"type:tinyInt"`
	ScheduleDate string `gorm:"size:100; not null"`
	Ustadz       Ustadz
	UstadzID     uint
	StartTime    string `gorm:"size:100;not null"`
	EndTime      string `gorm:"size:100;not null"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
