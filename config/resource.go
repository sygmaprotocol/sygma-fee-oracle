// The Licensed Work is (c) 2022 Sygma
// SPDX-License-Identifier: BUSL-1.1

package config

import (
	"fmt"
	"strings"
)

const (
	NativeCurrencyAddr = "0x0000000000000000000000000000000000000000"
)

type rawResource struct {
	ID       string `json:"resourceId"`
	Symbol   string `json:"symbol"`
	Decimals int    `json:"decimals"`
}

type DomainInfo struct {
	Decimals int `json:"decimals"`
}

type Resource struct {
	ID         string
	Symbol     string
	Decimals   int
	DomainInfo map[int]*DomainInfo
}

func newResource(resourceID string, symbol string) *Resource {
	return &Resource{
		ID:         resourceID,
		Symbol:     symbol,
		DomainInfo: make(map[int]*DomainInfo),
	}
}

// ResourceIDBuilder builds the resourceID according to fee handler contract
func ResourceIDBuilder(tokenAddress string, domainId int) string {
	return fmt.Sprintf("%s%d", strings.ToLower(tokenAddress), domainId)
}
