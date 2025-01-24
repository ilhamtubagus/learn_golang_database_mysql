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
