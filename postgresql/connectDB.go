package postgresql

import(
	"database/sql"
	"fmt"
)

func Connect()(*sql.DB, error){
	connStr := fmt.Sprintf("host=127.0.0.1 port=5432 user=onion0904 password=example dbname=db sslmode=disable")
	return sql.Open("postgres", connStr)
}