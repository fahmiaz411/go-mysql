package gomysql

import (
	"database/sql"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestOpenConnection(t *testing.T) {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/gosample")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	
	// gunakan db

}	