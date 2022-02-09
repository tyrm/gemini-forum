package postgres

import (
	"github.com/tyrm/gemini-forum/db"
)

// Update a struct
func (c *Client) Update(obj interface{}) error {
	switch obj := obj.(type) {
	default:
		_ = obj
		return db.ErrUnknownType
	}
}
