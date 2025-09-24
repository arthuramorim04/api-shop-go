package testutils

import (
	"database/sql"

	"github.com/DATA-DOG/go-sqlmock"
)

// NewSQLMock returns a sql.DB backed by sqlmock and the mock handle.
func NewSQLMock() (*sql.DB, sqlmock.Sqlmock, error) {
	return sqlmock.New()
}
