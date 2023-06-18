package runtime

import (
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	_ "github.com/mattn/go-sqlite3"
)

func NewSQLStore() (dialect.Driver, error) {
	// Open the database connection.
	drv, err := sql.Open(dialect.SQLite, "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		return nil, err
	}

	return drv, nil
}
