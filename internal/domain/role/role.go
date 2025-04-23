package role

import "slices"

type Role string

const (
	RoleSuperAdmin Role = "SUPER_ADMIN"
	RoleAdmin      Role = "ADMIN"
	RoleUser       Role = "USER"
	RoleAnonymous  Role = "ANONYMOUS"
	UnknownRole    Role = "UNKNOWN"
)

// NewRole creates a valid issuerCode instance.
// If the input string does not match a known issuer name (not recognized), it returns UnknownIssuer.
func NewRole(roleName string) Role {
	switch Role(roleName) {
	case RoleSuperAdmin:
		return RoleSuperAdmin
	case RoleAdmin:
		return RoleAdmin
	case RoleUser:
		return RoleUser
	case RoleAnonymous:
		return RoleAnonymous
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
