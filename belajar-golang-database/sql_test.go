package belajar_golang_database

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"
)

func TestExecSQL(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "INSERT INTO customer(id, name) VALUES('0001', 'zulhaditya')"
	_, err := db.ExecContext(ctx, script)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert new customer")
}

func TestQuerySQL(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "SELECT id, name FROM customer"
	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var id, name string
		err := rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}

		fmt.Println("ID:", id)
		fmt.Println("Name:", name)
	}
}

func TestQuerySQLComplex(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "SELECT id, name, email, balance, rating, birth_date, married, created_at FROM customer"
	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var id, name string
		var email sql.NullString
		var balance int32
		var rating float64
		var birthDate sql.NullTime
		var createdAt time.Time
		var married bool

		err := rows.Scan(&id, &name, &email, &balance, &rating, &birthDate, &married, &createdAt)
		if err != nil {
			panic(err)
		}

		fmt.Println("================ RESULT =====================")
		fmt.Println("ID:", id)
		fmt.Println("Name:", name)

		if email.Valid {
			fmt.Println("Email:", email.String)
		}

		fmt.Println("Balance:", balance)
		fmt.Println("Rating:", rating)

		if birthDate.Valid {
			fmt.Println("Birth Date:", birthDate.Time)
		}

		fmt.Println("Maried:", married)
		fmt.Println("Created At:", createdAt)
	}
}

func TestSQLInjection(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "admin'; #" // script sql injection
	password := "salah"

	script := "SELECT username FROM user WHERE username = '" + username +
		"' AND password = '" + password + "' LIMIT 1"
	fmt.Println(script)
	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	if rows.Next() {
		var username string
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("Success login", username)
	} else {
		fmt.Println("Failed to login")
	}
}

func TestSQLInjectionSafe(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "admin'; #" // script sql injection
	password := "salah"

	script := "SELECT username FROM user WHERE username = ? AND password = ? LIMIT 1"
	fmt.Println(script)
	rows, err := db.QueryContext(ctx, script, username, password)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	if rows.Next() {
		var username string
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("Success login", username)
	} else {
		fmt.Println("Failed to login")
	}
}

func TestExecSQLParameter(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "ackxle'; DROP TABLE user; #"
	password := "ackxle"

	script := "INSERT INTO user(username, password) VALUES(?,?)"
	_, err := db.ExecContext(ctx, script, username, password)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert new user")
}

func TestAutoIncrement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	email := "ackxle@gmail.com"
	comment := "test comment"

	script := "INSERT INTO comments(email, comment) VALUES(?,?)"
	result, err := db.ExecContext(ctx, script, email, comment)
	if err != nil {
		panic(err)
	}

	insertId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert new comment with id", insertId)
}
