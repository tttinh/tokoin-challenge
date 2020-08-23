package repo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const corruptedTicketData = `ticket`
const sampleTicketData = `
[
  {
    "_id": "436bf9b0-1147-4c0a-8439-6f79833bff5b",
    "url": "http://initech.tokoin.io.com/api/v2/tickets/436bf9b0-1147-4c0a-8439-6f79833bff5b.json",
    "external_id": "9210cdc9-4bee-485f-a078-35396cd74063",
    "created_at": "2016-04-28T11:19:34 -10:00",
    "type": "incident",
    "subject": "A Catastrophe in Korea (North)",
    "description": "Nostrud ad sit.",
    "priority": "high",
    "status": "pending",
    "submitter_id": 1,
    "assignee_id": 2,
    "organization_id": 119,
    "tags": [
      "Ohio",
      "Pennsylvania",
      "American Samoa",
      "Northern Mariana Islands"
    ],
    "has_incidents": false,
    "due_at": "2016-07-31T02:37:50 -10:00",
    "via": "web"
  },
  {
    "_id": "1a227508-9f39-427c-8f57-1b72f3fab87c",
    "url": "http://initech.tokoin.io.com/api/v2/tickets/1a227508-9f39-427c-8f57-1b72f3fab87c.json",
    "external_id": "3e5ca820-cd1f-4a02-a18f-11b18e7bb49a",
    "created_at": "2016-04-14T08:32:31 -10:00",
    "type": "incident",
    "subject": "A Catastrophe in Micronesia",
    "description": "Aliquip excepteur fugiat ex minim ea aute eu labore. Sunt eiusmod esse eu non commodo est veniam consequat.",
    "priority": "low",
    "status": "hold",
    "submitter_id": 2,
    "assignee_id": 1,
    "organization_id": 119,
    "tags": [
      "Puerto Rico",
      "Idaho",
      "Oklahoma",
      "Louisiana"
    ],
    "has_incidents": false,
    "due_at": "2016-08-15T05:37:32 -10:00",
    "via": "chat"
  },
  {
    "_id": "2217c7dc-7371-4401-8738-0a8a8aedc08d",
    "url": "http://initech.tokoin.io.com/api/v2/tickets/2217c7dc-7371-4401-8738-0a8a8aedc08d.json",
    "external_id": "3db2c1e6-559d-4015-b7a4-6248464a6bf0",
    "created_at": "2016-07-16T12:05:12 -10:00",
    "type": "problem",
    "subject": "A Catastrophe in Hungary",
    "description": "Ipsum fugiat voluptate reprehenderit cupidatat aliqua dolore consequat. Consequat ullamco minim laboris veniam ea id laborum et eiusmod excepteur sint laborum dolore qui.",
    "priority": "normal",
    "status": "closed",
    "submitter_id": 1,
    "assignee_id": 2,
    "organization_id": 105,
    "tags": [
      "Massachusetts",
      "New York",
      "Minnesota",
      "New Jersey"
    ],
    "has_incidents": true,
    "due_at": "2016-08-06T04:16:06 -10:00",
    "via": "web"
  }
]
  `

func TestTicket_LoadFromByte(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"Should pass", args{[]byte(sampleTicketData)}, false},
		{"Should fail", args{[]byte(corruptedTicketData)}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &Ticket{}
			err := repo.LoadFromByte(tt.args.data)
			assert.Equal(t, tt.wantErr, (err != nil), err)
		})
	}
}

func TestTicket_Find(t *testing.T) {
	mockRepo := &Ticket{}
	assert.Nil(t, mockRepo.LoadFromByte([]byte(sampleTicketData)))

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
		{"Find by existed _id", args{"_id", "436bf9b0-1147-4c0a-8439-6f79833bff5b"}, 1, false},
		{"Find by existed url", args{"url", "http://initech.tokoin.io.com/api/v2/tickets/436bf9b0-1147-4c0a-8439-6f79833bff5b.json"}, 1, false},
		{"Find by existed external_id", args{"external_id", "9210cdc9-4bee-485f-a078-35396cd74063"}, 1, false},
		{"Find by existed created_at", args{"created_at", "2016-04-28T11:19:34 -10:00"}, 1, false},
		{"Find by existed type", args{"type", "incident"}, 2, false},
		{"Find by existed subject", args{"subject", "A Catastrophe in Korea (North)"}, 1, false},
		{"Find by existed description", args{"description", "Nostrud ad sit."}, 1, false},
		{"Find by existed priority", args{"priority", "high"}, 1, false},
		{"Find by existed status", args{"status", "pending"}, 1, false},
		{"Find by existed submitter_id", args{"submitter_id", "1"}, 2, false},
		{"Find by existed assignee_id", args{"assignee_id", "2"}, 2, false},
		{"Find by existed organization_id", args{"organization_id", "119"}, 2, false},
		{"Find by existed tags", args{"tags", "Pennsylvania"}, 1, false},
		{"Find by existed has_incidents", args{"has_incidents", "false"}, 2, false},
		{"Find by existed due_at", args{"due_at", "2016-07-31T02:37:50 -10:00"}, 1, false},
		{"Find by existed via", args{"via", "web"}, 2, false},
		// TODO: Add not found test cases.
		// TODO: Add error test cases.
		{"Find by error has_incidents", args{"has_incidents", "fasdf"}, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := mockRepo.Find(tt.args.key, tt.args.value)
			assert.Equal(t, tt.wantErr, (err != nil), err)
			assert.Equal(t, tt.wantResult, len(got))
		})
	}
}

func TestTicket_GetUserTickets(t *testing.T) {
	mockRepo := &Ticket{}
	assert.Nil(t, mockRepo.LoadFromByte([]byte(sampleTicketData)))

	type args struct {
		id int
	}
	tests := []struct {
		name       string
		args       args
		wantResult int
	}{
		{"Get tickets of an existing user", args{1}, 3},
		{"Get tickets of an unexisting user", args{3}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := mockRepo.GetUserTickets(tt.args.id)
			assert.Equal(t, tt.wantResult, len(got))
		})
	}
}

func TestTicket_GetTicketsByOrg(t *testing.T) {
	mockRepo := &Ticket{}
	assert.Nil(t, mockRepo.LoadFromByte([]byte(sampleTicketData)))

	type args struct {
		id int
	}
	tests := []struct {
		name       string
		args       args
		wantResult int
	}{
		{"Get tickets of an existing org", args{119}, 2},
		{"Get tickets of an unexisting org", args{1}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := mockRepo.GetTicketsByOrg(tt.args.id)
			assert.Equal(t, tt.wantResult, len(got))
		})
	}
}
