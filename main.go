package main

import (
    "context"
    "fmt"
	"log"
	"time"
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Main structure for products
type Product struct {
	ID primitive.ObjectID `json:"id" bson:"_id"`
    Name string `json:"name,omitempty"`
    Group  string `json:"group,omitempty"`
	Kcal float64 `json:"kcal,omitempty"`
	Composition Composition `json:"composition,omitempty"`
}
type Composition struct {
    Carbohydrate float64 `json:"carbohydrate,omitempty"`
    Protein  float64 `json:"protein,omitempty"`
	Fat float64 `json:"fat,omitempty"`
}

//Getting collection from DB
func db_get_collection(param, search string) *mongo.Cursor {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {log.Fatal(err)}
	collection := client.Database("products").Collection("product")	
	ctx,_= context.WithTimeout(context.Background(), 30*time.Second)
	var cur *mongo.Cursor
	if err != nil { log.Fatal(err) }
	filter := bson.D{}
	switch param {
		case "name":
			filter = bson.D{{param, search}}
		case "group":
			filter = bson.D{{param, search}}
		case "_id":
			objid,_ :=primitive.ObjectIDFromHex(search)
			filter = bson.D{{param, objid}}
		case "search":
			filter = bson.D{{"name", bson.M{"$regex":search,"$options":"i"}}}
	}
	if param!=""{}
	cur, err = collection.Find(ctx, filter)
	if err != nil { log.Fatal(err) }					
	client.Disconnect(ctx)
	cur.Close(ctx)
	return cur
}

func GetProductByParam(param,name string)[]byte{
	var result []*Product
	ctx,_ := context.WithTimeout(context.Background(), 30*time.Second)
	cur:= db_get_collection(param,name)
	for cur.Next(ctx) {
		var elem Product
		err := cur.Decode(&elem)
		if err != nil {log.Fatal(err)}
		result = append(result, &elem)
	}
	if err := cur.Err(); err != nil {log.Fatal(err)}
	json_result, _ := json.Marshal(result)
	return json_result
}

//API Methods
func service_all_products(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write(GetProductByParam("",""))
}

func service_find_products_by_name (w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")
	w.Write(GetProductByParam("name",vars["product"]))
}

func service_find_products_by_id (w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")
	w.Write(GetProductByParam("_id",vars["product_id"]))
}

func service_find_products_by_group (w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")
	w.Write(GetProductByParam("group",vars["group"]))
}

func service_search_products (w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")
	w.Write(GetProductByParam("search",vars["searchrequest"]))
}

//Default Location method
func DefaultPage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK! Nothing!\n"))
}

//Main Func
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", DefaultPage).Methods("GET")

	r.HandleFunc("/GetAllProducts",service_all_products).Methods("GET")
	r.HandleFunc("/GetProductByName/{product}",service_find_products_by_name).Methods("GET")
	r.HandleFunc("/GetProductByID/{product_id}",service_find_products_by_id).Methods("GET")
	r.HandleFunc("/GetProductByGroup/{group}",service_find_products_by_group).Methods("GET")

	r.HandleFunc("/SearchProducts/{searchrequest}",service_search_products).Methods("GET")

//TODOs
//
//	r.HandleFunc("/AddProduct/",add_new_product).Methods("POST")
//
//	r.HandleFunc("/AddProduct/",remove_product).Methods("DELETE")
//

	fmt.Printf("Starting server for testing HTTP POST...\n")
	log.Fatal(http.ListenAndServe(":8080", r))
}