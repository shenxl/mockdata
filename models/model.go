package models

import "time"

type Model struct {
	Id        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"createdtime"`
	UpdatedAt time.Time  `json:"updatetime"`
	DeletedAt *time.Time `json:"deletetime"`
}
