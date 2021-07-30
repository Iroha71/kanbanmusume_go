package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	db  *gorm.DB
	err error
)

func Init() {
	db, err = gorm.Open("postgres", "host=localhost port5432 user=postgres password=postgres dbname=kanbanmusume")
	if err != nil {
		panic(err)
	}
}
