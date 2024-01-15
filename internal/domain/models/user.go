package models

import (
	"github.com/OddEer0/task-manager-server/internal/domain/valuesobject"
	"time"
)

type User struct {
	Id             string                 `json:"id" validate:"required,uuidv4"`
	Nick           string                 `json:"nick" validate:"required,min=3,max=50"`
	Password       valuesobject.Password  `json:"password" validate:"required"`
	Email          string                 `json:"email" validate:"required,email"`
	FirstName      string                 `json:"firstName" validate:"required,min=3,max=50"`
	LastName       string                 `json:"lastName" validate:"required,min=3,max=50"`
	SubTitle       *string                `json:"subTitle" validate:"omitempty,min=3,max=500"`
	Avatar         *string                `json:"avatar" validate:"omitempty,isLink"`
	Birthday       *valuesobject.Birthday `json:"birthday"`
	Role           string                 `json:"role" validate:"required,userRole"`
	ActivationLink string                 `json:"activationLink" validate:"required,isLink"`
	IsActivate     bool                   `json:"isActivate"`
	IsBanned       bool                   `json:"isBanned"`
	BanReason      *string                `json:"banReason" validate:"omitempty,min=3,max=255"`
	CreatedAt      time.Time              `json:"createdAt" validate:"required,dateIsLessNow"`
	UpdatedAt      time.Time              `json:"updatedAt" validate:"required,dateIsLessNow"`
}
