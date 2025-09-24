package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var sqlDB *sql.DB

func Connect(dsn string) error {
	var err error
	sqlDB, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	if err = sqlDB.Ping(); err != nil {
		return fmt.Errorf("db ping failed: %w", err)
	}
	return nil
}

func Close() { if sqlDB != nil { _ = sqlDB.Close() } }

func DB() *sql.DB { return sqlDB }
