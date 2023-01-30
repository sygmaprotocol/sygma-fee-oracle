// The Licensed Work is (c) 2022 Sygma
// SPDX-License-Identifier: BUSL-1.1

package config

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
	"strings"
)

type domainConfigFile struct {
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

// loadDomains registers and load all pre-defined domains
func loadDomains(domainConfigPath string) (map[int]Domain, map[string]*Resource) {
	domainData, err := ioutil.ReadFile(filepath.Clean(domainConfigPath))
	if err != nil {
		panic(ErrLoadDomainConfig.Wrap(err))
	}

	return parseDomains(domainData)
}

func parseDomains(domainData []byte) (map[int]Domain, map[string]*Resource) {
	var content domainConfigFile

	err := json.Unmarshal(domainData, &content)
	if err != nil {
		panic(ErrLoadDomainConfig.Wrap(err))
	}

	domains := make(map[int]Domain)
	resources := make(map[string]*Resource)
	for _, domain := range content.Domains {
		domains[domain.ID] = *newDomain(domain.ID, domain.Name, domain.BaseCurrencySymbol, domain.BaseCurrencyDecimals)
		parseResources(resources, domain)
	}

	return domains, resources
}

func parseResources(resources map[string]*Resource, domain Domain) {
	for _, resource := range domain.Resources {
		domainInfo := DomainInfo{
			Decimals: resource.Decimals,
		}

		storedResource, ok := resources[strings.ToLower(resource.ID)]
		if !ok {
			resource := newResource(resource.ID, resource.Symbol)
			resource.DomainInfo[domain.ID] = &domainInfo
			resources[strings.ToLower(resource.ID)] = resource
		} else {
			storedResource.DomainInfo[domain.ID] = &domainInfo
		}
	}

	// add native resource
	resource := newResource(ResourceIDBuilder(NativeCurrencyAddr, domain.ID), domain.BaseCurrencySymbol)
	resource.DomainInfo[domain.ID] = &DomainInfo{Decimals: domain.BaseCurrencyDecimals}
	resources[ResourceIDBuilder(NativeCurrencyAddr, domain.ID)] = resource
}
