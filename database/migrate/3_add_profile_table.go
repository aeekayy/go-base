package migrate

import (
	"fmt"

	"github.com/go-pg/migrations"
)

const profileTable = `
CREATE TABLE profiles (
id serial NOT NULL,
updated_at timestamp with time zone NOT NULL DEFAULT current_timestamp,
account_id uuid NOT NULL REFERENCES accounts(id),
theme text NOT NULL DEFAULT 'default',
PRIMARY KEY (id)
)`

const bootstrapAccountProfiles = `
INSERT INTO profiles(account_id) VALUES('00000000-0000-0000-0000-000000000001');
INSERT INTO profiles(account_id) VALUES('00000000-0000-0000-0000-000000000002');
`

func init() {
	up := []string{
		profileTable,
		bootstrapAccountProfiles,
	}

	down := []string{
		`DROP TABLE profiles`,
	}

	migrations.Register(func(db migrations.DB) error {
		fmt.Println("create profile table")
		for _, q := range up {
			_, err := db.Exec(q)
			if err != nil {
				return err
			}
		}
		return nil
	}, func(db migrations.DB) error {
		fmt.Println("drop profile table")
		for _, q := range down {
			_, err := db.Exec(q)
			if err != nil {
				return err
			}
		}
		return nil
	})
}
