// Copyright 2021 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package store

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"strings"

	"github.com/ChainSafe/chainbridge-fee-oracle/types"
	"github.com/syndtr/goleveldb/leveldb"
)

const conversionRateStoreKeyPrefix = "conversionrate:"

type ConversionRateStore struct {
	db Store
}

func NewConversionRateStore(db Store) *ConversionRateStore {
	return &ConversionRateStore{
		db: db,
	}
}

func (c *ConversionRateStore) StoreConversionRate(conversionRate *types.ConversionRate) error {
	data, err := json.Marshal(conversionRate)
	if err != nil {
		return err
	}
	return c.db.Set(c.storeKeyFormat(conversionRate.OracleName, conversionRate.Base, conversionRate.Foreign), data)
}

func (c *ConversionRateStore) GetConversionRate(oracleName, baseCurrency, foreignCurrency string) (*types.ConversionRate, error) {
	data, err := c.db.Get(c.storeKeyFormat(oracleName, baseCurrency, foreignCurrency))
	if err != nil {
		if errors.Is(err, leveldb.ErrNotFound) {
			return nil, ErrNotFound
		}
		return nil, err
	}

	var conversionRateData *types.ConversionRate
	err = json.Unmarshal(data, &conversionRateData)
	if err != nil {
		return nil, err
	}

	return conversionRateData, nil
}

func (c *ConversionRateStore) GetConversionRatesByCurrencyPair(base, foreign string) ([]types.ConversionRate, error) {
	key := bytes.Buffer{}
	key.WriteString(conversionRateStoreKeyPrefix)
	var dataReceiver *types.ConversionRate

	conversionRateData, err := c.db.GetByPrefix(key.Bytes(), dataReceiver)
	if err != nil {
		return nil, err
	}

	re := make([]types.ConversionRate, 0)
	for _, data := range conversionRateData {
		var cr types.ConversionRate
		err = mapstructure.Decode(data, &cr)
		if err != nil {
			return nil, err
		}
		if cr.Base == base && cr.Foreign == foreign {
			re = append(re, cr)
		}
	}

	return re, nil
}

func (c *ConversionRateStore) storeKeyFormat(oracleName, baseCurrency, foreignCurrency string) []byte {
	key := bytes.Buffer{}
	key.WriteString(fmt.Sprintf("%s%s:%s:%s", conversionRateStoreKeyPrefix, strings.ToLower(oracleName),
		strings.ToLower(baseCurrency), strings.ToLower(foreignCurrency)))

	return key.Bytes()
}
