package main

import (
	"github.com/go-pg/pg/orm"
	"github.com/robinjoseph08/go-pg-migrations"
)

func init() {
	up := func(db orm.DB) error {
		if _, err := db.Exec("ALTER TABLE pokemon ADD COLUMN game_family_id TEXT"); err != nil {
			return err
		}
		if _, err := db.Exec("ALTER TABLE pokemon ADD CONSTRAINT pokemon_game_family_id_foreign FOREIGN KEY (game_family_id) REFERENCES game_families (id)"); err != nil {
			return err
		}
		if _, err := db.Exec("CREATE INDEX pokemon_game_family_id_index ON pokemon (game_family_id)"); err != nil {
			return err
		}
		if _, err := db.Exec(`
			UPDATE pokemon
			SET
				game_family_id = gf.id
			FROM (
				SELECT
					DISTINCT(f1.id),
					f1.generation
				FROM game_families f1
				INNER JOIN game_families f2 ON f1.generation = f2.generation
				GROUP BY
					1, 2
				HAVING f1.order = MIN(f2.order)
			) AS gf
			WHERE
				pokemon.generation = gf.generation

		`); err != nil {
			return err
		}
		_, err := db.Exec("ALTER TABLE pokemon ALTER COLUMN game_family_id SET NOT NULL")
		return err
	}

	down := func(db orm.DB) error {
		_, err := db.Exec("ALTER TABLE pokemon DROP COLUMN game_family_id")
		return err
	}

	opts := migrations.MigrationOptions{}

	migrations.Register("20180818090625_add_game_family_id_to_pokemon", up, down, opts)
}
