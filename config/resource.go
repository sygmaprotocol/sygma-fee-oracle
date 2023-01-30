// The Licensed Work is (c) 2022 Sygma
// SPDX-License-Identifier: BUSL-1.1

package config

import (
	"encoding/json"
	"fmt"
	"strings"
)

const (
	NativeCurrencyAddr = "0x0000000000000000000000000000000000000000"
)

type resourceConfigFile struct {
	Resources []Resource `json:"resources"`
}

type DomainInfo struct {
	Decimals int `json:"decimals"`
}

type Resource struct {
	ID         string `json:"id"`
	Symbol     string `json:"symbol"`
	DomainInfo map[int]*DomainInfo
}

func newResource(resourceID string, symbol string) *Resource {
	return &Resource{
		ID:     resourceID,
		Symbol: symbol,
	}
}

// ResourceIDBuilder builds the resourceID according to fee handler contract
func ResourceIDBuilder(tokenAddress string, domainId int) string {
	return fmt.Sprintf("%s%d", strings.ToLower(tokenAddress), domainId)
}

func parseResources(resourceData []byte) map[string]Resource {
	var content resourceConfigFile
	err := json.Unmarshal(resourceData, &content)
	if err != nil {
		panic(ErrLoadResourceConfig.Wrap(err))
	}

	resources := make(map[string]Resource, 0)
	for _, resource := range content.Resources {
		resources[strings.ToLower(resource.ID)] =
			*newResource(resource.ID, resource.Symbol)
	}

	return resources
}
