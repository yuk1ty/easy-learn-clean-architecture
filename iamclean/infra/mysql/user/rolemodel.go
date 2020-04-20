package mysql

import (
	"github.com/jinzhu/gorm"
	"time"
)

type GormRole struct {
	gorm.Model
	Id string
	Name string
	UserId string
	CreatedAt time.Time `gorm:"type:datetime(6)"`
	UpdatedAt time.Time `gorm:"type:datetime(6)"`
}
