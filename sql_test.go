package learn_golang_database_mysql

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
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
