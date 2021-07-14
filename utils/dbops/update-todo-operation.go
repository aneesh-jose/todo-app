package dbops

import (
	"database/sql"
	"fmt"
)

func UpdateTodoOperation(id int, status bool) (sql.Result, error) {
	query := fmt.Sprintf("update todos set status=%v where id=%v", status, id)
	return Exec(query)

}
