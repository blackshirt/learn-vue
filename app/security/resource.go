package security

import (
	"errors"
	"sync"
)

var (
	ErrPermNotExist = errors.New("Permission does not exist")
	// ErrRoleExist occurred if a role shouldn't be found
	ErrPermExist = errors.New("Permission has already existed")
)

type Resource struct {
	Name string
}
type Resources map[string]Resource

func NewResource(name string) *Resource {
	res := &Resource{
		Name: name,
	}
}
