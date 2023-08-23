// The Licensed Work is (c) 2022 Sygma
// SPDX-License-Identifier: BUSL-1.1

package store

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/ChainSafe/sygma-fee-oracle/types"
	"github.com/mitchellh/mapstructure"
	"github.com/syndtr/goleveldb/leveldb"
)

const gasPriceStoreKeyPrefix = "gasprice:"

type GasPriceStore struct {
	db Store
}

func NewGasPriceStore(db Store) *GasPriceStore {
	return &GasPriceStore{
		db: db,
	}
}

func (g *GasPriceStore) StoreGasPrice(gasPrice *types.GasPrices) error {
	data, err := json.Marshal(gasPrice)
	if err != nil {
		return err
	}

	return g.db.Set(g.storeKeyFormat(gasPrice.OracleSource, gasPrice.DomainID), data)
}

func (g *GasPriceStore) GetGasPrice(OracleSource string, domainID int) (*types.GasPrices, error) {
	gasPriceData, err := g.db.Get(g.storeKeyFormat(OracleSource, domainID))
	if err != nil {
		if errors.Is(err, leveldb.ErrNotFound) {
			return nil, ErrNotFound
		}
		return nil, err
	}

	var gasPrice *types.GasPrices
	err = json.Unmarshal(gasPriceData, &gasPrice)
	if err != nil {
		return nil, err
	}

	return gasPrice, nil
}

func (g *GasPriceStore) GetGasPricesByDomain(domainID int) ([]types.GasPrices, error) {
	key := bytes.Buffer{}
	key.WriteString(gasPriceStoreKeyPrefix)
	var dataReceiver *types.GasPrices

	gasPriceData, err := g.db.GetByPrefix(key.Bytes(), dataReceiver)
	if err != nil {
		return nil, err
	}

	re := make([]types.GasPrices, 0)
	for _, data := range gasPriceData {
		var gp types.GasPrices
		err = mapstructure.Decode(data, &gp)
		if err != nil {
			return nil, err
		}
		if gp.DomainID == domainID {
			re = append(re, gp)
		}
	}

	return re, nil
}

func (g *GasPriceStore) storeKeyFormat(oracleSource string, domainID int) []byte {
	key := bytes.Buffer{}
	key.WriteString(fmt.Sprintf("%s%s:%d", gasPriceStoreKeyPrefix, strings.ToLower(oracleSource), domainID))

	return key.Bytes()
}
