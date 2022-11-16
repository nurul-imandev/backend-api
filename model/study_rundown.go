package model

import "time"

type StudyRundown struct {
	ID          uint
	Title       string
	OnScheduled int
	Ustadz      Ustadz
	UstadzID    uint
	StartTime   time.Time
	EndTime     time.Time
}
