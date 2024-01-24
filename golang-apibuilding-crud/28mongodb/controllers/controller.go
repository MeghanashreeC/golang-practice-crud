package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	model "command-line-arguments/Users/meghanashree.reddy/go/src/mygolang-study-1/28mongodb/model/models.go"

	"github.com/gorilla/mux"

	"go.mongo.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/mongo/driver/mongocrypt/options"
)

const connectionString = "mongodb+srv://meghanareddy1200:42djYWw61MGCQqV4@cluster0.yxq7qsl.mongodb.net/" //connectionString is where your database is connected
const dbName = "netflix"
const colName = "watchlist"

// MOST IMPORTANT FOR MAKING CONNECTION

var collection *mongo.Collection //taking reference of the mongodb collection

// connect with mongodb

func init() { // init() runs only for the first time when this appln is running
	// 1. client option
	clientOption := options.Client().ApplyURI(connectionString) //common syntax to make the connection

	// 2. connect to mongodb
	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDB connection success")
	// to get connected inside a db and inside a collection
	collection = client.Database(dbName).Collection(colName)

	// collection instance

	fmt.Println("Connection Instance is ready")
}

// Action Plan
// Basic methods - contain all requets and response params
// Monngodb helpers - fetch data from basic method
// Helper methods
// To insert
func insertOneMovie(movie model.Netflix) {
	inserted, err := collection.InsertOne(context.Background())

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted 1 movie in db with id:", inserted.InsertedID) // Whenever a value is inserted into the db, that value gets a uniqueid --> InsertedID
}

// 2. To update
// filter and flag

func updateOneMovie(movieId string) {
	// how to convert the id such that mongodb understands
	id, err := primitive.ObjectIDFromHex(movieId)
	if err != nil {
		log.Fatal(err)
	}
	// filter the id I need to update
	filter := bson.M{"_id": id}                       //why bson --> everything that is stored in mongodb is bson
	update := bson.M{"$set": bson.M{"watched": true}} // flag $set --> to pass on the updated information

	result, err := collection.updateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Modified count:", result.ModifiedCount)
}

// 3. Delete one and delete many
// provide a filter --> to delete one particular record
// provide an empty set --> to delete all if it

// delete 1 record
func deleteOneMovie(movieId string) {
	id, err := primitive.ObjectIDFromHex(movieId)
	if err != nil {
		log.Fatal(err)
	}
	filter := bson.M{"_id": id}
	result, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Movie got deleted with delete count:", result)

}

// delete all records from mongodb

func deleteAllMovie() {
	// filter:= bson.D{{}}   // since its a common operation
	deleteResult, err := collection.DeleteMany(context.Background(), bson.D{{}}, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Number of movies deleted:", deleteResult)

}

// Read operation -> get all movies from mongodb

func getAllMovies() []primitive.M {
	cursor, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	// since i have cursor, i need to declare a variable

	var movies []primitive.M // using primitive.M for variables not bson.M

	for cursor.Next(context.Background()) { //.Next helps to loop to the next value
		var movie bson.M // for one movie
		err := cursor.Decode(&movie)
		if err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie)
	}

	defer cursor.CLose(context.Background())
	return movies

}

// Actual controllers  ---> to get all movies from DB in golang

func GetMyAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "GET")

	allMovies := getAllMovies()
	json.NewEncoder(w).Encode(allMovies)
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "POST") //What type of content you're allowing(GET,PUT,POST,etc)

	var movie model.Netflix
	_ = json.NewDecoder(r.Body).Decode(&movie)
	insertOneMovie(movie)
	json.NewEncoder(w).Encode(movie)

}

// to mark the movie as watched

func MarkAsWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")
	// we need unique value of that movie.UniqueId comes from the variables of the URL.Whenever we click on that it gives us the uniqueid
	params := mux.Vars(r)
	updateOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])

}

// Delete 1 and all movie in golang
func DeleteAMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")
	params := mux.Vars(r)
	deleteOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])

}

// Delete all movies
func DeleteAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")
	count := DeleteAllMovies()
	json.NewEncoder(w).Encode(count)

}
