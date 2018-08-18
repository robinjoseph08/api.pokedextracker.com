package main

import (
	"github.com/go-pg/pg/orm"
	"github.com/robinjoseph08/go-pg-migrations"
)

func init() {
	up := func(db orm.DB) error {
		_, err := db.Exec(`
			ALTER TABLE users
				ADD COLUMN last_ip VARCHAR(45),
				ADD COLUMN last_login TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
		`)
		return err
	}

	down := func(db orm.DB) error {
		_, err := db.Exec(`
			ALTER TABLE users
				DROP COLUMN last_ip,
				DROP COLUMN last_login
		`)
		return err
	}

	opts := migrations.MigrationOptions{}

	migrations.Register("20180813200146_add_last_login_and_last_ip_to_users", up, down, opts)
}
