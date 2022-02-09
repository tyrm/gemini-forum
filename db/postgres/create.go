package postgres

import (
	"github.com/tyrm/gemini-forum/db"
)

// Create a struct
func (c *Client) Create(obj interface{}) error {
	switch obj := obj.(type) {
	default:
		_ = obj
		return db.ErrUnknownType
	}
}
