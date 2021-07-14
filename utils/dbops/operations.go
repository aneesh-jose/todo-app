package dbops

import (
	"database/sql"
)

func Query(instruction string) (*sql.Rows, error) {
	psqlInfo := GetDbCreds()                  //obtain database credentials
	db, err := sql.Open("postgres", psqlInfo) //connect to database
	if err != nil {
		// database connection error
		// maybe database server is down or the
		// authentication credentials might have changed
		return nil, err
	}
	data, err := db.Query(instruction)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	return data, nil
}

func Exec(instruction string) (sql.Result, error) {

	psqlInfo := GetDbCreds()                  //obtain database credentials
	db, err := sql.Open("postgres", psqlInfo) //connect to database
	if err != nil {
		// database connection error
		// maybe database server is down or the
		// authentication credentials might have changed
		return nil, err
	}
	data, err := db.Exec(instruction)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	return data, nil
}
