package database

import (
	"fmt"
	"time"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file" // file driver
	_ "github.com/lib/pq"                                // postgres driver
)

func (d *database) Migrate(url string) error {
	for {
		err := d.db.Ping()
		if err == nil {
			break
		}
		fmt.Println("Database migration failed: ", err)
		fmt.Println("Retrying in 5 seconds...")
		time.Sleep(5 * time.Second)
	}

	driver, err := postgres.WithInstance(d.db, &postgres.Config{})
	if err != nil {
		return err
	}
	m, err := migrate.NewWithDatabaseInstance(url, "postgres", driver)
	if err != nil {
		return err
	}
	m.Up()
	fmt.Println("Database migration complete")
	return nil
}
