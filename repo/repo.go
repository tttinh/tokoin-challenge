package repo

import (
	"github.com/tttinh/tokoin-challenge/config"
)

// Results represents a collection of data records.
type Results interface {
	Print()
	Length() int
}

// Repository interface.
type Repository interface {
	Init(cfg *config.Configuration) error
	FindUsers(key, value string) (Results, error)
	FindTickets(key, value string) (Results, error)
	FindOrganizations(key, value string) (Results, error)
}

// An implementation of Repository interface.
type repositoryImpl struct {
	users         *User
	tickets       *Ticket
	organizations *Organization
}

// New returns an instance of Repository interface.
func New() Repository {
	return &repositoryImpl{}
}

func (rp *repositoryImpl) Init(cfg *config.Configuration) error {
	rp.users = &User{}
	if err := rp.users.Load(cfg.UserData); err != nil {
		return err
	}

	rp.tickets = &Ticket{}
	if err := rp.tickets.Load(cfg.TicketData); err != nil {
		return err
	}

	rp.organizations = &Organization{}
	if err := rp.organizations.Load(cfg.OrganizationData); err != nil {
		return err
	}

	return nil
}

// FindUser searchs for users based on key and value.
func (rp *repositoryImpl) FindUsers(key, value string) (Results, error) {
	results, err := rp.users.Find(key, value)
	if err != nil {
		return nil, err
	}

	for _, usr := range results {
		usr.OrganizationName = rp.organizations.GetName(usr.OrganizationID)
		usr.Tickets = rp.tickets.GetUserTickets(usr.ID)
	}

	return results, nil
}

// FindTickets searchs for tickets based on key and value.
func (rp *repositoryImpl) FindTickets(key, value string) (Results, error) {
	results, err := rp.tickets.Find(key, value)
	if err != nil {
		return nil, err
	}

	for _, tck := range results {
		tck.OrganizationName = rp.organizations.GetName(tck.OrganizationID)
		tck.SubmitterName = rp.users.GetName(tck.SubmitterID)
		tck.AssigneeName = rp.users.GetName(tck.AssigneeID)
	}

	return results, nil
}

// FindOrganizations searchs for organization based on key and value.
func (rp *repositoryImpl) FindOrganizations(key, value string) (Results, error) {
	results, err := rp.organizations.Find(key, value)
	if err != nil {
		return nil, err
	}

	for _, org := range results {
		for _, usr := range rp.users.GetUsersByOrg(org.ID) {
			org.Users = append(org.Users, usr.Name)
		}

		for _, tck := range rp.tickets.GetTicketsByOrg(org.ID) {
			org.Tickets = append(org.Tickets, tck.Subject)
		}
	}

	return results, nil
}
