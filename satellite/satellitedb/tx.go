// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information.

package satellitedb

import (
	"context"
	"database/sql"

	dbx "storj.io/storj/satellite/satellitedb/dbx"
)

// RawMethods implements dbx.Methods with standard sql interface
type RawMethods interface {
	dbx.Methods
	Rebinder

	Exec(query string, args ...interface{}) (sql.Result, error)
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
}

// DBXTx implements dbx.Methods that matches RawMethods interface
type DBXTx struct {
	*sql.Tx
	dbx.Methods
	Rebinder
}

// Rebinder provides Rebind command
type Rebinder interface {
	Rebind(s string) string 
}