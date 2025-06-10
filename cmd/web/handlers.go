package main

import (
	"fmt"
	"net/http"
)

func (app *application) productView(w http.ResponseWriter, r *http.Request) {
	// product, err := app.queries.GetProduct(r.Context(), id)
	// if err != nil {

	// }
}

func (app *application) productAdd(w http.ResponseWriter, r *http.Request) {
	// get the category id from the category name

	// result, _ := app.queries.AddProduct(r.Context(), repository.AddProductParams{
	// 	Title:              "Ginger",
	// 	ProductDescription: "My amazing ginger product",
	// 	Category:           sql.NullInt32{Int32: 1, Valid: true},
	// 	Price:              "33.2",
	// 	ImageUrl:           "https://yoo.com",
	// 	StockQuantity:      33,
	// })
	// println(result.LastInsertId())
}

func (app *application) productEdit(w http.ResponseWriter, r *http.Request) {
	// result, err := app.queries.EditProduct(r.Context(), repository.EditProductParams{
	// 	ProductID: ,
	// })
}

func (app *application) productDelete(w http.ResponseWriter, r *http.Request) {
	// err := app.queries.DeleteProduct(r.Context(), 2)
}

func (app *application) productLatest(w http.ResponseWriter, r *http.Request) {
	// products, err := app.queries.GetLatestProducts(r.Context(), 10)
	// if err != nil {
	// 	app.serverError(w, r, err)
	// }
	// for _, product := range products {
	// 	j, _ := json.Marshal(product)
	// 	w.Write(j)
	// }
}

func (app *application) productAll(w http.ResponseWriter, r *http.Request) {
	products, err := app.queries.GetAllProducts(r.Context())
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	for _, product := range products {
		fmt.Fprintf(w, "Product: %s\n", product.Title)
	}
}
