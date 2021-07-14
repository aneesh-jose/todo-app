package dbops

import (
	"database/sql"
	"fmt"
)

func ReadTodosFromDb(username string) (*sql.Rows, error) {

	query := fmt.Sprintf("select * from todos where username='%v'", username)

	return Query(query)
}
