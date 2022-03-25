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

func (c *ConversionRateStore) StoreConversionRate(conversionRate *types.ConversionRateResp) error {
	data, err := json.Marshal(conversionRate)
	if err != nil {
		return err
	}
	return c.db.Set(c.storeKeyFormat(conversionRate.OracleName, conversionRate.Base, conversionRate.Foreign), data)
}

func (c *ConversionRateStore) GetConversionRate(oracleName, baseCurrency, foreignCurrency string) (*types.ConversionRateResp, error) {
	data, err := c.db.Get(c.storeKeyFormat(oracleName, baseCurrency, foreignCurrency))
	if err != nil {
		if errors.Is(err, leveldb.ErrNotFound) {
			return nil, ErrNotFound
		}
		return nil, err
	}

	var conversionRateData *types.ConversionRateResp
	err = json.Unmarshal(data, &conversionRateData)
	if err != nil {
		return nil, err
	}

	return conversionRateData, nil
}

func (c *ConversionRateStore) GetConversionRateByCurrencyPair(base, foreign string) ([]types.ConversionRateResp, error) {
	key := bytes.Buffer{}
	key.WriteString(conversionRateStoreKeyPrefix)
	var dataReceiver *types.ConversionRateResp

	conversionRateData, err := c.db.GetByPrefix(key.Bytes(), dataReceiver)
	if err != nil {
		if errors.Is(err, leveldb.ErrNotFound) {
			return nil, ErrNotFound
		}
		return nil, err
	}

	fmt.Println(conversionRateData)

	re := make([]types.ConversionRateResp, 0)
	for _, data := range conversionRateData {
		var cr types.ConversionRateResp
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
