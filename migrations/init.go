package migrations

import (
	"github.com/camarin24/go-studio/core"
)

func Migrate(app core.App) {
	if _, err := app.DB().NewQuery(`
		CREATE TABLE IF NOT EXISTS projects (
			id TEXT PRIMARY KEY NOT NULL,
			name TEXT UNIQUE NOT NULL,
			created TEXT DEFAULT (strftime('%Y-%m-%d %H:%M:%fZ')) NOT NULL,
			updated TEXT DEFAULT (strftime('%Y-%m-%d %H:%M:%fZ')) NOT NULL
		)
	`).Execute(); err != nil {
		panic(err)
	}
}
