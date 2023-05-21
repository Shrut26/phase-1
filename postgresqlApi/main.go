package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Person struct {
	gorm.Model
	Name  string
	Email string `gorm:"typevarchar(100);unique_index"`
	Books []Book
}

type Book struct {
	gorm.Model
	Title      string
	Author     string
	CallNumber int
	PersonID   int `gorm:"unique_index"`
}

// var (
// 	person = &Person{Name: "Jack", Email: "jack@email.com"}
// 	books  = []Book{
// 		{Title: "firstbook", Author: "firstauthor", CallNumber: 1234, PersonID: 1},
// 		{Title: "secondbook", Author: "secondauthor", CallNumber: 2345, PersonID: 1},
// 		{Title: "thirdbook", Author: "thirdauthor", CallNumber: 3456, PersonID: 1},
// 	}
// )

var db *gorm.DB

// var err error

func main() {
	// environmental variables
	// dialect := os.Getenv("DIALECT")
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
	host := os.Getenv("HOST")
	dbport := os.Getenv("DBPORT")
	user := os.Getenv("USER")
	dbname := os.Getenv("NAME")
	password := os.Getenv("PASSWORD")

	//connection string
	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%s", host, user, dbname, password, dbport)
	fmt.Println(dbURI)
	//opeing connection
	db, err := gorm.Open(postgres.Open(dbURI), &gorm.Config{})

	if err != nil {
		fmt.Println(dbURI)

		log.Fatal(err)
	} else {
		fmt.Println("Successfully connected to database")
	}

	db.AutoMigrate(&Person{})
	db.AutoMigrate(&Book{})

	r := mux.NewRouter()
	r.HandleFunc("/people", getPeople).Methods("GET")

	http.ListenAndServe(":4000", r)
}

func getPeople(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "GET")
	var people []Person
	err := db.Find(&people)
	if err.Error != nil {
		json.NewEncoder(w).Encode("No record")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(people)

	if err != nil {
		log.Fatal(err)
	}
}
