package model

import (
	"github.com/tttinh/tokoin-challenge/utils"
	"fmt"
)

// User model
type User struct {
	ID               int      `json:"_id"`
	URL              string   `json:"url"`
	ExternalID       string   `json:"external_id"`
	Name             string   `json:"name"`
	Alias            string   `json:"alias"`
	CreatedAt        string   `json:"created_at"`
	Active           bool     `json:"active"`
	Verified         bool     `json:"verified"`
	Shared           bool     `json:"shared"`
	Locale           string   `json:"locale"`
	Timezone         string   `json:"timezone"`
	LastLoginAt      string   `json:"last_login_at"`
	Email            string   `json:"email"`
	Phone            string   `json:"phone"`
	Signature        string   `json:"signature"`
	OrganizationID   int      `json:"organization_id"`
	Tags             []string `json:"tags"`
	Suspended        bool     `json:"suspended"`
	Role             string   `json:"role"`
	OrganizationName string
	Tickets          []string
}

// Users represents an array of users.
type Users []*User

// UserByID represents map from ID to User.
type UserByID map[int]*User

// UsersByOrg represents a map from OrganizationID to Users.
type UsersByOrg map[int]Users

// Print prints all items.
func (users Users) Print() {
	for _, usr := range users {
		utils.PrintObject(usr)
		utils.Writef("  %-30v%v\n", "organization_name", usr.OrganizationName)
		for i, tck := range usr.Tickets {
			tckNumber := fmt.Sprintf("ticket_%v", i)
			utils.Writef("  %-30v%v\n", tckNumber, tck)
		}
	}
}

// Length returns the number of items.
func (users Users) Length() int {
	return len(users)
}
