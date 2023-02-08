// The Licensed Work is (c) 2022 Sygma
// SPDX-License-Identifier: BUSL-1.1

package config

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strings"
)

const TEST_SYMBOL = "ERC20TST"

type domainConfig struct {
	Domains []Domain `json:"domains"`
}

type Domain struct {
	ID                   int           `json:"id"`
	Name                 string        `json:"name"`
	BaseCurrencySymbol   string        `json:"nativeTokenSymbol"`
	BaseCurrencyDecimals int           `json:"nativeTokenDecimals"`
	Resources            []rawResource `json:"resources"`
}

func newDomain(id int, name, baseCurrencySymbol string, baseCurrencyDecimals int) *Domain {
	return &Domain{
		ID:                   id,
		Name:                 name,
		BaseCurrencySymbol:   baseCurrencySymbol,
		BaseCurrencyDecimals: baseCurrencyDecimals,
	}
}

// loadDomainsFromFile registers and load all pre-defined domains from file
func loadDomainsFromFile(domainConfigPath string) (map[int]Domain, map[string]*Resource, error) {
	domainData, err := ioutil.ReadFile(filepath.Clean(domainConfigPath))
	if err != nil {
		return nil, nil, err
	}

	return parseDomains(domainData)
}

// loadDomainsFromNetwork registers and load all pre-defined domains from config stored on IPFS
func loadDomainsFromNetwork(url string) (map[int]Domain, map[string]*Resource, error) {
	resp, err := http.Get(url) // #nosec G107
	if err != nil {
		return nil, nil, err
	}

	domainData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, err
	}

	return parseDomains(domainData)
}

func parseDomains(domainData []byte) (map[int]Domain, map[string]*Resource, error) {
	var content domainConfig

	err := json.Unmarshal(domainData, &content)
	if err != nil {
		return nil, nil, err
	}

	domains := make(map[int]Domain)
	resources := make(map[string]*Resource)
	for _, domain := range content.Domains {
		domains[domain.ID] = *newDomain(domain.ID, domain.Name, domain.BaseCurrencySymbol, domain.BaseCurrencyDecimals)
		parseResources(resources, domain)
	}

	return domains, resources, nil
}

func parseResources(resources map[string]*Resource, domain Domain) {
	for _, resource := range domain.Resources {
		domainInfo := DomainInfo{
			Decimals: resource.Decimals,
		}

		storedResource, ok := resources[strings.ToLower(resource.ID)]
		if !ok {
			var r *Resource
			if resource.Symbol == TEST_SYMBOL {
				r = newResource(resource.ID, "usdt")
			} else {
				r = newResource(resource.ID, resource.Symbol)
			}

			r.DomainInfo[domain.ID] = &domainInfo
			resources[strings.ToLower(resource.ID)] = r
		} else {
			storedResource.DomainInfo[domain.ID] = &domainInfo
		}
	}

	// add native resource
	resource := newResource(ResourceIDBuilder(NativeCurrencyAddr, domain.ID), domain.BaseCurrencySymbol)
	resource.DomainInfo[domain.ID] = &DomainInfo{Decimals: domain.BaseCurrencyDecimals}
	resources[ResourceIDBuilder(NativeCurrencyAddr, domain.ID)] = resource
}
