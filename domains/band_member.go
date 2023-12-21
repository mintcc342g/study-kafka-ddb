package domains

import (
	"study-kafka-ddb/domains/enums"
	"time"
)

type BandMember struct {
	BandID    enums.BandID
	UserID    enums.UserID
	Role      enums.BandRole
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
