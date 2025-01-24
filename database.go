package learn_golang_database_mysql

import (
	"database/sql"
	"time"
)

func GetConnection() (*sql.DB, error) {
	// set parseTime to true to parse datetime field automatically into time.Time
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/learn_golang_database?parseTime=true")
	if err != nil {
		return nil, err
	}
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db, nil
}
