package security

import (
	"sync"
)

type Rbac struct {
	mutex sync.Mutex
	roles Roles
}

func New() *Rbac {
	return &Rbac{
		roles: make(Roles),
	}
}

// cek eksistensi role di rbac
func (rbac *Rbac) roleExist(role Role) bool {
	if _, ok := rbac.roles[role.Name()]; ok {
		return true
	}
	return false
}

// Add role
func (rbac *Rbac) AddRole(role Role) (err error) {
	rbac.mutex.Lock()
	defer rbac.mutex.Unlock()
	// tak ada role di roles
	if !rbac.roleExist(role) {
		rbac.roles[role.Name()] = role
	} else {
		err = ErrRoleExist
	}
	return
}

// remove role from rbac
func (rbac *Rbac) RemoveRole(role Role) (err error) {
	rbac.mutex.Lock()
	defer rbac.mutex.Unlock()
	// role ada di roles
	if rbac.roleExist(role) {
		delete(rbac.roles, role.Name())
	} else {
		err = ErrRoleNotExist
	}
	return
}

// Grant a permission to the specified role.
func (rbac *Rbac) Grant(role Role, perm Permission) (err error) {
	rbac.mutex.Lock()
	defer rbac.mutex.Unlock()
	if rbac.roleExist(role) {
		if err := role.Assign(perm); err != nil {
			return err
		}
	}
	return
}

// Revoke a permission from the specified role
func (rbac *Rbac) Revoke(role Role, perm Permission) (err error) {
	rbac.mutex.Lock()
	defer rbac.mutex.Unlock()
	if rbac.roleExist(role) {
		if err := role.Revoke(perm); err != nil {
			return err
		}
	}
	return nil
}

// Test whether the given role has access with the specified permission
func (rbac *Rbac) Check(role Role, perm Permission) bool {
	rbac.mutex.Lock()
	defer rbac.mutex.Unlock()
	if rbac.roleExist(role) {
		if role.Permit(perm) {
			return true
		}
	}
	return false
}
