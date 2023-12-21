package domains

import (
	"study-kafka-ddb/domains/enums"
	"time"
)

type ApplicationForm struct {
	ID            int64
	CreatedAt     time.Time
	UpdatedAt     time.Time
	IsOnGoing     bool
	UserID        enums.UserID
	FavoriteGenre enums.Genre
	DesiredRole   enums.BandRole
	Contents      string
}

type RecruitmentNotice struct {
	ID           int64
	CreatedAt    time.Time
	UpdatedAt    time.Time
	IsOnGoing    bool
	BandID       enums.BandID
	OpenPosition enums.BandRole
	Contents     string
}
