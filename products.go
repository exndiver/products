package main

import (
    "context"
	"log"
	"time"
	"encoding/json"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	ID primitive.ObjectID `json:"id" bson:"_id,omitempty"`
    Name string `json:"name"`
    Group  string `json:"group"`
	Kcal float64 `json:"kcal"`
	Composition Composition `json:"composition"`
}
type Composition struct {
    Carbohydrate float64 `json:"carbohydrate"`
    Protein  float64 `json:"protein"`
	Fat float64 `json:"fat"`
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

func AddNewProduct(product Product)[]byte{
	res := db_add(product)
	err := res.Decode(&product)
	if err != nil {log.Fatal(err)}
	result,_ := json.Marshal(product)
	return result
}