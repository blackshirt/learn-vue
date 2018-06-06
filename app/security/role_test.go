package security

import (
	"testing"
)

func TestStdRoleName(t *testing.T) {
	if adminrole.Name() != "admin" {
		t.Fatalf("[admin] expected, but %s got", adminrole.Id)
	}

}

func TestStdRoleAssignPermit(t *testing.T) {
	if err := adminrole.Assign(createUser); err != nil {
		t.Fatal(err)
	}
}

func TestStdRoleRevokePermit(t *testing.T) {
	if err := adminrole.Revoke(createUser); err != nil {
		t.Fatal(err)
	}
	if err := adminrole.Revoke(deleteUser); err != nil {
		t.Fatal(err)
	}
	if err := adminrole.Revoke(pingUser); err != nil {
		t.Fatal(err)
	}
}

func TestStdRolePermit(t *testing.T) {
	if err := adminrole.Assign(createUser); err != nil {
		t.Fatal(err)
	}
	if !adminrole.Permit(createUser) {
		t.Fatalf("have %s", createUser.Name())
	}

}

func TestStdRoleLenPermission(t *testing.T) {
	if err := adminrole.Revoke(createUser); err != nil {
		t.Fatal(err)
	}
	if err := adminrole.Revoke(readUser); err != nil {
		t.Fatal(err)
	}
	if err := adminrole.Revoke(pingUser); err != nil {
		t.Fatal(err)
	}
	if len(adminrole.Permissions()) != 0 {
		t.Fatal("%s should not have any permission", adminrole.Name())
	}
}
