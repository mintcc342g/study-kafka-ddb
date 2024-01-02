package domains

import (
	"study-kafka-ddb/domains/enums"
	"time"
)

type BandMember struct {
	BandID    enums.BandID
	UserID    enums.UserID
	Role      enums.BandPosition
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
