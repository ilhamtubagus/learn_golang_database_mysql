package learn_golang_database_mysql

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"
)

func TestExecSql(t *testing.T) {
	db, err := GetConnection()
	if err != nil {
		t.Fatal(err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			t.Fatal(err)
		}
	}(db)

	ctx := context.Background()

	script := "INSERT INTO customer (id, name) VALUES ('1', 'ilham')"
	_, err = db.ExecContext(ctx, script) // ExecContext can be used for insert, update, delete
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("Data inserted successfully")
}

func TestQuerySql(t *testing.T) {
	db, err := GetConnection()
	if err != nil {
		t.Fatal(err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			t.Fatal(err)
		}
	}(db)

	ctx := context.Background()

	script := "SELECT id, name FROM customer WHERE id = '1'"
	rows, err := db.QueryContext(ctx, script) // QueryContext can be used for query
	if err != nil {
		t.Fatal(err)
	}

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			t.Fatal(err)
		}
	}(rows)

	for rows.Next() {
		var id, name string
		err := rows.Scan(&id, &name)
		if err != nil {
			t.Fatal(err)
		}

		fmt.Println("id", id)
		fmt.Println("name", name)
	}

}

func TestQuerySqlComplex(t *testing.T) {
	db, err := GetConnection()
	if err != nil {
		t.Fatal(err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			t.Fatal(err)
		}
	}(db)

	ctx := context.Background()

	script := "SELECT id, name, email, balance, rating, birth_date, married, created_at FROM customer"
	rows, err := db.QueryContext(ctx, script) // QueryContext can be used for query
	if err != nil {
		t.Fatal(err)
	}

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			t.Fatal(err)
		}
	}(rows)

	for rows.Next() {
		var id, name string
		var email sql.NullString // to prevent error when value is NULL
		var balance int32
		var rating float64
		var createdAt time.Time
		var birthDate sql.NullTime // to prevent error when value is NULL
		var married bool
		err := rows.Scan(&id, &name, &email, &balance, &rating, &birthDate, &married, &createdAt)
		if err != nil {
			t.Fatal(err)
		}

		fmt.Println("--- Customer ---")
		fmt.Println("id", id)
		fmt.Println("name", name)
		if email.Valid {
			fmt.Println("email", email.String)
		}
		fmt.Println("balance", balance)
		fmt.Println("rating", rating)
		birthDateTime, _ := birthDate.Value()
		fmt.Println("birth_date", birthDateTime)
		fmt.Println("married", married)
		fmt.Println("created_at", createdAt)
	}

}

func TestSqlInjection(t *testing.T) {
	db, _ := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "admin'; #"
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
		fmt.Println("Sukses Login", username)
	} else {
		fmt.Println("Gagal Login")
	}
}

func TestSqlInjectionSafe(t *testing.T) {
	db, _ := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "admin"
	password := "admin"

	script := "SELECT username FROM user WHERE username = ? AND password = ? LIMIT 1"
	fmt.Println(script)
	rows, err := db.QueryContext(ctx, script, username, password) // variadic arguments username and password
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
		fmt.Println("Sukses Login", username)
	} else {
		fmt.Println("Gagal Login")
	}
}

func TestExecSqlParameter(t *testing.T) {
	db, err := GetConnection()
	if err != nil {
		t.Fatal(err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			t.Fatal(err)
		}
	}(db)

	ctx := context.Background()

	username := "fian'; DROP TABLE user; #"
	password := "fian"

	script := "INSERT INTO user (username, password) VALUES (?, ?)" // sql parameter to prevent sql injection
	_, err = db.ExecContext(ctx, script, username, password)        // ExecContext can be used for insert, update, delete
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("Data inserted successfully")
}

func TestAutoIncrement(t *testing.T) {
	db, err := GetConnection()
	if err != nil {
		t.Fatal(err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			t.Fatal(err)
		}
	}(db)

	ctx := context.Background()

	script := "INSERT INTO comments (email, comment) VALUES ('fian@gmail.com', 'bismillah')"
	result, err := db.ExecContext(ctx, script) // ExecContext can be used for insert, update, delete
	if err != nil {
		t.Fatal(err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("Inserted comment id:", id)

	fmt.Println("Data inserted successfully")
}
