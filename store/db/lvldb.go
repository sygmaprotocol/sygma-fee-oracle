package db

import (
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

func (db *LvlDB) Close() error {
	return db.db.Close()
}
