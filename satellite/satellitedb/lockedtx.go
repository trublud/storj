// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information.

package satellitedb

import (
	"context"
	"sync"

	"storj.io/storj/satellite/console"
	"storj.io/storj/satellite/orders"
)

// BeginTx is a method for opening transaction
func (m *lockedConsole) BeginTx(ctx context.Context) (console.DBTx, error) {
	m.Lock()
	db, err := m.db.BeginTx(ctx)

	txlocked := &lockedConsole{&sync.Mutex{}, db}
	return &lockedConsoleTx{m, txlocked, db, sync.Once{}}, err
}

// lockedConsoleTx extends Database with transaction scope
type lockedConsoleTx struct {
	parent *lockedConsole
	*lockedConsole
	tx   console.DBTx
	once sync.Once
}

// Commit is a method for committing and closing transaction
func (db *lockedConsoleTx) Commit() error {
	err := db.tx.Commit()
	db.once.Do(db.parent.Unlock)
	return err
}

// Rollback is a method for rollback and closing transaction
func (db *lockedConsoleTx) Rollback() error {
	err := db.tx.Rollback()
	db.once.Do(db.parent.Unlock)
	return err
}


// BeginTx is a method for opening transaction
func (m *lockedOrders) BeginTx(ctx context.Context) (orders.DBTx, error) {
	m.Lock()
	db, err := m.db.BeginTx(ctx)

	txlocked := &lockedOrders{&sync.Mutex{}, db}
	return &lockedOrdersTx{m, txlocked, db, sync.Once{}}, err
}

// lockedOrdersTx extends Database with transaction scope
type lockedOrdersTx struct {
	parent *lockedOrders
	*lockedOrders
	tx   orders.DBTx
	once sync.Once
}

// Commit is a method for committing and closing transaction
func (db *lockedOrdersTx) Commit() error {
	err := db.tx.Commit()
	db.once.Do(db.parent.Unlock)
	return err
}

// Rollback is a method for rollback and closing transaction
func (db *lockedOrdersTx) Rollback() error {
	err := db.tx.Rollback()
	db.once.Do(db.parent.Unlock)
	return err
}
