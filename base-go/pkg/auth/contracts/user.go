package contracts

import "github.com/pkg/errors"

const (
	USER_DOESNT_HAVE_ROLE_ERROR = "user_doesnt_have_role"
)

type Role struct {
	Name        string   `json:"name"`
	Permissions []string `json:"permissions"`
	IsDefault   bool     `json:"is_default"`
}

type User struct {
	id             string
	activeRole     string
	roles          []Role
	hashedPassword string
}

func NewUser(id string) *User {
	return &User{
		id:         id,
		activeRole: "",
		roles:      []Role{},
	}
}

func (u *User) Id() string {
	return u.id
}

func (u *User) ActiveRole() string {
	return u.activeRole
}

func (u *User) Roles() []Role {
	return u.roles
}

func (u *User) AddRole(name string, permissions []string, isDefault bool) {
	u.roles = append(u.roles, Role{
		Name:        name,
		Permissions: permissions,
		IsDefault:   isDefault,
	})

	if isDefault && u.activeRole == "" {
		u.SetActiveRole(name)
	}
}

func (u *User) SetActiveRole(name string) error {
	for _, role := range u.roles {
		if role.Name == name {
			u.activeRole = name
			return nil
		}
	}

	return errors.New(USER_DOESNT_HAVE_ROLE_ERROR)
}

func (u *User) HasPermission(permission string) bool {
	for _, role := range u.roles {
		if role.Name == u.activeRole {
			for _, perm := range role.Permissions {
				if perm == permission {
					return true
				}
			}
		}
	}

	return false
}

func (u *User) HashedPassword() string {
	return u.hashedPassword
}

func (u *User) SetHashedPassword(hashedPassword string) {
	u.hashedPassword = hashedPassword
}
