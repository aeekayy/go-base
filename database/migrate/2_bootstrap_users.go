package migrate

import (
	"fmt"

	"github.com/go-pg/migrations"
)

const bootstrapAdminAccount = `
INSERT INTO accounts (id, email, name, active, roles)
VALUES ('00000000-0000-0000-0000-000000000001', 'admin@boot.io', 'Admin Boot', true, '{admin}')
`

const bootstrapUserAccount = `
INSERT INTO accounts (id, email, name, active)
VALUES ('00000000-0000-0000-0000-000000000002', 'user@boot.io', 'User Boot', true)
`

func init() {
	up := []string{
		bootstrapAdminAccount,
		bootstrapUserAccount,
	}

	down := []string{
		`TRUNCATE accounts CASCADE`,
	}

	migrations.Register(func(db migrations.DB) error {
		fmt.Println("add bootstrap accounts")
		for _, q := range up {
			_, err := db.Exec(q)
			if err != nil {
				return err
			}
		}
		return nil
	}, func(db migrations.DB) error {
		fmt.Println("truncate accounts cascading")
		for _, q := range down {
			_, err := db.Exec(q)
			if err != nil {
				return err
			}
		}
		return nil
	})
}
