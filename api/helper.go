// The Licensed Work is (c) 2022 Sygma
// SPDX-License-Identifier: BUSL-1.1

package api

import (
	"fmt"
	"strconv"

	"github.com/ChainSafe/sygma-fee-oracle/config"
	"github.com/ChainSafe/sygma-fee-oracle/types"
)

func (h *Handler) dataExpirationManager(baseTimestamp int64) int64 {
	return baseTimestamp + h.conf.DataValidInterval
}

func (h *Handler) parseDomains(fromID string, toID string) (from *config.Domain, to *config.Domain, err error) {
	fromDomainID, err := strconv.Atoi(fromID)
	if err != nil {
		return from, to, err
	}
	toDomainID, err := strconv.Atoi(toID)
	if err != nil {
		return from, to, err
	}
	if fromDomainID == toDomainID {
		return from, to, fmt.Errorf("from and to domain equal")
	}

	from, err = h.conf.Domain(fromDomainID)
	if err != nil {
		return from, to, err
	}
	to, err = h.conf.Domain(toDomainID)
	if err != nil {
		return from, to, err
	}

	return from, to, nil
}

func (h *Handler) parseResource(resourceID string) (*config.Resource, error) {
	resource, err := h.conf.Resource(resourceID)
	if err != nil {
		return &config.Resource{}, err
	}

	return resource, nil
}

func (h *Handler) calculateTokenRate(resource *config.Resource, ber *types.ConversionRate, from *config.Domain, to *config.Domain) (*types.ConversionRate, error) {
	ter := &types.ConversionRate{}
	if resource.Symbol == from.BaseCurrencySymbol {
		ter = ber
	} else {
		if resource.Symbol == to.BaseCurrencySymbol {
			ter.Rate = 1.0
		} else {
			ter, err := h.consensus.FilterLocalConversionRateData(h.conversionRateStore, to.BaseCurrencySymbol, resource.Symbol)
			if err != nil {
				return ter, err
			}

			return ter, nil
		}
	}

	return ter, nil
}
