package postgres

import (
	"database/sql"
	"github.com/google/uuid"
	"github.com/tyrm/gemini-forum/models"
)

// ReadUser will retrieve a user by their uuid from the database
func (c *Client) ReadUser(id uuid.UUID) (*models.User, error) {
	var user models.User
	err := c.db.
		Get(&user, `SELECT id, certhash, username, created_at, updated_at, deleted_at
		FROM public.users WHERE id = $1 AND deleted_at IS NULL;`, id)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	groups, err := c.readGroupsByUser(user.ID)
	if err != nil {
		return nil, err
	}
	user.Groups = groups

	return &user, nil
}

// ReadUserByCertHash will read a user by certhash from the database
func (c *Client) ReadUserByCertHash(certhash string) (*models.User, error) {
	var user models.User
	err := c.db.
		Get(&user, `SELECT id, certhash, username, created_at, updated_at, deleted_at
		FROM public.users WHERE certhash = $1 AND deleted_at IS NULL;`, certhash)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	groups, err := c.readGroupsByUser(user.ID)
	if err != nil {
		return nil, err
	}
	user.Groups = groups

	return &user, nil
}

// ReadUserByUsername will read a user by username from the database
func (c *Client) ReadUserByUsername(username string) (*models.User, error) {
	var user models.User
	err := c.db.
		Get(&user, `SELECT id, certhash, username, created_at, updated_at, deleted_at
		FROM public.users WHERE lower(username) = lower($1) AND deleted_at IS NULL;`, username)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	groups, err := c.readGroupsByUser(user.ID)
	if err != nil {
		return nil, err
	}
	user.Groups = groups

	return &user, nil
}

func (c *Client) readGroupsByUser(userID uuid.UUID) ([]uuid.UUID, error) {
	var groups []uuid.UUID
	err := c.db.
		Select(&groups, `SELECT group_id 
		FROM public.group_membership WHERE user_id = $1 AND deleted_at IS NULL;`, userID)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return groups, nil
}
