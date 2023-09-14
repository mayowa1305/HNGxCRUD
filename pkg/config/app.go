package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	db *gorm.DB
)

func Connect() {
	d, err := gorm.Open("postgres", "postgres://hngx_crud_dbdb_user:TVpTB08hDwGn1LNGc3IQpp0XTwBFsdNy@dpg-ck0sh5m3ktkc73dua60g-a.oregon-postgres.render.com/hngx_crud_dbdb")
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDb() *gorm.DB {
	return db
}
