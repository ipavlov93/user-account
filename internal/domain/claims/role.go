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
	if roleName == "ROLE_SUPER_ADMIN" {
		return RoleSuperAdmin
	} else if roleName == "ROLE_ADMIN" {
		return RoleAdmin
	} else if roleName == "ROLE_USER" {
		return RoleUser
	} else if roleName == "ROLE_ANONYMOUS" {
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
