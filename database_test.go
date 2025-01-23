package learn_golang_database_mysql

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"testing"
	"time"
)

func Test(t *testing.T) {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/learn_golang_database")
	if err != nil {
		t.Fatal(err)
	}
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			t.Fatal(err)
		}
	}(db)
}
