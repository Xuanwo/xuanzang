package model

import (
	"strconv"

	"github.com/rs/zerolog/log"

	"github.com/Xuanwo/xuanzang/common/db"
)

// CloseTx will close tx with err.
// If err is nil, we will try to commit this tx.
// If err is not nil, we will rollback.
func CloseTx(tx *db.Tx, err error) {
	// If not writable, just rollback and skip.
	if !tx.Writable() {
		tx.Rollback()
		return
	}

	if err != nil {
		tx.Rollback()
		return
	}
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		log.Error().Msgf("Tx failed to commit for %v.", err)
	}
}

// FormatDocIDKey will format doc ID's key.
func FormatDocIDKey(id uint64) []byte {
	return []byte(docIDPrefix + strconv.FormatUint(id, 10))
}
