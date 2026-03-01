package db

import (
	"github.com/danielleit241/internal/config"
)

var DBConnectionString string

func InitDB(config *config.Config) error {
	DBConnectionString = config.GetDatabaseDSN()

	return nil
}
