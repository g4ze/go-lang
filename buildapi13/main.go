package main

// use ctrl + c to exit the build
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

//model for course
type Course struct {
	CourseId   string  `json:"id"`
	CourseName string  `json:"coursename"`
	Author     *Author `json:"author"`
}
type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// a fake database
var courses []Course

//midleware, helper
func (c *Course) ISEmpty() bool {
	return c.CourseId == "" && c.CourseName == ""
}

func main() {
	fmt.Print("We do mod gorilla stuff andAPI")
	r := mux.NewRouter()
	//seeding(making false database)
	courses = append(courses, Course{CourseId: "2", CourseName: "ReactJs", Author: &Author{Fullname: "Hitesh", Website: "lco.dev"}})
	courses = append(courses, Course{CourseId: "3", CourseName: "JAVA", Author: &Author{Fullname: "Nilay Gupta", Website: "lco.dev"}})
	//routing
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")
	//listen to a port
	log.Fatal(http.ListenAndServe(":4000", r))
}

//controllers - file

//serve home route

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>hellowww<h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type", "aapplication/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course")
	w.Header().Set("Content-Type", "application/jspn")
	//garb id form request
	params := mux.Vars(r)
	//loop thru courses find matchjomg id
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No course found with the goven id")
	return
}
func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("create one course")
	w.Header().Set("Content-Type", "application/jspn")

	//taking empty body precaution
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please some data")
	}
	//what at {}
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.ISEmpty() {
		json.NewEncoder(w).Encode("Please some data")
		return
	}

	// generate a unique id; cov them to string; append the new course
	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	return
}
func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("update one course")
	w.Header().Set("Content-Type", "application/jspn")
	// first grab the id
	params := mux.Vars(r)

	//loop, id ,remove add my id
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}

	}
	//TODO: send a respons wjen id is not found

}
func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete one course")
	w.Header().Set("Content-Type", "application/jspn")

	params := mux.Vars(r)
	//loop, id, remoce
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			break
		}
	}
}
