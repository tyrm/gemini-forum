package db

// DB represents the required commands for a database.
type DB interface {
	Create(interface{}) error
	Delete(interface{}) error
	Update(interface{}) error
}
