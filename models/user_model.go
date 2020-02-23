package models

import (
	"time"
)

type OrmCustom struct {
	ID        uint `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`
}

type User struct {
	OrmCustom
	Email string `json:"email"`
	Password string `json:"password"`
	Name string `json:"name"`
	Age int `json:"age" sql:"DEFAULT:18"`
	Address string `json:"address"`
}
