package repo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const corruptedOrgData = `org`
const sampleOrgData = `
[
  {
    "_id": 101,
    "url": "http://initech.tokoin.io.com/api/v2/organizations/101.json",
    "external_id": "9270ed79-35eb-4a38-a46f-35725197ea8d",
    "name": "Enthaze",
    "domain_names": [
      "kage.com",
      "ecratic.com",
      "endipin.com",
      "zentix.com"
    ],
    "created_at": "2016-05-21T11:10:28 -10:00",
    "details": "MegaCorp",
    "shared_tickets": false,
    "tags": [
      "Fulton",
      "West",
      "Rodriguez",
      "Farley"
    ]
  },
  {
    "_id": 119,
    "url": "http://initech.tokoin.io.com/api/v2/organizations/102.json",
    "external_id": "7cd6b8d4-2999-4ff2-8cfd-44d05b449226",
    "name": "Nutralab",
    "domain_names": [
      "trollery.com",
      "datagen.com",
      "bluegrain.com",
      "dadabase.com"
    ],
    "created_at": "2016-04-07T08:21:44 -10:00",
    "details": "Non profit",
    "shared_tickets": false,
    "tags": [
      "Cherry",
      "Collier",
      "Fuentes",
      "Trevino"
    ]
  }
]
`

func TestOrganization_LoadFromByte(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"Should pass", args{[]byte(sampleOrgData)}, false},
		{"Should fail", args{[]byte(corruptedOrgData)}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &Organization{}
			err := repo.LoadFromByte(tt.args.data)
			assert.Equal(t, tt.wantErr, (err != nil), err)
		})
	}
}

func TestOrganization_Find(t *testing.T) {
	mockRepo := &Organization{}
	assert.Nil(t, mockRepo.LoadFromByte([]byte(sampleOrgData)))

	type args struct {
		key   string
		value string
	}
	tests := []struct {
		name       string
		args       args
		wantResult int
		wantErr    bool
	}{
		// found test cases.
		{"Find by existed _id", args{"_id", "119"}, 1, false},
		{"Find by existed url", args{"url", "http://initech.tokoin.io.com/api/v2/organizations/102.json"}, 1, false},
		{"Find by existed external_id", args{"external_id", "7cd6b8d4-2999-4ff2-8cfd-44d05b449226"}, 1, false},
		{"Find by existed name", args{"name", "Nutralab"}, 1, false},
		{"Find by existed domain_names", args{"domain_names", "datagen.com"}, 1, false},
		{"Find by existed created_at", args{"created_at", "2016-04-07T08:21:44 -10:00"}, 1, false},
		{"Find by existed details", args{"details", "Non profit"}, 1, false},
		{"Find by existed shared_tickets", args{"shared_tickets", "false"}, 2, false},
		{"Find by existed tags", args{"tags", "Collier"}, 1, false},
		// TODO: Add not found test cases.
		// TODO: Add error test cases.
		{"Find by error shared_tickets", args{"shared_tickets", "fasdf"}, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := mockRepo.Find(tt.args.key, tt.args.value)
			assert.Equal(t, tt.wantErr, (err != nil), err)
			assert.Equal(t, tt.wantResult, len(got))
		})
	}
}

func TestOrganization_GetName(t *testing.T) {
	mockRepo := &Organization{}
	assert.Nil(t, mockRepo.LoadFromByte([]byte(sampleOrgData)))

	type args struct {
		id int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Get name of an existing org", args{119}, "Nutralab"},
		{"Get name of an unexisting org", args{3}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := mockRepo.GetName(tt.args.id)
			assert.Equal(t, tt.want, got)
		})
	}
}
