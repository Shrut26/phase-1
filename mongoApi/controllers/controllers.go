package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Shrut26/mongoApi/models"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://mon_go_api:mon_go_api@cluster0.hi59yys.mongodb.net/?retryWrites=true&w=majority"
const dbName = "Netflix"
const colName = "watchlist"

var collection *mongo.Collection

//mongo DB helpers

func init() {
	clientOption := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Mongo Connected")
	collection = (*mongo.Collection)(client.Database(dbName).Collection(colName))

	fmt.Println("Collection ready")
}

func insertOneMovie(movie models.Netflix) *mongo.InsertOneResult {
	inserted, err := collection.InsertOne(context.Background(), movie)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%T", inserted)
	fmt.Println(inserted)
	fmt.Println("Inserted with id", inserted.InsertedID)
	// fmt.Printf("%T ", inserted.InsertedID)
	insertedID := inserted.InsertedID.(primitive.ObjectID)
	fmt.Println("Converted id is", insertedID)
	mongoId := inserted.InsertedID
	stringObjectID := mongoId.(primitive.ObjectID).Hex()
	// return inserted.InsertedID
	fmt.Println(stringObjectID)
	// return stringObjectID
	return inserted
	// return string(inserted.InsertedID)
}

func updateOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Updated ", result.ModifiedCount)
}

func deleteOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	result, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Deleted count is : ", result.DeletedCount)
}

func deleteAllMovie() int {
	filter := bson.D{{}}
	result, err := collection.DeleteMany(context.Background(), filter)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Deleted count is: ", result.DeletedCount)
	return int(result.DeletedCount)
}

func getAllMovies() []primitive.M {
	cursor, err := collection.Find(context.Background(), bson.D{{}})

	if err != nil {
		log.Fatal(err)
	}

	var movies []primitive.M

	for cursor.Next(context.Background()) {
		var movie bson.M
		err := cursor.Decode(&movie)

		if err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie)
	}
	defer cursor.Close(context.Background())
	return movies
}

// actual controllers

func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	allMovies := getAllMovies()
	json.NewEncoder(w).Encode(allMovies)
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")
	var movie models.Netflix
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie_mongo := insertOneMovie(movie)
	// movie.ID := result.InsertedID
	// fmt.Println("*********************", result)
	// movie.ID = primitive.ObjectID{}
	// movie.ID.Hex() = movieID
	// movie.ID, _ := primitive.ObjectIDFromHex(movieID)

	result_map := make(map[string]interface{})
	result_map["_id"] = movie_mongo.InsertedID
	result_map["movie"] = movie.Movie
	result_map["watched"] = movie.Watched
	json.NewEncoder(w).Encode(result_map)
	// json.NewEncoder(w).Encode(movie)
}

func MarkAsWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")
	params := mux.Vars(r)
	updateOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteOneMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")
	params := mux.Vars(r)
	deleteOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteAllMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")
	count := deleteAllMovie()
	json.NewEncoder(w).Encode(count)
}
