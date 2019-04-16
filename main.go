package main

import (
    "fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
)

var Config = get_config()

func main() {
	
	r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/", DefaultPage).Methods("GET")
	r.HandleFunc("/GetAllProducts",service_all_products).Methods("GET").Headers("X-Secret-key",Config.Secret)
	r.HandleFunc("/GetProductByName/{product}",service_find_products_by_name).Methods("GET").Headers("X-Secret-key",Config.Secret)
	r.HandleFunc("/GetProductByID/{product_id}",service_find_products_by_id).Methods("GET").Headers("X-Secret-key",Config.Secret)
	r.HandleFunc("/GetProductByGroup/{group}",service_find_products_by_group).Methods("GET").Headers("X-Secret-key",Config.Secret)

	r.HandleFunc("/SearchProducts/{searchrequest}",service_search_products).Methods("GET").Headers("X-Secret-key",Config.Secret)

	r.HandleFunc("/AddProduct",add_new_product).Methods("POST").Headers("X-Secret-key",Config.Secret)
//TODOs
//
//	r.HandleFunc("/RemoveProduct/{product_id}",remove_product).Methods("DELETE")
//

	fmt.Printf("Starting server for testing HTTP POST...\n")
	log.Print(http.ListenAndServe(Config.Hosts.Service, handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
}