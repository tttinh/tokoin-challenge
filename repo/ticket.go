package repo

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"strconv"

	"github.com/tttinh/tokoin-challenge/model"
	"github.com/tttinh/tokoin-challenge/utils"
)

// Ticket manages all tickets.
type Ticket struct {
	tickets            model.Tickets
	ticketByID         model.TicketByID
	ticketsByOrg       model.TicketsMap
	ticketsByAssignee  model.TicketsMap
	ticketsBySubmitter model.TicketsMap
}

// LoadFromFile loads data from file.
func (repo *Ticket) LoadFromFile(filename string) error {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	return repo.LoadFromByte(file)
}

// LoadFromByte loads data from []byte.
func (repo *Ticket) LoadFromByte(data []byte) error {
	err := json.Unmarshal(data, &repo.tickets)
	if err != nil {
		return err
	}

	repo.ticketByID = make(model.TicketByID)
	repo.ticketsByOrg = make(model.TicketsMap)
	repo.ticketsByAssignee = make(model.TicketsMap)
	repo.ticketsBySubmitter = make(model.TicketsMap)
	for _, tck := range repo.tickets {
		repo.ticketByID[tck.ID] = tck
		repo.ticketsByOrg[tck.OrganizationID] = append(repo.ticketsByOrg[tck.OrganizationID], tck)
		repo.ticketsByAssignee[tck.AssigneeID] = append(repo.ticketsByAssignee[tck.AssigneeID], tck)
		repo.ticketsBySubmitter[tck.SubmitterID] = append(repo.ticketsBySubmitter[tck.SubmitterID], tck)
	}

	return nil
}

// Find look for tickets by key and value.
func (repo *Ticket) Find(key, value string) (model.Tickets, error) {
	results := make(model.Tickets, 0)
	switch key {
	case "_id":
		for _, tck := range repo.tickets {
			if tck.ID == value {
				results = append(results, tck)
			}
		}
	case "url":
		for _, tck := range repo.tickets {
			if tck.URL == value {
				results = append(results, tck)
			}
		}
	case "external_id":
		for _, tck := range repo.tickets {
			if tck.ExternalID == value {
				results = append(results, tck)
			}
		}
	case "created_at":
		for _, tck := range repo.tickets {
			if tck.CreatedAt == value {
				results = append(results, tck)
			}
		}
	case "type":
		for _, tck := range repo.tickets {
			if tck.Type == value {
				results = append(results, tck)
			}
		}
	case "subject":
		for _, tck := range repo.tickets {
			if tck.Subject == value {
				results = append(results, tck)
			}
		}
	case "description":
		for _, tck := range repo.tickets {
			if tck.Description == value {
				results = append(results, tck)
			}
		}
	case "priority":
		for _, tck := range repo.tickets {
			if tck.Priority == value {
				results = append(results, tck)
			}
		}
	case "status":
		for _, tck := range repo.tickets {
			if tck.Status == value {
				results = append(results, tck)
			}
		}
	case "submitter_id":
		id, err := strconv.Atoi(value)
		if err != nil {
			return nil, err
		}

		for _, tck := range repo.tickets {
			if tck.SubmitterID == id {
				results = append(results, tck)
			}
		}
	case "assignee_id":
		id, err := strconv.Atoi(value)
		if err != nil {
			return nil, err
		}

		for _, tck := range repo.tickets {
			if tck.AssigneeID == id {
				results = append(results, tck)
			}
		}
	case "organization_id":
		id, err := strconv.Atoi(value)
		if err != nil {
			return nil, err
		}

		for _, tck := range repo.tickets {
			if tck.OrganizationID == id {
				results = append(results, tck)
			}
		}
	case "tags":
		for _, tck := range repo.tickets {
			for _, tag := range tck.Tags {
				if tag == value {
					results = append(results, tck)
					break
				}
			}
		}
	case "has_incidents":
		v, err := utils.StringToBool(value)
		if err != nil {
			return nil, err
		}

		for _, tck := range repo.tickets {
			if tck.HasIncidents == v {
				results = append(results, tck)
			}
		}
	case "due_at":
		for _, tck := range repo.tickets {
			if tck.DueAt == value {
				results = append(results, tck)
			}
		}
	case "via":
		for _, tck := range repo.tickets {
			if tck.Via == value {
				results = append(results, tck)
			}
		}
	default:
		return nil, errors.New("unsupport search term")
	}

	return results, nil
}

// GetUserTickets returns all ticket subjects related to a user.
func (repo *Ticket) GetUserTickets(id int) []string {
	var subjects []string
	if tickets, found := repo.ticketsByAssignee[id]; found {
		for _, tck := range tickets {
			subjects = append(subjects, tck.Subject)
		}
	}

	if tickets, found := repo.ticketsBySubmitter[id]; found {
		for _, tck := range tickets {
			subjects = append(subjects, tck.Subject)
		}
	}

	return subjects
}

// GetTicketsByOrg returns all tickets of an organization.
func (repo *Ticket) GetTicketsByOrg(id int) model.Tickets {
	if tickets, found := repo.ticketsByOrg[id]; found {
		return tickets
	}

	return nil
}
