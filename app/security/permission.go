/*
Map resource to permission
*/
package security

import "errors"

var (
	ErrPermissionNotExist = errors.New("Permission does not exist")
	ErrPermissionExist    = errors.New("Permission has already existed")
)

type Permission interface {
	Name() string
	Match(Permission) bool
}

type Permissions map[string]Permission

type StdPermission struct {
	Id string
}

func NewStdPermission(name string) Permission {
	return &StdPermission{name}
}

func (sp *StdPermission) Name() string {
	return sp.Id
}

func (sp *StdPermission) Match(op Permission) bool {
	return sp.Id == op.Name()
}
