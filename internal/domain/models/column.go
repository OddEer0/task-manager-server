package models

type Column struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Bg        string `json:"bg"`
	Color     string `json:"color"`
	Order     int    `json:"order"`
	ProjectId string `json:"projectId"`
}
