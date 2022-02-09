package Dependencies

import (
	"FirstAPI/pkg/Dependencies/Database"
)

type Dependencies struct {
	Database *Database.Database
}

func InitDependencies() Dependencies {
	return Dependencies{
		Database: Database.InitDatabase(),
	}
}
