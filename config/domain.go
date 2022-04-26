// Copyright 2021 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package config

type domain struct {
	ID                   int    `json:"id"`
	Name                 string `json:"name"`
	BaseCurrencyFullName string `json:"baseCurrencyFullName"`
	BaseCurrencySymbol   string `json:"baseCurrencySymbol"`
}

func newDomain(id int, name, baseCurrencyFullName, baseCurrencySymbol string) *domain {
	return &domain{
		ID:                   id,
		Name:                 name,
		BaseCurrencyFullName: baseCurrencyFullName,
		BaseCurrencySymbol:   baseCurrencySymbol,
	}
}

// loadDomains registers and load all pre-defined domains
func loadDomains() map[int]domain {
	domains := make(map[int]domain, 0)

	domains[1] = *newDomain(1, "ethereum", "ether", "eth")
	domains[2] = *newDomain(2, "polygon", "matic", "matic")

	return domains
}
