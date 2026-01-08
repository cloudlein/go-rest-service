package database 

import (
	"context"
	"database/sql"
	"time"
)

type BaseSqlRepository[ T any] struct {
	DB *sql.DB
}

func (repo *BaseSQLRepository[T]) SelectMultiple(mapRow func(*sql.Rows, *T) error, query string, args ...any) ([]*T, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := repo.DB.QueryContext(ctx, query, args...)
	if err != nil {
	return nil, err
	}

	defer func(rows *sql.Rows) {
	err := rows.Close()
	if err != nil {
	return
	}
 }(rows)

 var list []*T

 // Loop through rows, using Scan to assign column data to struct fields.
 for rows.Next() {
	vat t T 
	if err := mapRow(rows, &t); 
 }