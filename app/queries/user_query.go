package queries

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/prabhatpankaj/go-fiber-rest-api/app/models"
)

// UserQueries struct for queries from User model.
type UserQueries struct {
	*sqlx.DB
}

// GetRoles method for getting all roles.
func (q *UserQueries) GetRoles() ([]models.Role, error) {
	// Define roles variable.
	roles := []models.Role{}

	// Define query string.
	query := `SELECT * FROM roles`

	// Send query to database.
	err := q.Select(&roles, query)
	if err != nil {
		// Return empty object and error.
		return roles, err
	}

	// Return query result.
	return roles, nil
}

// GetRole query for getting one Role by given ID.
func (q *UserQueries) GetRoleByName(name string) (models.Role, error) {
	// Define Role variable.
	role := models.Role{}

	// Define query string.// Define query string.
	query := `SELECT * FROM roles WHERE name = $1`

	// Send query to database.
	err := q.Get(&role, query, name)
	if err != nil {
		// Return empty object and error.
		return role, err
	}

	// Return query result.
	return role, nil
}

// GetRole query for getting one Role by given ID.
func (q *UserQueries) GetRoleByUser(id uuid.UUID) ([]models.Role, error) {
	// Define Role variable.
	roles := []models.Role{}
	// Define query string.// Define query string.
	query := `SELECT r.id, r.name FROM roles r INNER JOIN user_roles ur ON r.id = ur.role_id WHERE ur.user_id = $1`

	// Send query to database.
	err := q.Select(&roles, query, id)

	if err != nil {
		// Return empty object and error.
		return roles, err
	}

	// Return query result.
	return roles, nil
}

// GetUserByID query for getting one User by given ID.
func (q *UserQueries) GetUserByID(id uuid.UUID) (models.User, error) {
	// Define User variable.
	user := models.User{}

	// Define query string.
	query := `SELECT * FROM users WHERE id = $1`

	// Send query to database.
	err := q.Get(&user, query, id)
	if err != nil {
		// Return empty object and error.
		return user, err
	}

	// Return query result.
	return user, nil
}

// GetUserByEmail query for getting one User by given Email.
func (q *UserQueries) GetUserByEmail(email string) (models.User, error) {
	// Define User variable.
	user := models.User{}

	// Define query string.
	query := `SELECT * FROM users WHERE email = $1`

	// Send query to database.
	err := q.Get(&user, query, email)
	if err != nil {
		// Return empty object and error.
		return user, err
	}

	// Return query result.
	return user, nil
}

// CreateUser query for creating a new user by given email and password hash.
func (q *UserQueries) CreateUser(u *models.User) error {
	// Define query string.
	query := `INSERT INTO users VALUES ($1, $2, $3, $4, $5, $6, $7)`

	// Send query to database.
	_, err := q.Exec(
		query,
		u.ID, u.CreatedAt, u.UpdatedAt, u.Email, u.FullName, u.PasswordHash, u.UserStatus,
	)
	if err != nil {
		// Return only error.
		return err
	}

	// This query returns nothing.
	return nil
}

// CreateUserRole query for creating a new user role by given user and role id.
func (q *UserQueries) CreateUserRole(u *models.UserRole) error {
	// Define query string.
	query := `INSERT INTO user_roles VALUES ($1, $2)`

	// Send query to database.
	_, err := q.Exec(
		query,
		u.UserID, u.RoleID,
	)
	if err != nil {
		// Return only error.
		return err
	}

	// This query returns nothing.
	return nil
}
