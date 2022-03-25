package store

import "github.com/pkg/errors"

var (
	ErrNotFound = errors.New("key not found")
)

type Store interface {
	Set([]byte, []byte) error
	Get([]byte) ([]byte, error)
	GetByPrefix([]byte, interface{}) ([]interface{}, error)
	Close() error
}
