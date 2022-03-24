package store

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ChainSafe/chainbridge-fee-oracle/types"
	"github.com/syndtr/goleveldb/leveldb"
	"strings"
)

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

func (c *ConversionRateStore) GetConversionRateByCurrencyPair(base, foreign string) ([]*types.ConversionRateResp, error) {
	conversionRateData, err := c.db.GetBySuffix([]byte(fmt.Sprintf(":%s:%s", strings.ToLower(base), strings.ToLower(foreign))))
	if err != nil {
		if errors.Is(err, leveldb.ErrNotFound) {
			return nil, ErrNotFound
		}
		return nil, err
	}

	re := make([]*types.ConversionRateResp, 0)
	for _, data := range conversionRateData {
		var cr *types.ConversionRateResp
		err = json.Unmarshal(data, &cr)
		if err != nil {
			return nil, err
		}
		re = append(re, cr)
	}

	return re, nil
}

func (c *ConversionRateStore) storeKeyFormat(oracleName, baseCurrency, foreignCurrency string) []byte {
	return []byte(fmt.Sprintf("%s:%s:%s", strings.ToLower(oracleName), strings.ToLower(baseCurrency), strings.ToLower(foreignCurrency)))
}
