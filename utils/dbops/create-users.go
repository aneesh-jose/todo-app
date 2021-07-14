package dbops

import (
	"database/sql"
	"fmt"
)

func AddUserToDb(username string, password string, name string) (sql.Result, error) {
	// insert to the new `USERS` tanle
	instruction := fmt.Sprintf("insert into users values('%v','%v','%v')", username, password, name)
	return Exec(instruction)
}
