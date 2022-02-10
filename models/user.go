package models

import (
	"database/sql"
	"github.com/google/uuid"
	"github.com/tyrm/gemini-forum/util"
	"time"
)

// User is used to login and keep authentication information
type User struct {
	CertHash string `db:"certhash"`
	Username string `db:"username"`

	Groups []uuid.UUID

	ID        uuid.UUID    `db:"id"`
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt time.Time    `db:"updated_at"`
	DeletedAt sql.NullTime `db:"deleted_at"`
}

// IsMemberOfGroup checks if a user is in a given set of groups
func (u *User) IsMemberOfGroup(groups ...uuid.UUID) bool {
	return util.ContainsOneOfUUIDs(&u.Groups, &groups)
}
