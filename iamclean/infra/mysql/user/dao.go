package mysql

import "github.com/jinzhu/gorm"

type UserDao struct {
	db *gorm.DB
}

func NewUserDao(db *gorm.DB) UserDao {
	return UserDao{db: db}
}
