package domains

import (
	"encoding/json"
	"study-kafka-ddb/domains/enums"
	"study-kafka-ddb/utils/deftype"
	"time"

	"go.uber.org/zap"
)

const (
	validityPeriodOfPost = 6 * 30 * 24 * time.Hour // 6 months
)

type Post struct {
	ID        int64
	WriterID  enums.UserID
	CreatedAt time.Time
	UpdatedAt time.Time
	Contents  string
	Type      enums.PostType

	// NOTE: wanted & resume
	IsOpened bool
	Position enums.BandPosition

	// NOTE: wanted
	BandID enums.BandID

	// NOTE: resume
	FavoriteGenre enums.Genre
}

func isValidContents(contents string) bool {
	return len(contents) <= 800
}

func NewPost(userID enums.UserID, contents string) *Post {
	return &Post{
		WriterID:  userID,
		CreatedAt: time.Now(),
		Contents:  contents,
	}
}

func NewWantedPost(userID enums.UserID, contents string, position enums.BandPosition, bandID enums.BandID) *Post {
	post := NewPost(userID, contents)
	post.Type = enums.PostTypeWanted
	post.IsOpened = true
	post.Position = position
	post.BandID = bandID

	return post
}

func NewResumePost(userID enums.UserID, contents string, position enums.BandPosition, genre enums.Genre) *Post {
	post := NewPost(userID, contents)
	post.Type = enums.PostTypeResume
	post.IsOpened = true
	post.Position = position
	post.FavoriteGenre = genre

	return post
}

func (r *Post) IsWanted() bool {
	return r.Type == enums.PostTypeWanted
}

func (r *Post) IsResume() bool {
	return r.Type == enums.PostTypeResume
}

func (r *Post) IsExpired() bool {
	expired := r.CreatedAt.Add(validityPeriodOfPost)
	return time.Now().After(expired)
}

func (r *Post) MakeMessage() ([]byte, deftype.Error) {
	if r.IsWanted() || r.IsResume() {
		m := map[string]interface{}{
			"id": r.ID,
		}
		bytm, err := json.Marshal(m)
		if err != nil {
			zap.S().Error("fail to marshal a post", "err", err, "post_id", r.ID)
			return nil, deftype.ErrInternalServerError
		}
		return bytm, nil
	}

	zap.S().Error("invalid post type", "post_id", r.ID)
	return nil, deftype.ErrInvalidRequestData
}
