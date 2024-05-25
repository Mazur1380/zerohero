package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const mongoURL = "mongodb://user:pass@127.0.0.1:27017/?retryWrites=true&w=majority"

type User struct {
	ID      primitive.ObjectID `bson:"_id,omitempty"`
	Name    string             `bson:"name,omitempty"`
	Courses []string           `bson:"courses,omitempty"`
	City    string             `bson:"city,omitempty"`
	Age     int                `bson:"age,omitempty"`
}

func main() {

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(mongoURL))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.Background())

	col := client.Database("mydb").Collection("users")
	cursor, err := col.Find(context.Background(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}

	var results []User
	err = cursor.All(context.Background(), &results)
	if err != nil {
		log.Fatal(err)
	}
	for _, r := range results {
		fmt.Printf("%v+\n", r)
	}
}

// 	u := []interface{}{
// 		User{
// 			Name:    "Djas",
// 			Courses: []string{"Python"},
// 			City:    "Hamburg",
// 			Age:     24,
// 		},
// 		User{
// 			Name: "Dima",
// 			City: "SPB",
// 			Age:  24,
// 		},
// 	}

// 	// u := User{
// 	// 	Name:    "Vlad",
// 	// 	Courses: []string{"Golang", "Python"},
// 	// 	City:    "SPB",
// 	// 	Age:     28,
// 	// }

// 	res, err := client.Database("mydb").Collection("users").InsertMany(context.Background(), u)
// 	if err != nil {
// 		log.Fatal(err)

// 	}
// 	fmt.Println("Insert ids:", res.InsertedIDs)
// }
