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
	resources[3] = *newResource(3, "matic", domains[2])
	resources[4] = *newResource(4, "doge", domains[1])

	return resources
}
