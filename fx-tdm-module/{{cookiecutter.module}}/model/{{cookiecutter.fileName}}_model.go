package model

import "time"

type {{cookiecutter.className}} struct {
	Id        string     `json:"id" bson:"id"`
	AppId     string     `json:"appId" bson:"appId"`
	CreatedAt *time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt" bson:"updatedAt"`
}

type Paginate{{cookiecutter.className}} struct {
	Page    int64      `json:"page"`
	PerPage int64      `json:"perPage"`
	Total   int64      `json:"total"`
	Data    []*{{cookiecutter.className}} `json:"data"`
}
