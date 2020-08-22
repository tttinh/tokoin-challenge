package model

import "github.com/tttinh/tokoin-challenge/utils"

// Ticket model
type Ticket struct {
	ID               string   `json:"_id"`
	URL              string   `json:"url"`
	ExternalID       string   `json:"external_id"`
	CreatedAt        string   `json:"created_at"`
	Type             string   `json:"type"`
	Subject          string   `json:"subject"`
	Description      string   `json:"description"`
	Priority         string   `json:"priority"`
	Status           string   `json:"status"`
	SubmitterID      int      `json:"submitter_id"`
	AssigneeID       int      `json:"assignee_id"`
	OrganizationID   int      `json:"organization_id"`
	Tags             []string `json:"tags"`
	HasIncidents     bool     `json:"has_incidents"`
	DueAt            string   `json:"due_at"`
	Via              string   `json:"via"`
	OrganizationName string
	SubmitterName    string
	AssigneeName     string
}

// Tickets represents an array of tickets.
type Tickets []*Ticket

// TicketByID represents a map from ID to Ticket.
type TicketByID map[string]*Ticket

// TicketsMap represents a map from an int to Tickets.
type TicketsMap map[int]Tickets

// Print prints all items.
func (tickets Tickets) Print() {
	for _, tck := range tickets {
		utils.PrintObject(tck)
		utils.Writef("  %-30v%v\n", "organization_name", tck.OrganizationName)
		utils.Writef("  %-30v%v\n", "submitter_name", tck.SubmitterName)
		utils.Writef("  %-30v%v\n", "assignee_name", tck.AssigneeName)
	}
}

// Length returns the number of items.
func (tickets Tickets) Length() int {
	return len(tickets)
}
