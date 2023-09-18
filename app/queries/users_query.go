package queries

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/prabhatpankaj/go-fiber-rest-api/app/models"
)

// UserQueries struct for queries from Users model.
type UserQueries struct {
	*sqlx.DB
}

// GetUserByID method for getting one user by given ID.
func (q *UserQueries) GetUserByID(id uuid.UUID, b *models.Users) (models.Users, error) {
	// Define user variable.
	user := models.Users{}

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

// GetUserByUsername method for getting one user by given username.
func (q *UserQueries) GetUserByUsername(username string) (models.Users, error) {
	// Define user variable.
	user := models.Users{}

	// Define query string.
	query := `SELECT * FROM users WHERE username = $1`

	// Send query to database.
	err := q.Get(&user, query, username)
	if err != nil {
		// Return empty object and error.
		return user, err
	}

	// Return query result.
	return user, nil
}

// GetProfileByEmail method for getting one profile by given email.
func (q *UserQueries) GetProflebyEmail(email string) (models.Profile, error) {
	// Define user variable.
	profile := models.Profile{}

	// Define query string.
	query := `SELECT * FROM profile WHERE email = $1`

	// Send query to database.
	err := q.Get(&profile, query, email)
	if err != nil {
		// Return empty object and error.
		return profile, err
	}

	// Return query result.
	return profile, nil
}

// CreateUser method for creating user by given Users object.
func (q *UserQueries) CreateUser(b *models.Users) error {
	// Define query string.
	query := `INSERT INTO users VALUES ($1, $2, $3, $4, $5, $6)`
	// Send query to database.
	_, err := q.Exec(query, b.ID, b.Username, b.Password, b.ActiveStatus, b.CreatedAt, b.UpdatedAt)
	if err != nil {
		// Return only error.
		return err
	}

	// This query returns nothing.
	return nil
}
