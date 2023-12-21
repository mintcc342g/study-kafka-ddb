package domains

import (
	"study-kafka-ddb/domains/enums"
	"time"
)

type Band struct {
	ID        enums.BandID
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
	ReaderID  enums.UserID
	MemberNum int
	Genres    []enums.Genre
}
