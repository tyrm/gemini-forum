package postgres

import (
	"github.com/tyrm/gemini-forum/db"
)

// Delete a struct
func (c *Client) Delete(obj interface{}) error {
	switch obj := obj.(type) {
	default:
		_ = obj
		return db.ErrUnknownType
	}
}
