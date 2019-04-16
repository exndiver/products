package main

import (
    "context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func db_connect() *mongo.Client{
	client, err := mongo.NewClient(options.Client().ApplyURI(Config.Hosts.Mongodb))
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	if err = client.Connect(ctx); err != nil {log.Fatal(err)}
	return client
}

//Getting collection from DB
func db_get_collection(param, search string) *mongo.Cursor {
	client := db_connect()
	collection := client.Database("products").Collection("product")	
	var cur *mongo.Cursor
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
	ctx,_:= context.WithTimeout(context.Background(), 30*time.Second)
	cur, err := collection.Find(ctx, filter)
	if err != nil { log.Fatal(err) }					
	client.Disconnect(ctx)
	cur.Close(ctx)
	return cur
}

func db_add (product Product) *mongo.SingleResult{
	client := db_connect()
	ctx,_:= context.WithTimeout(context.Background(), 30*time.Second)
	collection := client.Database("products").Collection("product")
	res, err := collection.InsertOne(ctx,product)
	if err != nil { log.Fatal(err) }
	cur := collection.FindOne(ctx, bson.D{{"_id",res.InsertedID}})
	client.Disconnect(ctx)
	return cur
}