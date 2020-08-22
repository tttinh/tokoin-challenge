package main

import (
	"github.com/tttinh/tokoin-challenge/model"
	"github.com/tttinh/tokoin-challenge/repo"
	"github.com/tttinh/tokoin-challenge/utils"
)

const seperateLine = "\n------------------------------------------------------------"
const welcomeMessage = `Type 'quit' to exit at anytime, press Enter to continue.

	Select search options:
	 * Type '1' to search.
	 * Type '2' to view a list of searchable fields.
`

// App implements our main application logic.
type App struct {
	repository repo.Repository
}

// NewApp creates a new application instance.
func NewApp(rp repo.Repository) *App {
	return &App{repository: rp}
}

// Run starts the main loop of the application.
func (app *App) Run() {
	for {
		utils.WriteLine(welcomeMessage)
		cmd, quit := readUserInput()
		if quit {
			return
		}

		if quit := app.handleCommand(cmd); quit {
			return
		}
	}
}

func readUserInput() (line string, quit bool) {
	if line = utils.ReadLine(); line == "quit" {
		quit = true
		return
	}

	quit = false
	return
}

func (app *App) handleCommand(cmd string) bool {
	switch cmd {
	case "1":
		return app.handleSearch()
	case "2":
		showSearchableFields()
	}

	return false
}

// handleSearch searchs information based on user input.
func (app *App) handleSearch() bool {
	utils.WriteLine(seperateLine)
	utils.WriteLine("Select 1) Users or 2) Tickets or 3) Organizations")
	var searchCategory string
	for {
		line, quit := readUserInput()
		if quit {
			return true
		}

		if line == "1" || line == "2" || line == "3" {
			searchCategory = line
			break
		}
		utils.WriteLine("Please enter '1', '2' or '3':")
	}

	utils.WriteLine("Enter search term:")
	searchTerm, quit := readUserInput()
	if quit {
		return true
	}

	utils.WriteLine("Enter search value:")
	searchValue, quit := readUserInput()
	if quit {
		return true
	}

	var results repo.Results
	var err error
	switch searchCategory {
	case "1":
		utils.WriteLine("Searching users for", searchTerm, "with a value of", searchValue)
		results, err = app.repository.FindUsers(searchTerm, searchValue)
	case "2":
		utils.WriteLine("Searching tickets for", searchTerm, "with a value of", searchValue)
		results, err = app.repository.FindTickets(searchTerm, searchValue)
	case "3":
		utils.WriteLine("Searching organizations for", searchTerm, "with a value of", searchValue)
		results, err = app.repository.FindOrganizations(searchTerm, searchValue)
	}

	if err != nil {
		utils.WriteLine("Error: ", err)
	} else {
		utils.WriteLine(results.Length(), "results found.")
		results.Print()
	}

	utils.WriteLine(seperateLine)
	return false
}

// showSearchableFields shows all searchable fields.
func showSearchableFields() {
	utils.WriteLine(seperateLine)
	utils.WriteLine("Search Users with:")
	utils.PrintJSONTags(&model.User{})
	utils.WriteLine(seperateLine)
	utils.WriteLine("Search Tickets with:")
	utils.PrintJSONTags(&model.Ticket{})
	utils.WriteLine(seperateLine)
	utils.WriteLine("Search Organizations with:")
	utils.PrintJSONTags(&model.Organization{})
	utils.WriteLine(seperateLine)
}
