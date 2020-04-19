package registry

import (
	"github.com/jinzhu/gorm"
	"github.com/yuk1ty/easy-learn-clean-architecture/iamclean/adapter/infra/mysql"
	userdao "github.com/yuk1ty/easy-learn-clean-architecture/iamclean/adapter/infra/mysql/user"
	"log"
)

type MySqlDaoRegistry struct {
	db *gorm.DB
}

func NewMySqlDaoRegistry() MySqlDaoRegistry {
	db, err := mysql.MySqlContext()
	if err != nil {
		log.Fatal(err.Error())
	}
	return MySqlDaoRegistry{db: db}
}

func (r MySqlDaoRegistry) UserDao() userdao.UserDao {
	return userdao.NewUserDao(r.db)
}
