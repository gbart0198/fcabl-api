// Package repository provides the methods used to interact with the database.
// The methods within the repository are meant to be used by passing in a database
// connection pool object, which will be passed to the generated SQLC code.
package repository

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

func ListUsers1(pg *pgxpool.Pool) {
}
