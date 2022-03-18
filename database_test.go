package gomysql

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func TestExecContext(t *testing.T) {
	db := GetConnection()
	defer db.Close()
	
	// gunakan db

	ctx := context.Background()

	query := "INSERT INTO customer(id, name) VALUES('p1', 'fahmi')"
	
	var _, errExecContext = db.ExecContext(ctx, query)
	if errExecContext != nil {
		fmt.Println(errExecContext)
		panic("error exec query")
	}
	fmt.Println("success")
}

func TestQueryContext (t *testing.T){
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	query := "SELECT id, name FROM customer"

	rows, err := db.QueryContext(ctx, query)
	defer rows.Close()

	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var i, n string
		err := rows.Scan(&i, &n)
		if err != nil {
			panic(err)  
		}

		fmt.Println("Id", i,)
		fmt.Println("Name", n)
	}

}

func TestComplexQueryContext (t *testing.T){
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	query := "SELECT id, name, email, balance, rating, birth_date, married, created_at FROM customer"

	rows, err := db.QueryContext(ctx, query)
	defer rows.Close()

	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var id, name string
		var email sql.NullString
		var balance int
		var rating float64
		var birthDate sql.NullTime 
		var createdAt time.Time
		var married bool

		err := rows.Scan(
			&id, 
			&name, 
			&email, 
			&balance, 
			&rating, 
			&birthDate, 
			&married, 
			&createdAt,
		)
		if err != nil {
			panic(err)  
		}

		fmt.Println("=================")
		fmt.Println("Id:", id)
		fmt.Println("Name:", name)
		if email.Valid {
			fmt.Println("Email:", email.String)
		}
		fmt.Println("Balance:", balance)
		fmt.Println("Rating:", rating)
		if birthDate.Valid {
			fmt.Println("Birth:", birthDate.Time)
		}
		fmt.Println("Married:", married)
		fmt.Println("Create:", createdAt)
		fmt.Println("=================")
	}

}

func TestSQLInjection (t *testing.T){


	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "admin'; -- "
	password := "salah"

	query := "SELECT username FROM user WHERE username = '" +
	 username + 
	 "' AND password = '" + 
	 password + 
	 "' LIMIT 1"

	rows, err := db.QueryContext(ctx, query)
	defer rows.Close()

	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var username string

		err := rows.Scan(
			&username,	
		)
		if err != nil {
			panic(err)  
		}

		fmt.Println("Username:", username)
	}
}

func TestSQLSafeInjection (t *testing.T){


	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "admin"
	password := "admin"

	query := "SELECT username FROM user WHERE username = ? AND password = ? LIMIT 1"

	rows, err := db.QueryContext(ctx, query, username, password)
	defer rows.Close()

	if err != nil {
		panic(err)
	}

	if rows.Next() {
		var username string

		err := rows.Scan(
			&username,	
		)
		if err != nil {
			panic(err)  
		}

		fmt.Println("Login")
		fmt.Println("Username:", username)
	} else {
		fmt.Println("Gagal Login")
	}
}

func TestExecSafeInjection(t *testing.T) {
	db := GetConnection()
	defer db.Close()
	
	// gunakan db

	ctx := context.Background()

	username := "fahmi"
	password:="1234"

	query := "INSERT INTO user(username, password) VALUES(?, ?)"
	
	var _, errExecContext = db.ExecContext(ctx, query, username, password)
	if errExecContext != nil {
		fmt.Println(errExecContext)
		panic("error exec query")
	}
	fmt.Println("success")
}

func TestInsertAutoIncrement(t *testing.T) {
	db := GetConnection()
	defer db.Close()
	
	// gunakan db

	ctx := context.Background()

	email := "eg@gmail.com"
	comment:="1234"

	query := "INSERT INTO comments(email, comment) VALUES(?, ?)"
	
	var result, errExecContext = db.ExecContext(ctx, query,email, comment)
	if errExecContext != nil {
		fmt.Println(errExecContext)
		panic("error exec query")
	}

	var id, err = result.LastInsertId()

	if err != nil {
		panic(err)
	}

	fmt.Println(id)
}

func TestPrepareStatement (t *testing.T){
	db:= GetConnection()
	defer db.Close()

	ctx:= context.Background()
	script := "INSERT INTO comments(email, comment) VALUES (? , ?)"
	statement, err := db.PrepareContext(ctx, script)
	defer statement.Close()

	if err != nil {
		panic(err)
	}

	for i := 0; i < 10; i++ {
		email := "fahmi" + strconv.Itoa(i) + "@gmail.com"
		comment := "comment ke " + strconv.Itoa(i)

		result, err := statement.ExecContext(ctx, email, comment)
		if err != nil {
			panic(err)
		}

		id, err := result.LastInsertId()

		fmt.Println("Comment id:", id)

	}

}

func TestTransaction (t *testing.T){
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	tx, err :=  db.Begin()

	if err != nil {
		panic(err)
	}

	// do transaction
	script := "INSERT INTO comments(email, comment) VALUES (? , ?)"
	stmt, err := tx.PrepareContext(ctx, script)
	defer stmt.Close()

	for i := 100000; i < 110000; i++ {
		email := "fahmi" + strconv.Itoa(i) + "@gmail.com"
		comment := "comment ke " + strconv.Itoa(i)
		result, err := stmt.ExecContext(ctx, email, comment)
		
		if err != nil {
			panic(err)
		}

		id, err := result.LastInsertId()

		fmt.Println("Comment id:", id)

	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		panic(err)
	}

}

func BenchmarkSelect(b *testing.B) {
	for i:= 0; i < b.N; i++{
		db := GetConnection()
		defer db.Close()

		ctx := context.Background()

		script := "SELECT id, email, comment FROM comments WHERE email = ?"
		stmt, err := db.PrepareContext(ctx, script)
		defer stmt.Close()

		if err != nil {
			panic(err)
		}
		
		email := "fahmi105999@gmail.com"

		rows, err := stmt.QueryContext(ctx, email)

		var id int
		var comment string

		if err != nil {
			panic(err)
		}

		if rows.Next() {
			err := rows.Scan(&id, &email, &comment)

			if err != nil {
				panic(err)
			}

		} else {
			fmt.Println("gagal")
		}

		fmt.Println("Id:", id)
		fmt.Println("Email:", email)
		fmt.Println("Comment:", comment)
		
	}
}