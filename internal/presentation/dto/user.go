package dto

import "time"

type (
	CreateUserDto struct {
		Nick           string `json:"nick" validate:"required,min=3"`
		Password       string `json:"password" validate:"required,min=8"`
		Email          string `json:"email" validate:"required, min=6"`
		FirstName      string `json:"firstName" validate:"required, min=3"`
		LastName       string `json:"lastName" validate:"required, min=3"`
		ActivationLink string `json:"activationLink"`
		Role           string `json:"role"`
	}

	ResponseUserDto struct {
		Id         string     `json:"id"`
		Nick       string     `json:"nick"`
		Email      string     `json:"email"`
		FirstName  string     `json:"firstName"`
		LastName   string     `json:"lastName"`
		SubTitle   *string    `json:"subTitle"`
		Avatar     *string    `json:"avatar"`
		Birthday   *time.Time `json:"birthday"`
		Role       string     `json:"role"`
		IsActivate bool       `json:"isActivate"`
		IsBanned   bool       `json:"isBanned"`
		BanReason  *string    `json:"banReason"`
		CreatedAt  *time.Time `json:"createdAt,omitempty"`
		UpdatedAt  *time.Time `json:"updatedAt,omitempty"`
	}
)
