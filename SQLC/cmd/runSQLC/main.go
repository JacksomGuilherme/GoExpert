package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/JacksomGuilherme/GoExpert/SQLC/internal/db"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	ctx := context.Background()
	dbConn, err := sql.Open("mysql", "golang:golang@tcp(localhost:3306)/goexpert")
	if err != nil {
		panic(err)
	}
	defer dbConn.Close()

	queries := db.New(dbConn)

	// err = queries.CreateCategory(ctx, db.CreateCategoryParams{
	// 	ID:          uuid.New().String(),
	// 	Name:        "Backend",
	// 	Description: sql.NullString{String: "Backend description", Valid: true},
	// })
	// if err != nil {
	// 	panic(err)
	// }

	// err = queries.UpdateCategory(ctx, db.UpdateCategoryParams{
	// 	ID:          "0d54a5bd-1c9c-4e62-991e-957eee859daa",
	// 	Name:        "Backend updated",
	// 	Description: sql.NullString{String: "Backend description updated", Valid: true},
	// })

	err = queries.DeleteCategory(ctx, "0d54a5bd-1c9c-4e62-991e-957eee859daa")

	categories, err := queries.ListCategories(ctx)
	if err != nil {
		panic(err)
	}

	for _, category := range categories {
		fmt.Println(category.ID, category.Name, category.Description)
	}
}
