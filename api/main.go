package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// models
type Course struct {
	CourseId    string  `json : "courseid"`
	Coursename  string  `json : "coursename`
	Courseprice int     `json : "price"`
	Author      *Author `json : "author"`
}

type Author struct {
	Fullname string `json : "fullname"`
	Website  string `json : "website"`
}

// fake DB
var courses []Course

// middlewares
func (c *Course) Isempty() bool {
	return c.Coursename == ""
}

// controllers
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>hello</h1>"))
}
func getAllCourses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if len(courses) > 0 {
		json.NewEncoder(w).Encode(courses)
	} else {
		json.NewEncoder(w).Encode("Empty List")
	}
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	fmt.Println(params)

	for key, value := range courses {
		if courses[key].CourseId == params["id"] {
			json.NewEncoder(w).Encode(value)
			return
		}
	}

	json.NewEncoder(w).Encode("No Course found")
	return
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Tyep", "application/json")

	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
	}

	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.Isempty() {
		json.NewEncoder(w).Encode("No data inside json")
		return
	}

	rand.Seed(time.Now().Unix())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return
}

func updatecourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	fmt.Println(params)
	for key, _ := range courses {
		if courses[key].Coursename == params["coursename"] {
			var course Course
			json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = courses[key].CourseId
			courses[key].Coursename = course.Coursename
			courses[key].Courseprice = course.Courseprice
			courses[key].Author.Fullname = course.Author.Fullname
			courses[key].Author.Website = course.Author.Website
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No such course found")
	return
}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for key, _ := range courses {
		if courses[key].Coursename == params["coursename"] {
			courses = append(courses[:key], courses[key+1:]...)
			json.NewEncoder(w).Encode(courses)
			return
		}
	}
	json.NewEncoder(w).Encode("No such course found")
}
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/course", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course/create", createOneCourse).Methods("POST")
	r.HandleFunc("/course/update/{coursename}", updatecourse).Methods("PUT")
	r.HandleFunc("/course/delete/{coursename}", deleteOneCourse).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":4000", r))
}
