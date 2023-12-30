package app

import (
	"database/sql"
	"time"

	"github.com/midoon/simple-crud-go/helper"
)

func NewDB() *sql.DB {
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/simple_crud_go_migration")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}

// cara untuk membuat db migration
//   migrate create -ext sql -dir [nama direktori dari root project] [nama_file_migration]
// conoth: migrate create -ext sql -dir db/migrations create_table_category

// Menjalankan DB migration
// migrate -database ["koneksi database"] -path [folder filel migration] up
// migrate -database "mysql://root@tcp(localhost:3306)/simple_crud_go_migration" -path db/migrations up
