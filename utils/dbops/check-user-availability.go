package dbops

import (
	"database/sql"
	"fmt"

	"github.com/aneesh-jose/simple-server/models"
)

func CheckUserAvailability(user models.User) (*sql.Rows, error) {
	query := fmt.Sprintf("select username from users where username='%v' and password='%v'", user.Username, user.Password)

	return Query(query)
}
