package repo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const corruptedUserData = `user`
const sampleUserData = `
[
  {
    "_id": 1,
    "url": "http://initech.tokoin.io.com/api/v2/users/1.json",
    "external_id": "74341f74-9c79-49d5-9611-87ef9b6eb75f",
    "name": "Francisca Rasmussen",
    "alias": "Miss Coffey",
    "created_at": "2016-04-15T05:19:46 -10:00",
    "active": true,
    "verified": true,
    "shared": false,
    "locale": "en-AU",
    "timezone": "Sri Lanka",
    "last_login_at": "2013-08-04T01:03:27 -10:00",
    "email": "coffeyrasmussen@flotonic.com",
    "phone": "8335-422-718",
    "signature": "Don't Worry Be Happy!",
    "organization_id": 119,
    "tags": [
      "Springville",
      "Sutton",
      "Hartsville/Hartley",
      "Diaperville"
    ],
    "suspended": true,
    "role": "admin"
  },
  {
    "_id": 2,
    "url": "http://initech.tokoin.io.com/api/v2/users/2.json",
    "external_id": "c9995ea4-ff72-46e0-ab77-dfe0ae1ef6c2",
    "name": "Cross Barlow",
    "alias": "Miss Joni",
    "created_at": "2016-06-23T10:31:39 -10:00",
    "active": true,
    "verified": true,
    "shared": false,
    "locale": "zh-CN",
    "timezone": "Armenia",
    "last_login_at": "2012-04-12T04:03:28 -10:00",
    "email": "jonibarlow@flotonic.com",
    "phone": "9575-552-585",
    "signature": "Don't Worry Be Happy!",
    "organization_id": 106,
    "tags": [
      "Foxworth",
      "Woodlands",
      "Herlong",
      "Henrietta"
    ],
    "suspended": false,
    "role": "admin"
  }
]
`

func TestUser_LoadFromByte(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"Should pass", args{[]byte(sampleUserData)}, false},
		{"Should fail", args{[]byte(corruptedUserData)}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &User{}
			err := repo.LoadFromByte(tt.args.data)
			assert.Equal(t, tt.wantErr, (err != nil), err)
		})
	}
}

func TestUser_Find(t *testing.T) {
	mockRepo := &User{}
	assert.Nil(t, mockRepo.LoadFromByte([]byte(sampleUserData)))

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
		{"Find by existed _id", args{"_id", "1"}, 1, false},
		{"Find by existed url", args{"url", "http://initech.tokoin.io.com/api/v2/users/1.json"}, 1, false},
		{"Find by existed external_id", args{"external_id", "74341f74-9c79-49d5-9611-87ef9b6eb75f"}, 1, false},
		{"Find by existed name", args{"name", "Francisca Rasmussen"}, 1, false},
		{"Find by existed alias", args{"alias", "Miss Coffey"}, 1, false},
		{"Find by existed created_at", args{"created_at", "2016-04-15T05:19:46 -10:00"}, 1, false},
		{"Find by existed active", args{"active", "true"}, 2, false},
		{"Find by existed verified", args{"verified", "true"}, 2, false},
		{"Find by existed shared", args{"shared", "false"}, 2, false},
		{"Find by existed locale", args{"locale", "en-AU"}, 1, false},
		{"Find by existed timezone", args{"timezone", "Sri Lanka"}, 1, false},
		{"Find by existed last_login_at", args{"last_login_at", "2013-08-04T01:03:27 -10:00"}, 1, false},
		{"Find by existed email", args{"email", "coffeyrasmussen@flotonic.com"}, 1, false},
		{"Find by existed phone", args{"phone", "8335-422-718"}, 1, false},
		{"Find by existed signature", args{"signature", "Don't Worry Be Happy!"}, 2, false},
		{"Find by existed organization_id", args{"organization_id", "119"}, 1, false},
		{"Find by existed tags", args{"tags", "Springville"}, 1, false},
		{"Find by existed suspended", args{"suspended", "true"}, 1, false},
		{"Find by existed role", args{"role", "admin"}, 2, false},
		// TODO: Add not found test cases.
		// TODO: Add error test cases.
		{"Find by error _id", args{"_id", "fasdf"}, 0, true},
		{"Find by error active", args{"active", "fasdf"}, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := mockRepo.Find(tt.args.key, tt.args.value)
			assert.Equal(t, tt.wantErr, (err != nil), err)
			assert.Equal(t, tt.wantResult, len(got))
		})
	}
}

func TestUser_GetName(t *testing.T) {
	mockRepo := &User{}
	assert.Nil(t, mockRepo.LoadFromByte([]byte(sampleUserData)))

	type args struct {
		id int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Get name of an existing user", args{1}, "Francisca Rasmussen"},
		{"Get name of an unexisting user", args{3}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := mockRepo.GetName(tt.args.id)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestUser_GetUsersByOrg(t *testing.T) {
	mockRepo := &User{}
	assert.Nil(t, mockRepo.LoadFromByte([]byte(sampleUserData)))

	type args struct {
		id int
	}
	tests := []struct {
		name       string
		args       args
		wantResult int
	}{
		{"Get users of an existing org", args{119}, 1},
		{"Get users of an unexisting org", args{1}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := mockRepo.GetUsersByOrg(tt.args.id)
			assert.Equal(t, tt.wantResult, len(got))
		})
	}
}
