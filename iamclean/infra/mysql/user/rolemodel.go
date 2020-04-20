package mysql

import (
	"github.com/jinzhu/gorm"
	"time"
)

type GormRole struct {
	gorm.Model
	Id string `gorm:"primary_key"`
	Name string
	UserId string
	CreatedAt time.Time `gorm:"type:datetime(6)"`
	UpdatedAt time.Time `gorm:"type:datetime(6)"`
}

func (r GormRole) TableName() string {
	return "role"
}