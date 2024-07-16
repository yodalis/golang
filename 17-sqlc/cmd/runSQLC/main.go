package main

import (
	"context"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/yodalis/golang/17-sqlc/internal/db"
)

func main() {
	ctx := context.Background()
	dbConnection, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/courses")
	if err != nil {
		panic(err)
	}

	defer dbConnection.Close()

	queries := db.New(dbConnection)

	// err = queries.CreateCategory(ctx, db.CreateCategoryParams{
	// 	ID:          uuid.New().String(),
	// 	Name:        "backend",
	// 	Description: sql.NullString{String: "Backend description", Valid: true},
	// })

	// if err != nil {
	// 	panic(err)
	// }

	// categories, err := queries.ListCategories(ctx)
	// if err != nil {
	// 	panic(err)
	// }

	// for _, category := range categories {
	// 	println(category.ID, category.Name, category.Description.String)
	// }

	// err = queries.UpdateCategory(ctx, db.UpdateCategoryParams{
	// 	ID:          "c0c44d12-1311-4f9a-b04f-2fe8c90ed2af",
	// 	Name:        "Backend Updated",
	// 	Description: sql.NullString{String: "Updated Backend description", Valid: true},
	// })

	err = queries.DeleteCategory(ctx, "c0c44d12-1311-4f9a-b04f-2fe8c90ed2af")
	if err != nil {
		panic(err)
	}
	categories, err := queries.ListCategories(ctx)
	if err != nil {
		panic(err)
	}

	for _, category := range categories {
		println(category.ID, category.Name, category.Description.String)
	}
}
