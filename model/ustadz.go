package model

import "time"

type Ustadz struct {
	ID           uint   `gorm:"autoIncrement;not null;primaryKey"`
	Name         string `gorm:"size:255;not null"`
	StudyRundown []StudyRundown
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
