package db

import (
	"github.com/google/uuid"
	"github.com/tyrm/gemini-forum/models"
)

// DB represents the required commands for a database.
type DB interface {
	Create(interface{}) error
	Delete(interface{}) error
	ReadUser(uuid.UUID) (*models.User, error)
	ReadUserByCertHash(string) (*models.User, error)
	ReadUserByUsername(string) (*models.User, error)
	Update(interface{}) error
}
