package domains

import (
	"study-kafka-ddb/domains/enums"
	"time"
)

type User struct {
	ID        enums.UserID
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
	Email     string
}
