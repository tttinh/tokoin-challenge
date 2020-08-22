package repo

import (
	"github.com/tttinh/tokoin-challenge/model"
	"github.com/tttinh/tokoin-challenge/utils"
	"encoding/json"
	"errors"
	"io/ioutil"
	"strconv"
)

// Organization manages all organizations.
type Organization struct {
	organizations model.Organizations
	orgByID       model.OrgByID
}

// Load loads json data from file.
func (repo *Organization) Load(filename string) error {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	err = json.Unmarshal(file, &repo.organizations)
	if err != nil {
		return err
	}

	repo.orgByID = make(model.OrgByID)
	for _, org := range repo.organizations {
		repo.orgByID[org.ID] = org
	}
	return nil
}

// Find look for organizations by key and value.
func (repo *Organization) Find(key, value string) (model.Organizations, error) {
	results := make(model.Organizations, 0)
	switch key {
	case "_id":
		id, err := strconv.Atoi(value)
		if err != nil {
			return nil, err
		}
		if org, found := repo.orgByID[id]; found {
			results = append(results, org)
		}
	case "url":
		for _, org := range repo.organizations {
			if org.URL == value {
				results = append(results, org)
			}
		}
	case "external_id":
		for _, org := range repo.organizations {
			if org.ExternalID == value {
				results = append(results, org)
			}
		}
	case "name":
		for _, org := range repo.organizations {
			if org.Name == value {
				results = append(results, org)
			}
		}
	case "domain_names":
		for _, org := range repo.organizations {
			for _, d := range org.DomainNames {
				if d == value {
					results = append(results, org)
					break
				}
			}
		}
	case "created_at":
		for _, org := range repo.organizations {
			if org.CreatedAt == value {
				results = append(results, org)
			}
		}
	case "details":
		for _, org := range repo.organizations {
			if org.Details == value {
				results = append(results, org)
			}
		}
	case "shared_tickets":
		v, err := utils.StringToBool(value)
		if err != nil {
			return nil, err
		}

		for _, org := range repo.organizations {
			if org.SharedTickets == v {
				results = append(results, org)
			}
		}
	case "tags":
		for _, org := range repo.organizations {
			for _, tag := range org.Tags {
				if tag == value {
					results = append(results, org)
					break
				}
			}
		}
	default:
		return nil, errors.New("unsupport search term")
	}

	return results, nil
}

// GetName returns organization name by id.
func (repo *Organization) GetName(id int) string {
	if org, found := repo.orgByID[id]; found {
		return org.Name
	}

	return ""
}
