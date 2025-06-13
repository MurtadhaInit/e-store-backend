package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func (app *application) productView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	product, err := app.queries.GetProduct(r.Context(), int32(id))
	if err != nil {
		app.clientError(w, http.StatusNotFound)
		return
	}

	// TODO
	// category, err := app.queries.GetProductCategory(r.Context(), product.Category.Int32)
	// if err != nil {
	// }

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
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
