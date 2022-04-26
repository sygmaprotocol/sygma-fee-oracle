// Copyright 2021 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package config

type resource struct {
	ID     int    `json:"id"`
	Symbol string `json:"symbol"`
	Domain domain `json:"domain"`
}

func newResource(id int, symbol string, domain domain) *resource {
	return &resource{
		ID:     id,
		Symbol: symbol,
		Domain: domain,
	}
}

// loadResources registers and load all pre-defined resources
func loadResources(domains map[int]domain) map[int]resource {
	resources := make(map[int]resource, 0)

	resources[1] = *newResource(1, "eth", domains[1])
	resources[2] = *newResource(2, "usdt", domains[1])
	resources[3] = *newResource(3, "matic", domains[1])
	resources[4] = *newResource(4, "doge", domains[1])

	resources[5] = *newResource(5, "eth", domains[2])
	resources[6] = *newResource(6, "usdt", domains[2])
	resources[7] = *newResource(7, "matic", domains[2])
	resources[8] = *newResource(8, "doge", domains[2])

	return resources
}
