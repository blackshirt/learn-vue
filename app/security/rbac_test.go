package security

import (
	"testing"
)

var (
	anonymouserole = NewStdRole("anonymous")
	adminrole      = NewStdRole("admin")
	nonexistrole   = NewStdRole("nonexist")
	createUser     = NewStdPermission("createUser")
	deleteUser     = NewStdPermission("deleteUser")
	readUser       = NewStdPermission("readUser")
	pingUser       = NewStdPermission("pingUser")
	rbac           *Rbac
)

func assert(t *testing.T, err error) {
	if err != nil {
		t.Fatal(err)
	}
}

func TestRbacSetup(t *testing.T) {
	rbac = New()
}

func TestRbacAddRole(t *testing.T) {
	// test adding admin role to rbac
	assert(t, rbac.AddRole(adminrole))
	// test for adding already added role
	if err := rbac.AddRole(adminrole); err != ErrRoleExist {
		t.Error(" A role can not be added, already exist")
	}
	// add anonymouse role
	assert(t, rbac.AddRole(anonymouserole))
}

func TestRbacExist(t *testing.T) {
	if !rbac.roleExist(adminrole) {
		t.Fatalf("%s should exist, it already added before", adminrole.Name())
	}
	if rbac.roleExist(nonexistrole) {
		t.Fatalf("%s should not exist", nonexistrole.Name())
	}
}

func TestRbacRemoveRole(t *testing.T) {
	// test removing admin role
	assert(t, rbac.RemoveRole(adminrole))
	// cek still eksis atau gak
	if rbac.roleExist(adminrole) {
		t.Fatal("removing role failed, still exist")
	}
	// test if the role was non exist in rbac
	if err := rbac.RemoveRole(nonexistrole); err != ErrRoleNotExist {
		t.Fatalf("%s needed", ErrRoleNotExist)
	}
}

func TestRbacGrant(t *testing.T) {
	assert(t, rbac.AddRole(adminrole))
	if !rbac.roleExist(adminrole) {
		t.Fatalf("%s should exist", adminrole.Name())
	}
	assert(t, rbac.Grant(adminrole, createUser))
	assert(t, rbac.Grant(adminrole, deleteUser))
}

func TestRbacRevoke(t *testing.T) {
	if !rbac.roleExist(adminrole) {
		t.Fatalf("%s should exist", adminrole.Name())
	}
	if !rbac.Check(adminrole, createUser) {
		t.Fatalf("%s should have %s", adminrole.Name(), createUser.Name())
	}
	assert(t, rbac.Revoke(adminrole, createUser))
	assert(t, rbac.Revoke(adminrole, deleteUser))
}

func TestRbacCheck(t *testing.T) {

	if !rbac.roleExist(adminrole) {
		t.Fatalf("%s should exist", adminrole.Name())
	}
	assert(t, rbac.Grant(adminrole, createUser))
	if !rbac.Check(adminrole, createUser) {
		t.Fatalf("%s should have %s", adminrole.Name(), createUser.Name())
	}
	assert(t, rbac.Grant(adminrole, deleteUser))
	if !rbac.Check(adminrole, deleteUser) {
		t.Fatalf("%s should have %s", adminrole.Name(), deleteUser.Name())
	}
	if rbac.Check(adminrole, readUser) {
		t.Fatalf("%s should not have %s permit not yet assigned", adminrole.Name(), readUser.Name())
	}

}

func BenchmarkRbacCheck(b *testing.B) {
	rbac = New()
	rbac.AddRole(adminrole)
	rbac.AddRole(anonymouserole)
	rbac.Grant(adminrole, createUser)
	rbac.Grant(adminrole, readUser)

	rbac.Grant(anonymouserole, readUser)
	for i := 0; i < b.N; i++ {
		rbac.Check(adminrole, createUser)
	}

}

func BenchmarkRbacNotPermitted(b *testing.B) {
	rbac = New()
	rbac.AddRole(adminrole)
	rbac.AddRole(anonymouserole)
	rbac.Grant(adminrole, createUser)
	rbac.Grant(adminrole, readUser)

	rbac.Grant(anonymouserole, readUser)
	for i := 0; i < b.N; i++ {
		rbac.Check(anonymouserole, createUser)
	}
}
