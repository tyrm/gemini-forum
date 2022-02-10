package postgres

import (
	"github.com/tyrm/gemini-forum/db"
	"github.com/tyrm/gemini-forum/models"
)

// Create a struct
func (c *Client) Create(obj interface{}) error {
	switch obj := obj.(type) {
	case *models.User:
		return c.createUser(obj)
	default:
		return db.ErrUnknownType
	}
}

func (c *Client) createUser(u *models.User) error {
	// add to database
	return c.db.
		QueryRowx(`INSERT INTO "public"."users"("certhash", "username")
			VALUES ($1, $2) RETURNING id, created_at, updated_at;`, u.CertHash, u.Username).
		Scan(&u.ID, &u.CreatedAt, &u.UpdatedAt)
}
