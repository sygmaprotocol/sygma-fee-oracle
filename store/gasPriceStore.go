package store

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/ChainSafe/chainbridge-fee-oracle/types"
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

func (g *GasPriceStore) StoreGasPrice(gasPrice *types.GasPricesResp) error {
	data, err := json.Marshal(gasPrice)
	if err != nil {
		return err
	}

	return g.db.Set(g.storeKeyFormat(gasPrice.OracleName, gasPrice.DomainId), data)
}

func (g *GasPriceStore) GetGasPrice(oracleName, domainID string) (*types.GasPricesResp, error) {
	gasPriceData, err := g.db.Get(g.storeKeyFormat(oracleName, domainID))
	if err != nil {
		if errors.Is(err, leveldb.ErrNotFound) {
			return nil, ErrNotFound
		}
		return nil, err
	}

	var gasPrice *types.GasPricesResp
	err = json.Unmarshal(gasPriceData, &gasPrice)
	if err != nil {
		return nil, err
	}

	return gasPrice, nil
}

func (g *GasPriceStore) GetGasPriceByDomain(domainId string) ([]types.GasPricesResp, error) {
	key := bytes.Buffer{}
	key.WriteString(gasPriceStoreKeyPrefix)
	var dataReceiver *types.GasPricesResp

	gasPriceData, err := g.db.GetByPrefix(key.Bytes(), dataReceiver)
	if err != nil {
		if errors.Is(err, leveldb.ErrNotFound) {
			return nil, ErrNotFound
		}
		return nil, err
	}

	re := make([]types.GasPricesResp, 0)
	for _, data := range gasPriceData {
		var gp types.GasPricesResp
		err = mapstructure.Decode(data, &gp)
		if err != nil {
			return nil, err
		}
		if gp.DomainId == domainId {
			re = append(re, gp)
		}
	}

	return re, nil
}

func (g *GasPriceStore) storeKeyFormat(oracleName, domainID string) []byte {
	key := bytes.Buffer{}
	key.WriteString(fmt.Sprintf("%s%s:%s", gasPriceStoreKeyPrefix, strings.ToLower(oracleName), strings.ToLower(domainID)))

	return key.Bytes()
}
