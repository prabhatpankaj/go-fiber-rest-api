package utils

import (
	"fmt"

	"github.com/prabhatpankaj/go-fiber-rest-api/pkg/repository"
)

// VerifyRole func for verifying a given role.
func VerifyRole(role string) (string, error) {
	// Switch given role.
	switch role {
	case repository.AdminRoleName:
		// Nothing to do, verified successfully.
	case repository.UserRoleName:
		// Nothing to do, verified successfully.
	default:
		// Return error message.
		return "", fmt.Errorf("role '%v' does not exist", role)
	}

	return role, nil
}

func hasPermission(userRoles []string, requiredPermission string) bool {
	// Check if the user's roles contain the required permission
	for _, role := range userRoles {
		if role == requiredPermission {
			return true
		}
	}
	return false
}
