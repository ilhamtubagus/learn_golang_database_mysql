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
	_, err = db.ExecContext(ctx, script)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("Data inserted successfully")
}
