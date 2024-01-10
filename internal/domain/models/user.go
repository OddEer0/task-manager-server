package models

import (
	"github.com/OddEer0/task-manager-server/internal/domain/valuesobject"
	"time"
)

type User struct {
	Id             string                 `json:"id"`
	Nick           string                 `json:"nick"`
	Password       string                 `json:"password"`
	Email          valuesobject.Email     `json:"email"`
	FirstName      string                 `json:"firstName"`
	LastName       string                 `json:"lastName"`
	SubTitle       *string                `json:"subTitle"`
	Avatar         *string                `json:"avatar"`
	Birthday       *valuesobject.Birthday `json:"birthday"`
	Role           string                 `json:"role"`
	ActivationLink string                 `json:"activationLink"`
	IsActivate     bool                   `json:"isActivate"`
	IsBanned       bool                   `json:"isBanned"`
	BanReason      *string                `json:"banReason"`
	CreatedAt      time.Time              `json:"createdAt"`
	UpdatedAt      time.Time              `json:"updatedAt"`
}
