package mysql

import (
	"github.com/jinzhu/gorm"
	"github.com/yuk1ty/easy-learn-clean-architecture/iamclean/context/apperr"
)
import _ "github.com/jinzhu/gorm/dialects/mysql"

func MySqlContext() (*gorm.DB, error) {
	db, err := gorm.Open("mysql", "user:password@/db_name?charset=utf8&parseTime=True&loc=Local") // FIXME
	defer db.Close()
	if err != nil {
		return nil, apperr.NewMySqlConnectionErr("MySQL への接続に失敗しました: " + err.Error())
	}

	return db, nil
}
