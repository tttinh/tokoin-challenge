package model

import "github.com/tttinh/tokoin-challenge/utils"

// Organization model
type Organization struct {
	ID            int      `json:"_id"`
	URL           string   `json:"url"`
	ExternalID    string   `json:"external_id"`
	Name          string   `json:"name"`
	DomainNames   []string `json:"domain_names"`
	CreatedAt     string   `json:"created_at"`
	Details       string   `json:"details"`
	SharedTickets bool     `json:"shared_tickets"`
	Tags          []string `json:"tags"`
	Users         []string
	Tickets       []string
}

// Organizations represents an array of organizations.
type Organizations []*Organization

// OrgByID represents map from ID to Organization.
type OrgByID map[int]*Organization

// Print prints all items.
func (organizations Organizations) Print() {
	for _, org := range organizations {
		utils.PrintObject(org)
		utils.Writef("  %v\n", "users:")
		for _, usr := range org.Users {
			utils.Writef("    %v\n", usr)
		}

		utils.Writef("  %v\n", "tickets:")
		for _, tck := range org.Tickets {
			utils.Writef("    %v\n", tck)
		}
	}
}

// Length returns the number of items.
func (organizations Organizations) Length() int {
	return len(organizations)
}
