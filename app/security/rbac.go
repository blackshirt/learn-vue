package security

import (
	"errors"
	"sync"
)

var (
	ErrRoleNotExist = errors.New("Role does not exist")
	// ErrRoleExist occurred if a role shouldn't be found
	ErrRoleExist = errors.New("Role has already existed")
)

type Rbac struct {
	mutex     sync.Mutex
	roles     Roles
	resources Resources
}

func New() *Rbac {
	return &Rbac{
		roles:     make(Roles),
		resources: make(Resources),
	}
}

// Add role to the structure
func (rbac *Rbac) AddRole(role Role) (err error) {
	rbac.mutex.Lock()
	if _, ok := rbac.roles[role.Name()]; !ok {
		rbac.roles[role.Name()] = role
	} else {
		err = ErrRoleExist
	}
	rbac.mutex.Unlock()
	return
}

func (rbac *Rbac) RemoveRole(role Role) (err error) {
	rbac.mutex.Lock()
	defer rbac.mutex.Unlock()
	if _, ok := rbac.roles[role.Name()]; ok {
		delete(rbac.roles, role)
	} else {
		err = ErrRoleNotExist
	}
	return
}

func (rbac *Rbac) Grant(role Role, res Resource, perm Permission) (err error) {
	// Grant a permission over resource to the specified role.
	rbac.mutex.Lock()
	defer rbac.mutex.Unlock()
	// cek jika role ada didaftar
	if _, ok := rbac.roles[role.Name()]; ok {

	}
	return
}

func (rbac *Rbac) Revoke(role Role, res Resource, perm Permission) (err error) {
	// Revoke a permission over a resource from the specified role
}

func (rbac *Rbac) Check(role Role, res Resource, perm Permission) bool {
	// Test whether the given role has access to the resource with the specified permission
}
