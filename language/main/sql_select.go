package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

func main() {
	// Open a connection to the database
	db, _ := sql.Open("driverName", "dataSourceName")

	// Create a context with a timeout of 1 second
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Execute the query with the context
	rows, err := db.QueryContext(ctx, "SELECT * FROM table")
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	// Handle the query results
	// ...
}
