package role

import "slices"

type Role string

const (
	SuperAdmin  Role = "SUPER_ADMIN"
	Admin       Role = "ADMIN"
	User        Role = "USER"
	Anonymous   Role = "ANONYMOUS"
	UnknownRole Role = "UNKNOWN"
)

// NewRole creates a valid issuerCode instance.
// If the input string does not match a known issuer name (not recognized), it returns UnknownIssuer.
func NewRole(roleName string) Role {
	switch Role(roleName) {
	case SuperAdmin:
		return SuperAdmin
	case Admin:
		return Admin
	case User:
		return User
	case Anonymous:
		return Anonymous
	}
	return UnknownRole
}

// IsGivenRolePresent returns true if expectedRole is found in presentRoles.
func IsGivenRolePresent(expectedRole Role, presentRoles []Role) bool {
	return slices.Contains(presentRoles, expectedRole)
}

// IsGivenRolesPresent returns true if all expectedRoles are found in presentRoles.
// Note: It has O(n * m) complexity, where n = len(expectedRoles), m = len(presentRoles).
func IsGivenRolesPresent(expectedRoles []Role, presentRoles []Role) bool {
	for _, role := range expectedRoles {
		if !slices.Contains(presentRoles, role) {
			return false
		}
	}
	return true
}
