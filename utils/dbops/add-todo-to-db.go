package dbops

import (
	"database/sql"
	"fmt"
)

func AddTodoToDB(name string, description string, username string) (sql.Result, error) {
	query := fmt.Sprintf("insert into todos values(nextval('countsequence'),'%v','%v',%v,'%v')", name, description, false, username)
	return Exec(query)
}
