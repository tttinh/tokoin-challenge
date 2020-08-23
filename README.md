# tokoin-challenge

A small project for interview at Tokoin.

## Challenge

Using the provided data (tickets.json and users.json and organization.json), write a simple command line application to search the data and return the results in a human readable format.

* Where the data exists, values from any related entities MUST be included in the results, i.e.
  * Searching organization MUST return its **ticket subject** and **users name**.
  * Searching users MUST return his/her **assignee ticket subject** and **submitted ticket subject** and his/her **organization name**.
  * Searching tickets MUST return its **assignee name**, **submitter name**, and **organization name**.
* The user should be able to search on any field, full value matching is fine(e.g. "mar" won't return "mary").
* The user should also be able to search for empty values, e.g. where description is empty.

## How to build and run

### Prerequisites

Go 1.14 should be installed on your system. Please visit this [website](https://golang.org/) for more information.

### Run unit tests

```bash
go test ./... -v
```

### Build and run the program

```bash
go build
./tokoin-challenge
```

## Folder structure

```text
root
├── config
│   ├── config.go ( Implement configuration reading )
│   └── default.toml (Application config file)
│
├── data ( Contain JSON files which are search data )
├── model ( Define all object models )
├── repo ( Implement data layers which are responsible for searching )
├── utils ( Implement helper functions )
├── app.go ( Main application logic, handle user's input )
├── main.go ( Program entry point )
└── README.md ( Describe Project )
```

## Program workflow

* Loading configuration.
* Loading and building data layers.
* Start an infinite loop to handle user's input.

## Notes

* Program is built and tested on Ubuntu 18.04.
* Because the demo program is built on MacOS, so that I can't run it. Therefore, this implementation is based on the instruction file only.
* Program support searching on all fields of all collections (User, Ticket, Organization), and on each collection only **_id** field have been indexed using map data structure. But we can easily build index for other fields with the same technique.

* Our core logic functions are at the data layer, so unit tests focus most on this (in the **repo** folder). The test coverage is not so high (about 70%) but enough to illustrate the testing method.

## Thanks for reading