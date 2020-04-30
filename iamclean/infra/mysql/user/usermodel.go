package mysql

import (
	"github.com/jinzhu/gorm"
	entity "github.com/yuk1ty/easy-learn-clean-architecture/iamclean/model/entity/user"
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

func (u GormUser) ToUser() (*entity.User, error) {
	return entity.NewUser(u.Name, u.Roles) // TODO roles がめんどくさくて変えてないけど。
}