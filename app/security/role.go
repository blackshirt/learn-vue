package security

import "sync"

type Role interface {
	Name() string
	Permit(Permission) bool
}

type Roles map[string]Role

type StdRole struct {
	sync.RWMutex
	Name        string
	permissions Permissions
}

func NewStdRole(name string) *StdRole {
	role := &StdRole{
		name:        name,
		permissions: make(Permissions),
	}
	return role
}

func (sr *StdRole) Name() string {
	return sr.Name
}

func (sr *StdRole) Permit(p Permission) (result bool) {
	sr.RLock()
	for _, rp := range sr.permissions {
		if rp.Match(p) {
			result = true
			break
		}
	}
	sr.Unlock()
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
	sr.RUnlock()
	return result
}
