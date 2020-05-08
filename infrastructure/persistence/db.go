package persistence

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// NewConnection ...
func NewDbConnection(dbDriver, host, port, dbName, user, password string) (*sql.DB, error) {
	dbUrl := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, dbName)

	fmt.Println(">>> dbUrl: ", dbUrl)

	// Open up our database connection.
	db, err := sql.Open(dbDriver, dbUrl)
	if err != nil {
		return nil, err
	}

	return db, nil
}
