package lib

import (
	"fmt"
	"testing"
)

type MockRoleGetter struct {
	r Roles
}

func NewMockRoleGetter(r []string) RoleGetter {
	return &MockRoleGetter{r: Roles(r)}
}

func (m *MockRoleGetter) Roles() Roles {
	return m.r.Dedup()
}

func ExampleRoleGetter() {
	roles := []string{
		"mock3", "mock2", "mock1", "mock2", "mock4", "mock1",
	}
	m := NewMockRoleGetter(roles)
	for _, r := range m.Roles() {
		fmt.Println(r)
	}
	// Output:
	// mock1
	// mock2
	// mock3
	// mock4
}

func TestEmptyRoleGetter(t *testing.T) {
	m := NewMockRoleGetter([]string{})
	r := m.Roles()

	t.Logf("Empty role result: %v", r)
	if len(r) != 0 {
		t.Errorf("Found unexpected roles from empty input")
	}
}

func TestNilRoleGetter(t *testing.T) {
	m := NewMockRoleGetter(nil)
	r := m.Roles()

	t.Logf("Nil role result: %v", r)
	if len(r) != 0 {
		t.Errorf("Found unexpected roles from nil input")
	}
}
