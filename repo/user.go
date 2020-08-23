package repo

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"strconv"

	"github.com/tttinh/tokoin-challenge/model"
	"github.com/tttinh/tokoin-challenge/utils"
)

// User manages all users and underline data structures for searching.
type User struct {
	users      model.Users
	userByID   model.UserByID
	usersByOrg model.UsersByOrg
}

// LoadFromFile loads data from file.
func (repo *User) LoadFromFile(filename string) error {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	return repo.LoadFromByte(file)
}

// LoadFromByte loads data from []byte.
func (repo *User) LoadFromByte(data []byte) error {
	err := json.Unmarshal(data, &repo.users)
	if err != nil {
		return err
	}

	repo.userByID = make(model.UserByID)
	repo.usersByOrg = make(model.UsersByOrg)
	for _, usr := range repo.users {
		repo.userByID[usr.ID] = usr
		repo.usersByOrg[usr.OrganizationID] = append(repo.usersByOrg[usr.OrganizationID], usr)
	}

	return nil
}

// Find look for users by key and value.
func (repo *User) Find(key, value string) (model.Users, error) {
	results := make(model.Users, 0)
	switch key {
	case "_id":
		id, err := strconv.Atoi(value)
		if err != nil {
			return nil, err
		}
		if usr, found := repo.userByID[id]; found {
			results = append(results, usr)
		}
	case "url":
		for _, usr := range repo.users {
			if usr.URL == value {
				results = append(results, usr)
			}
		}
	case "external_id":
		for _, usr := range repo.users {
			if usr.ExternalID == value {
				results = append(results, usr)
			}
		}
	case "name":
		for _, usr := range repo.users {
			if usr.Name == value {
				results = append(results, usr)
			}
		}
	case "alias":
		for _, usr := range repo.users {
			if usr.Alias == value {
				results = append(results, usr)
			}
		}
	case "created_at":
		for _, usr := range repo.users {
			if usr.CreatedAt == value {
				results = append(results, usr)
			}
		}
	case "active":
		v, err := utils.StringToBool(value)
		if err != nil {
			return nil, err
		}

		for _, usr := range repo.users {
			if usr.Active == v {
				results = append(results, usr)
			}
		}
	case "verified":
		v, err := utils.StringToBool(value)
		if err != nil {
			return nil, err
		}

		for _, usr := range repo.users {
			if usr.Verified == v {
				results = append(results, usr)
			}
		}
	case "shared":
		v, err := utils.StringToBool(value)
		if err != nil {
			return nil, err
		}

		for _, usr := range repo.users {
			if usr.Shared == v {
				results = append(results, usr)
			}
		}
	case "locale":
		for _, usr := range repo.users {
			if usr.Locale == value {
				results = append(results, usr)
			}
		}
	case "timezone":
		for _, usr := range repo.users {
			if usr.Timezone == value {
				results = append(results, usr)
			}
		}
	case "last_login_at":
		for _, usr := range repo.users {
			if usr.LastLoginAt == value {
				results = append(results, usr)
			}
		}
	case "email":
		for _, usr := range repo.users {
			if usr.Email == value {
				results = append(results, usr)
			}
		}
	case "phone":
		for _, usr := range repo.users {
			if usr.Phone == value {
				results = append(results, usr)
			}
		}
	case "signature":
		for _, usr := range repo.users {
			if usr.Signature == value {
				results = append(results, usr)
			}
		}
	case "organization_id":
		id, err := strconv.Atoi(value)
		if err != nil {
			return nil, err
		}
		for _, usr := range repo.users {
			if usr.OrganizationID == id {
				results = append(results, usr)
			}
		}
	case "tags":
		for _, usr := range repo.users {
			for _, tag := range usr.Tags {
				if tag == value {
					results = append(results, usr)
					break
				}
			}
		}
	case "suspended":
		v, err := utils.StringToBool(value)
		if err != nil {
			return nil, err
		}

		for _, usr := range repo.users {
			if usr.Suspended == v {
				results = append(results, usr)
			}
		}
	case "role":
		for _, usr := range repo.users {
			if usr.Role == value {
				results = append(results, usr)
			}
		}
	default:
		return nil, errors.New("unsupport search term")
	}

	return results, nil
}

// GetName returns user name by id.
func (repo *User) GetName(id int) string {
	if usr, found := repo.userByID[id]; found {
		return usr.Name
	}

	return ""
}

// GetUsersByOrg returns all users of an organization.
func (repo *User) GetUsersByOrg(id int) model.Users {
	if users, found := repo.usersByOrg[id]; found {
		return users
	}

	return nil
}
