package db

import (
	"time"

	"github.com/boltdb/bolt"
	"github.com/rs/zerolog/log"

	"github.com/Xuanwo/xuanzang/constants"
)

// Database stores database connection.
type Database struct {
	*bolt.DB
}

// DatabaseOptions stores database options.
type DatabaseOptions struct {
	Address string
}

// NewDB will create a new database connection.
func NewDB(opt *DatabaseOptions) (d *Database, err error) {
	client, err := bolt.Open(opt.Address, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		log.Error().Msgf("Open database failed for %v.", err)
		return
	}

	d = &Database{DB: client}
	log.Debug().Msgf("Connected to Bolt database %s", opt.Address)
	return
}

// Init will init current database.
func (db *Database) Init() (err error) {
	tx, err := db.Begin(true)
	if err != nil {
		log.Error().Msgf("Start writable transaction failed for %v.", err)
		return
	}
	defer tx.Rollback()

	// Create task list bucket.
	_, err = tx.CreateBucketIfNotExists([]byte(constants.DefaultBucketName))
	if err != nil {
		log.Error().Msgf("Create task bucket failed for %v.", err)
		return
	}

	err = tx.Commit()
	if err != nil {
		log.Error().Msgf("DB commit failed for %v.", err)
		return
	}

	return
}
