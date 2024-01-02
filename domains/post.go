package domains

import (
	"study-kafka-ddb/domains/enums"
	"time"
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
	post.IsOpened = true
	post.Position = position
	post.BandID = bandID

	return post
}

func NewResumePost(userID enums.UserID, contents string, position enums.BandPosition, genre enums.Genre) *Post {
	post := NewPost(userID, contents)
	post.IsOpened = true
	post.Position = position
	post.FavoriteGenre = genre

	return post
}
