package model

import (
	"github.com/rs/zerolog/log"
	"github.com/vmihailenco/msgpack"

	"github.com/Xuanwo/xuanzang/constants"
	"github.com/Xuanwo/xuanzang/contexts"
)

const (
	docPrefix   = "d:"
	docIDPrefix = "di:"
)

// Doc is the base model in xuanzang.
type Doc struct {
	ID        uint64 `msgpack:"id"`
	UpdatedAt int64  `msgpack:"ua"`
	URL       string `msgpack:"u"`
	Title     string `msgpack:"t"`
}

// Save will save current url to database.
func (u *Doc) Save() (err error) {
	tx, err := contexts.DB.Begin(true)
	if err != nil {
		log.Error().Msgf("Start transaction failed for %v.", err)
		return
	}
	defer func() {
		CloseTx(tx, err)
	}()

	// Get bucket.
	b := tx.Bucket(constants.DefaultBucketName)

	content, err := msgpack.Marshal(u)
	if err != nil {
		return
	}

	err = b.Put([]byte(docPrefix+u.URL), content)
	if err != nil {
		return
	}
	return
}

// CreateDoc will create doc in database.
func CreateDoc(p string, updatedAt int64) (u *Doc, err error) {
	tx, err := contexts.DB.Begin(true)
	if err != nil {
		log.Error().Msgf("Start transaction failed for %v.", err)
		return
	}
	defer func() {
		CloseTx(tx, err)
	}()

	// Get bucket.
	b := tx.Bucket(constants.DefaultBucketName)

	id, err := b.NextSequence()
	if err != nil {
		return
	}

	u = &Doc{
		ID:        id,
		URL:       p,
		UpdatedAt: updatedAt,
	}

	content, err := msgpack.Marshal(u)
	if err != nil {
		return
	}

	err = b.Put([]byte(docPrefix+p), content)
	if err != nil {
		return
	}

	err = b.Put(FormatDocIDKey(id), []byte(p))
	if err != nil {
		return
	}
	return
}

// DeleteDoc will delete doc from database.
func DeleteDoc(p string) (err error) {
	tx, err := contexts.DB.Begin(true)
	if err != nil {
		log.Error().Msgf("Start transaction failed for %v.", err)
		return
	}
	defer func() {
		CloseTx(tx, err)
	}()

	// Get bucket.
	b := tx.Bucket(constants.DefaultBucketName)

	return b.Delete([]byte(docPrefix + p))
}

// DeleteDocID will delete a Doc ID.
func DeleteDocID(id uint64) (err error) {
	tx, err := contexts.DB.Begin(true)
	if err != nil {
		log.Error().Msgf("Start transaction failed for %v.", err)
		return
	}
	defer func() {
		CloseTx(tx, err)
	}()

	// Get bucket.
	b := tx.Bucket(constants.DefaultBucketName)

	return b.Delete(FormatDocIDKey(id))
}

// GetDoc will get a Doc from database.
func GetDoc(p string) (u *Doc, err error) {
	tx, err := contexts.DB.Begin(false)
	if err != nil {
		log.Error().Msgf("Start transaction failed for %v.", err)
		return
	}
	defer func() {
		CloseTx(tx, err)
	}()

	// Get bucket.
	b := tx.Bucket(constants.DefaultBucketName)

	content := b.Get([]byte(docPrefix + p))
	if content == nil {
		return nil, nil
	}

	u = &Doc{}
	err = msgpack.Unmarshal(content, u)
	if err != nil {
		return
	}
	return
}

// GetDocByID will get a doc by it's id.
func GetDocByID(id uint64) (u *Doc, err error) {
	tx, err := contexts.DB.Begin(true)
	if err != nil {
		log.Error().Msgf("Start transaction failed for %v.", err)
		return
	}
	defer func() {
		CloseTx(tx, err)
	}()

	// Get bucket.
	b := tx.Bucket(constants.DefaultBucketName)

	content := b.Get(FormatDocIDKey(id))
	if content == nil {
		return nil, nil
	}

	content = b.Get([]byte(docPrefix + string(content)))
	if content == nil {
		return nil, nil
	}

	u = &Doc{}
	err = msgpack.Unmarshal(content, u)
	if err != nil {
		return
	}
	return
}
