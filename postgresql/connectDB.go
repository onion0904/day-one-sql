package postgresql

import(
	"database/sql"
	"fmt"
)

func Connect()(*sql.DB, error){
	connStr := fmt.Sprintf("user=onion0904 password=example dbname=db sslmode=disable")
	return sql.Open("postgres", connStr)
}