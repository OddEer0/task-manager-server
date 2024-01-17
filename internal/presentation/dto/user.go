package dto

import "time"

type (
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
