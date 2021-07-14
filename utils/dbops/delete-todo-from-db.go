package dbops

import (
	"database/sql"
	"fmt"
)

func DeleteTodoFromDb(id int, username string) (sql.Result, error) {
	// insert to the new `USERS` tanle
	instruction := fmt.Sprintf("delete from todos where id=%v and username='%v'", id, username)
	return Exec(instruction)
}
