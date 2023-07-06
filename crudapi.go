package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

type Course struct {
	CourseId    string  `json:"courseid"`
	Coursename  string  `json:"coursename"`
	CoursePrice int     `json:"courseprice"`
	Author      *Author `json:"author"`
}

type Author struct {
	FullName string `json:"fullname"`
	Website  string `json:"website"`
}

// fake database.
var courses []Course

// middleware helper file;
func (c *Course) IsEmpty() bool {
	//can not rely on user to enter id's will generate on our own
	//return c.CourseId == "" && c.Coursename == ""
	return c.Coursename == ""
}

func main() {

	fmt.Println("API - Pratham")
	r := mux.NewRouter()

	//seeding
	courses = append(courses, Course{CourseId: "2", Coursename: "ReactJS", CoursePrice: 299, Author: &Author{FullName: "Pratham Goel", Website: "udemy.com"}})
	courses = append(courses, Course{CourseId: "4", Coursename: "MERN Stack", CoursePrice: 199, Author: &Author{FullName: "Pratham Goel", Website: "udemy.com"}})

	//routing
	r.HandleFunc("/", servehome).Methods("GET")
	r.HandleFunc("/courses", getallcourse).Methods("GET")
	r.HandleFunc("/course/{id}", getonecourse).Methods("GET")
	r.HandleFunc("/course", createonecourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")

	// listen to a port
	log.Fatal(http.ListenAndServe(":4000", r))

}

//controllers

func servehome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>API TESTING</h1>"))
}
func getallcourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get all course func")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getonecourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course func")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	fmt.Println(params)

	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	json.NewEncoder(w).Encode("No Course found with given id")
	return
}

func createonecourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course func")
	w.Header().Set("Content-Type", "application/json")

	//if the body is empty.
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send me some data")
	}

	//if the body is like {}

	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No data inside JSON")
		return
	}

	//to-check only if title is duplicate
	//loop, title matches with course.coursename, JSON

	// generate unique id, string
	// append course into courses

	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return

}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update one course")
	w.Header().Set("Content-Type", "application/json")

	// first - grab id from req
	params := mux.Vars(r)

	// loop, id, remove, add with my ID

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...) //append function expects more arguments
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}

}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete one course")
	w.Header().Set("Content-Type", "applicatioan/json")

	params := mux.Vars(r)

	//loop, id, remove (index, index+1)

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			break

		}
	}
}
