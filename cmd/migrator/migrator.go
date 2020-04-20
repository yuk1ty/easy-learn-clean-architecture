package main

import (
	"github.com/yuk1ty/easy-learn-clean-architecture/iamclean/infra/mysql"
	mysqluser "github.com/yuk1ty/easy-learn-clean-architecture/iamclean/infra/mysql/user"
	"log"
)

func main() {
	// migration
	db, err := mysql.MySqlContext()
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}

	db.CreateTable(&mysqluser.GormUser{})
	db.CreateTable(&mysqluser.GormRole{})
}