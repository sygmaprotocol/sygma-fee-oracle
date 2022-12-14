// The Licensed Work is (c) 2022 Sygma
// SPDX-License-Identifier: BUSL-1.1

package config

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
)

type domainConfigFile struct {
	Domains []domain `json:"domains"`
}

type domain struct {
	ID                   int    `json:"id"`
	Name                 string `json:"name"`
	BaseCurrencyFullName string `json:"baseCurrencyFullName"`
	BaseCurrencySymbol   string `json:"baseCurrencySymbol"`
	AddressPrefix        string `json:"addressPrefix"`
}

func newDomain(id int, name, baseCurrencyFullName, baseCurrencySymbol, addressPrefix string) *domain {
	return &domain{
		ID:                   id,
		Name:                 name,
		BaseCurrencyFullName: baseCurrencyFullName,
		BaseCurrencySymbol:   baseCurrencySymbol,
		AddressPrefix:        addressPrefix,
	}
}

// loadDomains registers and load all pre-defined domains
func loadDomains(domainConfigPath string) map[int]domain {
	domainData, err := ioutil.ReadFile(filepath.Clean(domainConfigPath))
	if err != nil {
		panic(ErrLoadDomainConfig.Wrap(err))
	}

	return parseDomains(domainData)
}

func parseDomains(domainData []byte) map[int]domain {
	var content domainConfigFile

	err := json.Unmarshal(domainData, &content)
	if err != nil {
		panic(ErrLoadDomainConfig.Wrap(err))
	}

	domains := make(map[int]domain, 0)
	for _, domain := range content.Domains {
		domains[domain.ID] = *newDomain(domain.ID, domain.Name, domain.BaseCurrencyFullName, domain.BaseCurrencySymbol, domain.AddressPrefix)
	}

	return domains
}
