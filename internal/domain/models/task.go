package models

type Task struct {
	Id          string  `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
	Priority    *string `json:"priority"`
	Order       int     `json:"order"`
	ColumnId    string  `json:"columnId"`
}
