package models

type Tag struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Color     string `json:"color"`
	Bg        string `json:"bg"`
	ProjectId string `json:"projectId"`
}
