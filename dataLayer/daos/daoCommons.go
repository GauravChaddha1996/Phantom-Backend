package daos

import (
	"database/sql"
	"errors"
)

func prepareAndExecuteInsertQuery(db *sql.DB, query string, args ...interface{}) error {

	// Prepare the statement
	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Execute the query
	result, err := stmt.Exec(args...)
	if err != nil {
		return err
	}

	// Check if rows were inserted or not
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("rows affected is 0")
	}
	return nil
}

func prepareAndExecuteSelectQuery(db *sql.DB, query string, args ...interface{}) (*sql.Rows, error) {

	// Prepare the statement
	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	// Execute the query
	rows, err := stmt.Query(args...)
	if err != nil {
		return nil, err
	}

	return rows, nil
}