package dtos

import (
	"study-kafka-ddb/domains/enums"
	"time"
)

type SignUpReq struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignUpResp struct {
	ID        enums.UserID `json:"id"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
	Name      string       `json:"name"`
	Email     string       `json:"email"`
}

type SignInReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignInResp struct {
	Success bool `json:"success"`
}

type OpenPositionReq struct {
	UserID   enums.UserID       `json:"user_id"`
	BandID   enums.BandID       `json:"band_id"`
	Position enums.BandPosition `json:"position"`
	Contents string             `json:"contents"`
}

type OpenPositionResp struct {
	ID        int64              `json:"id"`
	BandID    enums.BandID       `json:"band_id"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at"`
	Position  enums.BandPosition `json:"position"`
	Contents  string             `json:"contents"`
	IsOpened  bool               `json:"is_opened"`
}

type SeekPositionReq struct {
	UserID        enums.UserID       `json:"user_id"`
	BandID        enums.BandID       `json:"band_id"`
	Position      enums.BandPosition `json:"position"`
	Contents      string             `json:"contents"`
	FavoriteGenre enums.Genre        `json:"favorite_genre"`
}

type SeekPositionResp struct {
	ID            int64              `json:"id"`
	CreatedAt     time.Time          `json:"created_at"`
	UpdatedAt     time.Time          `json:"updated_at"`
	Position      enums.BandPosition `json:"position"`
	Contents      string             `json:"contents"`
	IsOpened      bool               `json:"is_opened"`
	FavoriteGenre enums.Genre        `json:"favorite_genre"`
}
