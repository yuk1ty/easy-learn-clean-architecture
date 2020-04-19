package mysql

import (
	"github.com/jinzhu/gorm"
	"time"
)

type User struct {
	gorm.Model
	Id string `gorm:"primary_key"`
	Name string
	// TODO Roles?
	CreatedAt time.Time `gorm:"type:datetime(6)"`
	UpdatedAt time.Time `gorm:"type:datetime(6)"`
}
