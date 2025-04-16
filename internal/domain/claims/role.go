package claims

type Role string

const (
	RoleSuperAdmin Role = "ROLE_SUPER_ADMIN"
	RoleAdmin      Role = "ROLE_ADMIN"
	RoleUser       Role = "ROLE_USER"
	RoleAnonymous  Role = "ROLE_ANONYMOUS"
	UnknownRole    Role = "ROLE_UNKNOWN"
)

func ConvertRole(roleName string) Role {
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

func IsGivenRolePresent(expectedRole Role, presentRoles []Role) bool {
	for _, role := range presentRoles {
		if role == expectedRole {
			return true
		}
	}
	return false
}

// IsGivenRolesPresent has n * n algorithm complexity that is accepted.
func IsGivenRolesPresent(expectedRoles []Role, presentRoles []Role) bool {
	var matchCount int
	for _, neededRole := range expectedRoles {
		for _, existingRole := range presentRoles {
			if existingRole == neededRole {
				matchCount++
				break
			}
		}
	}
	return matchCount == len(expectedRoles)
}
