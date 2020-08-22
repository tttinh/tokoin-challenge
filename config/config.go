package config

import (
	"github.com/spf13/viper"
)

// Configuration struct.
type Configuration struct {
	UserData         string
	TicketData       string
	OrganizationData string
}

// New creates a new config.
func New() *Configuration {
	return &Configuration{}
}

// Init ...
func (cfg *Configuration) Init(filepath, filename string) error {
	viper.AddConfigPath(filepath)
	viper.SetConfigName(filename)
	viper.SetConfigType("toml")

	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	cfg.UserData = viper.GetString("Data.User")
	cfg.TicketData = viper.GetString("Data.Ticket")
	cfg.OrganizationData = viper.GetString("Data.Organization")
	return nil
}
