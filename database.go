package gomysql

import (
	"database/sql"
	"time"
)

func GetConnection() *sql.DB {

	db, err := sql.Open("mysql", "root:1234@tcp(localhost:3306)/belajar_golang?parseTime=true")
	if err != nil {
		panic(err)
	}
	
	db.SetMaxIdleConns(10) // standby connection
	db.SetMaxOpenConns(50)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(time.Hour)

	return db

}