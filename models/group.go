package models

import "github.com/google/uuid"

// GroupSuperAdmin is the uuid of the Super Administrators group
var GroupSuperAdmin = uuid.Must(uuid.Parse("5cf41f36-a0ae-474c-ae26-73819e38c64d"))

// GroupTitle contains the titles of the groups.
var groupTitle = map[uuid.UUID]string{
	GroupSuperAdmin: "Super Admin",
}

// groups of groups

// GroupsAll contains the uuids of all groups
var GroupsAll = []uuid.UUID{
	GroupSuperAdmin,
}

// GroupsAllAdmins contains the uuids of all admin groups
var GroupsAllAdmins = []uuid.UUID{
	GroupSuperAdmin,
}

// GroupTitle return a pretty text name for the group
func GroupTitle(g uuid.UUID) string {
	if s, ok := groupTitle[g]; ok {
		return s
	}
	return ""
}
