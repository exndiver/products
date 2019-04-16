package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
)
//GetAllProducts
func service_all_products(w http.ResponseWriter, r *http.Request) {
	Logger1(r)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.Write(GetProductByParam("",""))
}
//GetProductByName/{product}
func service_find_products_by_name (w http.ResponseWriter, r *http.Request) {
	Logger1(r)
	vars := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")
	w.Write(GetProductByParam("name",vars["product"]))
}
//GetProductByID/{product_id}
func service_find_products_by_id (w http.ResponseWriter, r *http.Request) {
	Logger1(r)
	vars := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")
	w.Write(GetProductByParam("_id",vars["product_id"]))
}
//GetProductByGroup/{group}
func service_find_products_by_group (w http.ResponseWriter, r *http.Request) {
	Logger1(r)
	vars := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")
	w.Write(GetProductByParam("group",vars["group"]))
}
//SearchProducts/{searchrequest}
func service_search_products (w http.ResponseWriter, r *http.Request) {
	Logger1(r)
	vars := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")
	w.Write(GetProductByParam("search",vars["searchrequest"]))
}

//AddProduct (POST)
func add_new_product (w http.ResponseWriter, r *http.Request) {
	Logger1(r)
	var product Product
	body, _ := ioutil.ReadAll(r.Body)
	if err := r.Body.Close(); err != nil {panic(err)}
	if err := json.Unmarshal(body, &product); err != nil {
        w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422)
		fmt.Printf("%+s\n",err)
        if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
		return
    }
	fmt.Printf("%+s\n",body)
	w.Header().Set("Content-Type", "application/json")
	w.Write(AddNewProduct(product))
}

//Default Location method
func DefaultPage(w http.ResponseWriter, r *http.Request) {
	Logger1(r)
	w.Write([]byte("OK! Nothing!\n"))
}