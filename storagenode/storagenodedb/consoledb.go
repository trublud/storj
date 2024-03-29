// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information.

package storagenodedb

import (
	"context"
	"time"

	"github.com/zeebo/errs"

	"storj.io/storj/internal/date"
	"storj.io/storj/pkg/pb"
	"storj.io/storj/pkg/storj"
	"storj.io/storj/storagenode/console"
)

type consoleDB struct {
	SQLDB
}

// Bandwidth returns consoledb as console.Bandwidth
func (db *consoleDB) Bandwidth() console.Bandwidth {
	return db
}

// newConsoleDB returns a new instance of consoledb initialized with the specified database.
func newConsoleDB(db SQLDB) *consoleDB {
	return &consoleDB{
		SQLDB: db,
	}
}

// GetDaily returns slice of daily bandwidth usage for provided time range,
// sorted in ascending order for particular satellite
func (db *consoleDB) GetDaily(ctx context.Context, satelliteID storj.NodeID, from, to time.Time) (_ []console.BandwidthUsed, err error) {
	defer mon.Task()(&ctx)(&err)

	since, _ := date.DayBoundary(from.UTC())
	_, before := date.DayBoundary(to.UTC())

	return db.getDailyBandwidthUsed(ctx,
		"WHERE satellite_id = ? AND ? <= created_at AND created_at <= ?",
		satelliteID, since, before)
}

// GetDaily returns slice of daily bandwidth usage for provided time range,
// sorted in ascending order
func (db *consoleDB) GetDailyTotal(ctx context.Context, from, to time.Time) (_ []console.BandwidthUsed, err error) {
	defer mon.Task()(&ctx)(&err)

	since, _ := date.DayBoundary(from.UTC())
	_, before := date.DayBoundary(to.UTC())

	return db.getDailyBandwidthUsed(ctx,
		"WHERE ? <= created_at AND created_at <= ?",
		since, before)
}

// getDailyBandwidthUsed returns slice of grouped by date bandwidth usage
// sorted in ascending order and applied condition if any
func (db *consoleDB) getDailyBandwidthUsed(ctx context.Context, cond string, args ...interface{}) (_ []console.BandwidthUsed, err error) {
	defer mon.Task()(&ctx)(&err)

	query := `SELECT action, SUM(amount), created_at
				FROM bandwidth_usage
				` + cond + `
				GROUP BY DATE(created_at), action
				ORDER BY created_at ASC`

	rows, err := db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	defer func() {
		err = errs.Combine(err, rows.Close())
	}()

	var dates []time.Time
	dailyBandwidth := make(map[time.Time]*console.BandwidthUsed)

	for rows.Next() {
		var action int32
		var amount int64
		var createdAt time.Time

		err = rows.Scan(&action, &amount, &createdAt)
		if err != nil {
			return nil, err
		}

		from, to := date.DayBoundary(createdAt)

		bandwidthUsed, ok := dailyBandwidth[from]
		if !ok {
			bandwidthUsed = &console.BandwidthUsed{
				From: from,
				To:   to,
			}

			dates = append(dates, from)
			dailyBandwidth[from] = bandwidthUsed
		}

		switch pb.PieceAction(action) {
		case pb.PieceAction_GET:
			bandwidthUsed.Egress.Usage = amount
		case pb.PieceAction_GET_AUDIT:
			bandwidthUsed.Egress.Audit = amount
		case pb.PieceAction_GET_REPAIR:
			bandwidthUsed.Egress.Repair = amount
		case pb.PieceAction_PUT:
			bandwidthUsed.Ingress.Usage = amount
		case pb.PieceAction_PUT_REPAIR:
			bandwidthUsed.Ingress.Repair = amount
		}
	}

	var bandwidthUsedList []console.BandwidthUsed
	for _, date := range dates {
		bandwidthUsedList = append(bandwidthUsedList, *dailyBandwidth[date])
	}

	return bandwidthUsedList, nil
}
