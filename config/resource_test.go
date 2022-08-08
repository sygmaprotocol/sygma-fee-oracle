package config

import (
	"github.com/stretchr/testify/suite"
	"io/ioutil"
	"path/filepath"
	"testing"
)

type ResourceConfigTestSuite struct {
	suite.Suite
}

func TestRunConversionRateStoreTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceConfigTestSuite))
}

func (s *ResourceConfigTestSuite) TestParseResource() {
	resourceData, err := ioutil.ReadFile(filepath.Clean("../fixtures/dev/resource.json"))
	s.Nil(err)
	domains := loadDomains("../fixtures/dev/domain.json")

	resource := parseResources(resourceData, domains)
	s.Equal(len(resource), 1)
	ds := resource["0x0000000000000000000000000000000000000000000000000000000000000001"].Domain
	s.Equal(ds.Name, "ethereum")
	// WE have only one domain and resource

}
