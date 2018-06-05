package security

type Permission interface {
	Name() string
	Match(Permission) bool
}

type Permissions map[string]Permission

type StdPermission struct {
	Name      string
	resources Resources
}

func NewStdPermission(name string) Permission {
	return &StdPermission{name}
}

func (sp *StdPermission) Name() string {
	return sp.Name
}

func (sp *StdPermission) Match(op Permission) bool {
	return sp.Name == op.Name()
}
