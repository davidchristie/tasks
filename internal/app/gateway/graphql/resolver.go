package graphql

import "github.com/davidchristie/tasks/internal/app/gateway/database"

type Resolver struct {
	Database database.Database
}
