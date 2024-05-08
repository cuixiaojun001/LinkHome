package cache

import (
	"fmt"
)

var factories = make(map[string]Factory)

type Factory interface {
	New(cfg map[string]interface{}) (Cache, error)
}

func Register(name string, factory Factory) {
	if factory == nil {
		panic("Must not provide nil Factory")
	}
	_, registered := factories[name]
	if registered {
		panic(fmt.Sprintf("Factory named %s already registered", name))
	}

	factories[name] = factory
}

func New(name string, parameters map[string]interface{}) (Cache, error) {
	factory, ok := factories[name]
	if !ok {
		return nil, InvalidCacheError{Name: name}
	}
	return factory.New(parameters)
}

type InvalidCacheError struct {
	Name string
}

func (err InvalidCacheError) Error() string {
	return fmt.Sprintf("Cache not registered: %s", err.Name)
}
