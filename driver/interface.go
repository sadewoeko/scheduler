package driver

import "database/sql"

type DBHandler interface {
	Execute(statement string) (sql.Result, error)
	Query(statement string) (Row, error)
}

type Row interface {
	Scan(dest ...interface{}) error
	Next() bool
}
