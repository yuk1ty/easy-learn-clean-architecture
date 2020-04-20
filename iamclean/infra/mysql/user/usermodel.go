package mysql

import (
	"github.com/jinzhu/gorm"
	"time"
)

type GormUser struct {
	gorm.Model
	Id string `gorm:"primary_key"`
	Name string
	Roles []GormRole    `gorm:"foreignkey:UserId"`
	CreatedAt time.Time `gorm:"type:datetime(6)"`
	UpdatedAt time.Time `gorm:"type:datetime(6)"`
}

func (u GormUser) TableName() string {
	return "user"
}