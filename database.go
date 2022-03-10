package gomysql

import (
	"database/sql"
	"time"
)

func GetConnection() *sql.DB {

	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/gosample")
	if err != nil {
		panic(err)
	}
	
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(50)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(time.Hour)

	return db

}