package domains

import (
	"study-kafka-ddb/domains/enums"
	"study-kafka-ddb/utils/deftype"
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

func (r *Band) OpenPosition(user *User, position enums.BandPosition, contents string) (*Post, deftype.Error) {
	if !r.CheckReader(user) {
		return nil, deftype.ErrUnauthorized
	}

	if !isValidContents(contents) {
		return nil, deftype.ErrInvalidRequestData
	}

	return NewWantedPost(user.ID, contents, position, r.ID), nil
}

func (r *Band) CheckReader(user *User) bool {
	return r.ReaderID == user.ID
}
