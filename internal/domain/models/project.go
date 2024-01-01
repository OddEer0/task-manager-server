package models

type Project struct {
	Id          string  `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
	Bg          string  `json:"bg"`
	Color       string  `json:"color"`
	Order       int     `json:"order"`
	ProjectId   string  `json:"projectId"`
}
