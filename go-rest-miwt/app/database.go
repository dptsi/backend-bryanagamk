package app

import (
	"bryanagamk/go-rest-miwt/helper"
	"database/sql"
	"fmt"
)

func NewDB(driver interface{}, source interface{}, host interface{}, user interface{}, pass interface{}, name interface{}, port interface{}) *sql.DB {
	dataSource := fmt.Sprintf("%v://%v:%v@%v:%v?database=%v", source, user, pass, host, port, name)
	db, err := sql.Open(fmt.Sprint(driver), dataSource)
	helper.PanicIfError(err)

	return db
}
