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
