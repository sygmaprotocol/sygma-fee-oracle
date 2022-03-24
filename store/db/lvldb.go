package db

import (
	"bytes"
	"github.com/ChainSafe/chainbridge-fee-oracle/store"
	"github.com/pkg/errors"
	"github.com/syndtr/goleveldb/leveldb"
)

var _ store.Store = (*LvlDB)(nil)

type LvlDB struct {
	db *leveldb.DB
}

func NewLvlDB(path string) (*LvlDB, error) {
	ldb, err := leveldb.OpenFile(path, nil)
	if err != nil {
		return nil, errors.Wrap(err, "failed with levelDB.OpenFile")
	}
	return &LvlDB{db: ldb}, nil
}

func (db *LvlDB) Get(key []byte) ([]byte, error) {
	return db.db.Get(key, nil)
}

func (db *LvlDB) Set(key []byte, value []byte) error {
	return db.db.Put(key, value, nil)
}

// GetBySuffix get all items from db with given suffix
// e.g:
//	1. key pattern with `*:eth:usdt` for conversion rate, suffix is `:eth:usdt`
//  2. key pattern with `*:ethereum` for gas price, suffix is `:ethereum`
func (db *LvlDB) GetBySuffix(suffix []byte) ([][]byte, error) {
	re := make([][]byte, 0)
	iter := db.db.NewIterator(nil, nil)
	for iter.Next() {
		if bytes.HasSuffix(iter.Key(), suffix) {
			re = append(re, iter.Value())
		}
	}
	iter.Release()
	err := iter.Error()
	if err != nil {
		return nil, err
	}

	return re, nil
}

func (db *LvlDB) Close() error {
	return db.db.Close()
}
