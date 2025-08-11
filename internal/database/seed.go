package database

import (
	"context"
	"database/sql"
	"e-store-backend/internal/repository"
	"fmt"
)

func SeedDatabase(db *sql.DB, queries *repository.Queries) error {
	ctx := context.Background()

	categoryIDs := make(map[string]int32)

	// Check product categories
	existingCategories, err := queries.GetAllProductCategories(ctx)
	if err != nil {
		return fmt.Errorf("failed to check existing categories: %w", err)
	}
	if len(existingCategories) == 0 {
		// 1. Seed product categories first
		categories := []repository.AddProductCategoryParams{
			{CategoryName: "Coffee Beans", CategoryDescription: "Premium coffee beans from around the world"},
			{CategoryName: "Coffee Appliances", CategoryDescription: "Professional and home espresso machines"},
			{CategoryName: "Coffee Equipment", CategoryDescription: "Coffee grinders, filters, and others"},
		}

		for _, category := range categories {
			result, err := queries.AddProductCategory(ctx, category)
			if err != nil {
				return fmt.Errorf("failed to seed category %s: %w", category.CategoryName, err)
			}
			id, _ := result.LastInsertId()
			categoryIDs[category.CategoryName] = int32(id)
		}

		fmt.Println("Successfully seeded product categories!")
	} else {
		fmt.Println("Product categories already seeded...")
		for _, category := range existingCategories {
			categoryIDs[category.CategoryName] = category.CategoryID
		}
	}

	// Check products
	existingProducts, err := queries.GetAllProducts(ctx)
	if err != nil {
		return fmt.Errorf("failed to check existing products: %w", err)
	}
	if len(existingProducts) == 0 {
		// 2. Seed products using actual category IDs
		products := []repository.AddProductParams{
			{
				Title:              "Ethiopian Yirgacheffe Coffee Beans",
				ProductDescription: "Single-origin coffee with bright, floral notes and citrus undertones",
				Category:           categoryIDs["Coffee Beans"],
				Price:              "24.99",
				ImageUrl:           "https://example.com/ethiopian-beans.jpg",
				StockQuantity:      150,
			},
			{
				Title:              "Three Pillars",
				ProductDescription: "Smooth and sweet, medium roast with notes of nut, milk chocolate, and sweet vanilla",
				Category:           categoryIDs["Coffee Beans"],
				Price:              "21.99",
				ImageUrl:           "https://example.com/three-pillars.jpg",
				StockQuantity:      13,
			},
			{
				Title:              "Blue Bird Orchid Blend",
				ProductDescription: "Gorgeous coffee, impressively sturdy, and stays sweet and full with or without milk",
				Category:           categoryIDs["Coffee Beans"],
				Price:              "22.99",
				ImageUrl:           "https://example.com/blue-bird-orchid.jpg",
				StockQuantity:      15,
			},
			{
				Title:              "Lelit Victoria",
				ProductDescription: "Quality Italian espresso machine",
				Category:           categoryIDs["Coffee Appliances"],
				Price:              "799.99",
				ImageUrl:           "https://example.com/lelit-victoria.jpg",
				StockQuantity:      25,
			},
			{
				Title:              "Baratza Encore Coffee Grinder",
				ProductDescription: "Entry-level burr grinder perfect for home brewing",
				Category:           categoryIDs["Coffee Appliances"],
				Price:              "169.99",
				ImageUrl:           "https://example.com/baratza-encore.jpg",
				StockQuantity:      50,
			},
			{
				Title:              "Profitec Go",
				ProductDescription: "Single boiler espresso machine with ring brew group",
				Category:           categoryIDs["Coffee Appliances"],
				Price:              "1000",
				ImageUrl:           "https://example.com/profitec-go.jpg",
				StockQuantity:      3,
			},
			{
				Title:              "Hario V60 Pour Over Kit",
				ProductDescription: "Complete pour-over brewing set with dripper, filters, and server",
				Category:           categoryIDs["Coffee Equipment"],
				Price:              "39.99",
				ImageUrl:           "https://example.com/hario-v60.jpg",
				StockQuantity:      75,
			},
			{
				Title:              "KeepCup Reusable Coffee Cup",
				ProductDescription: "Eco-friendly reusable coffee cup with silicone band",
				Category:           categoryIDs["Coffee Equipment"],
				Price:              "19.99",
				ImageUrl:           "https://example.com/keepcup.jpg",
				StockQuantity:      200,
			},
			{
				Title:              "ROK EspressoGC Commercial Edition",
				ProductDescription: "An iconic manual espresso machine with factory-fitted pressure gauge",
				Category:           categoryIDs["Coffee Equipment"],
				Price:              "279",
				ImageUrl:           "https://example.com/rok-espressogc-ce.jpg",
				StockQuantity:      10,
			},
		}

		for _, product := range products {
			_, err := queries.AddProduct(ctx, product)
			if err != nil {
				return fmt.Errorf("failed to seed product %s: %w", product.Title, err)
			}
		}

		fmt.Println("Successfully seeded products!")
	} else {
		fmt.Println("Products already seeded...")
	}

	return nil
}
