package databasDaos

import (
	"database/sql"
	"errors"
	"log"
)

func prepareAndExecuteInsertQuery(db *sql.DB, query string, args ...interface{}) (*int64, error) {

	// Prepare the statement
	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer closeStmt(stmt)

	// Execute the query
	result, err := stmt.Exec(args...)
	if err != nil {
		return nil, err
	}

	// Check if rows were inserted or not
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}
	if rowsAffected == 0 {
		return nil, errors.New("rows affected is 0")
	}

	// Return last inserted Id
	lastInsertId, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	return &lastInsertId, nil
}

func prepareAndExecuteSingleRowSelectQuery(db *sql.DB, query string, args ...interface{}) (*sql.Row, error) {

	// Prepare the statement
	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer closeStmt(stmt)

	// Execute the query
	row := stmt.QueryRow(args...)

	return row, nil
}

func prepareAndExecuteSelectQuery(db *sql.DB, query string, args ...interface{}) (*sql.Rows, error) {

	// Prepare the statement
	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer closeStmt(stmt)

	// Execute the query
	rows, err := stmt.Query(args...)
	if err != nil {
		return nil, err
	}

	return rows, nil
}

func closeStmt(stmt *sql.Stmt) {
	err := stmt.Close()
	if err != nil {
		log.Print(err)
	}
}

func closeRows(rows *sql.Rows) {
	err := rows.Close()
	if err != nil {
		log.Print(err)
	}
}
