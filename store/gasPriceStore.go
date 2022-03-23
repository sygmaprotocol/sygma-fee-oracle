package store

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ChainSafe/chainbridge-fee-oracle/types"
	"github.com/syndtr/goleveldb/leveldb"
)

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

func (g *GasPriceStore) storeKeyFormat(oracleName, domainID string) []byte {
	return []byte(fmt.Sprintf("%s:%s", oracleName, domainID))
}
