// Copyright 2021 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package db

import (
	"encoding/json"
	"github.com/ChainSafe/chainbridge-fee-oracle/store"
	"github.com/pkg/errors"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/util"
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

// GetByPrefix get all items from db with given prefix
// dataReceiver is the var of the data type after Unmarshal
func (db *LvlDB) GetByPrefix(prefix []byte, dataReceiver interface{}) ([]interface{}, error) {
	re := make([]interface{}, 0)
	iter := db.db.NewIterator(util.BytesPrefix(prefix), nil)
	for iter.Next() {
		val := iter.Value()

		err := json.Unmarshal(val, &dataReceiver)
		if err != nil {
			return nil, err
		}

		re = append(re, dataReceiver)
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
