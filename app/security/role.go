package security

import (
	"errors"
	"sync"
)

var (
	ErrRoleNotExist = errors.New("Role does not exist")
	ErrRoleExist    = errors.New("Role has already existed")
)

type Role interface {
	Name() string
	Permit(Permission) bool
	Assign(Permission) error
	Revoke(Permission) error
}

type Roles map[string]Role

type StdRole struct {
	sync.RWMutex
	Id          string
	permissions Permissions
}

func NewStdRole(name string) *StdRole {
	role := &StdRole{
		Id:          name,
		permissions: make(Permissions),
	}
	return role
}

func (sr *StdRole) Name() string {
	return sr.Id
}

func (sr *StdRole) Permit(p Permission) (result bool) {
	sr.RLock()
	for _, rp := range sr.permissions {
		if rp.Match(p) {
			result = true
			break
		}
	}
	defer sr.RUnlock()
	return
}

func (sr *StdRole) Assign(p Permission) error {
	sr.Lock()
	defer sr.Unlock()
	sr.permissions[p.Name()] = p
	return nil
}

func (sr *StdRole) Revoke(p Permission) error {
	sr.Lock()
	defer sr.Unlock()
	delete(sr.permissions, p.Name())
	return nil
}

// Permissions returns all permissions into a slice.
func (sr *StdRole) Permissions() []Permission {
	sr.RLock()
	result := make([]Permission, 0, len(sr.permissions))
	for _, p := range sr.permissions {
		result = append(result, p)
	}
	defer sr.RUnlock()
	return result
}
