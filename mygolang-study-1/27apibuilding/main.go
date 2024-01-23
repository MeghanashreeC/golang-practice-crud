// This section deals with building a complete API without injecting database
// We are not going to breakdown code and put it into seperate files.
// we will work on only one file
// We will understand a clarity on how API is built using golang

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

// We need to define how our course is going to look like --->  Model for course - file
// 1. create struct
type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"` // for referencing from below

}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// lets create a fake database using slice
var courses []Course

// helper method / middleware ---> file 2 ---> they help to perform tasks like encrypting passwords,etc

// Example below is used to --> to make sure CourseId and CourseName should not be empty
// This below helper method tells us to proceed only if both the fields are filled and not empty

func (c *Course) ISEmpty() bool { // since this method is structure, define structure to get access to all the values from struct
	// return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == ""
}

func main() {

	fmt.Println("API - Golang")
	// create a router
	r := mux.NewRouter()
	// seeding of data -- injecting data into courses i.e., into the slice we created
	courses = append(courses, Course{CourseId: "2", CourseName: "ReactJS", CoursePrice: 299, Author: &Author{Fullname: "Meghana", Website: "golang.co"}})
	courses = append(courses, Course{CourseId: "4", CourseName: "GO", CoursePrice: 499, Author: &Author{Fullname: "Meghana", Website: "golang.co"}})

	// Routing Starts

	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET") // {id}: defined in params
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", DeleteOneCourse).Methods("DELETE")

	// 1. Listening to a port

	log.Fatal(http.ListenAndServe(":4000", r))

}

// Action Plan
//1. We will create a learncode online similar kind of looking backend api
//2. We want to give a feature that the user can get all the courses, create courses, update courses and delete courses
//3. I want to see if there is no unique id or title given to a course, then do not add the course.(creating a helper node)
//4. we will be creating a slice as there is no database yet
//5. We will be injecting gorilla/mux so that every single route is created for CRUD operation

// controllers - file 3
// servehome route

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to API using Golang</h1>"))

}

// I want to throw all values from database

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")

	// method to set headers
	w.Header().Set("Content-Type", "application/json")
	// how to throw all things in db as JSON
	json.NewEncoder(w).Encode(courses)
	return
}

// TO GET course based on request id in golang

func getOneCourse(w http.ResponseWriter, r *http.Request) { // They provide a way to send data back to the client and to extract information from the incoming HTTP request.
	// task 1 - grab request id
	// task 2 - run a loop to compare the id, if id exists
	fmt.Println("Get one course")
	w.Header().Set("Content-Type", "application/json")

	// grab id from a request --> params is a holder of key-value pair
	params := mux.Vars(r)

	// loop through courses, find matching id and return the response
	for _, course := range courses {
		// to find the matching id
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course) // to return the response. writer is going to write
			return
		}
	}
	json.NewEncoder(w).Encode("No Course Found with given Id")

}

// to inject values ---> create one course

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one course")
	w.Header().Set("Content-Type", "application/json")

	// what if: entire body is empty
	if r.Body == nil { // checks the incoming request's body
		json.NewEncoder(w).Encode("Please send some data")
	}

	// what if: data is sent like this {}
	var course Course //course --> var created for new course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.ISEmpty() {
		json.NewEncoder(w).Encode("No data inside the JSON you sent")
		return
	}
	// generate a unique id, string

	rand.Seed(time.Now().UnixNano())               // it is going to create a unique id
	course.CourseId = strconv.Itoa(rand.Intn(100)) // Itoa helps to convert int to str i.e., above int id to str
	// append new course into courses
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return
}

// perform update operation

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	// What is the course i want to update which is identified by unique courseid
	// I need all those values to be updated
	fmt.Println("Update one courses")
	w.Header().Set("Content-Type", "application/json")

	// step 1 : Grabbing Id from request
	params := mux.Vars(r)

	// step 2: loop to get the id, remove that value from the index operation, add the value again operation(add with my ID that we get in the params)

	// looping
	for index, course := range courses {
		if course.CourseId == params["id"] {
			// removing the value
			courses = append(courses[:index], courses[index+1:]...)
			// add a new value based on the id
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course) // I need to grab the id which I have in params from the body and I will be decoding the reference
			// we need to have a unique id since it is update operation
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	// TODO : send a response when ID is not found

}

// DELETE OPERATION

func DeleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete one courses")
	w.Header().Set("Content-Type", "application/json")
	// I need to check which course to delete i.e., I need to have the unique id of that course

	params := mux.Vars(r)

	// 1- loop, id, remove(index, index+1)

	// step 1 - loop

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			break // it just breaks the entire loop

		}
	}

}
