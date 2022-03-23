package oracle

import _ "github.com/golang/mock/mockgen/model"

type Oracle interface {
	Name() string
	IsEnabled() bool
}
