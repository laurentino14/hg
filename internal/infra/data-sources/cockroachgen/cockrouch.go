package cockroachgen

import (
	"database/sql"
	"os"
	//_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/lib/pq"
)

func GenerateCockrouchDataSource() *sql.DB {
	url := os.Getenv("DATABASE_URL")
	DBConn, err := sql.Open("postgres", url)

	if err != nil {
		panic(err)
	}
	return DBConn
}
