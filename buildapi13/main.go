package main

import "fmt"

//model for course
type Course struct {
	CourseId   string  `json:"id"`
	CourseName string  `json:"coursename"`
	Author     *Author `json:"author"`
}
type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"ebsite"`
}

// a fake database
var course []Course

//midleware, helper
func (c *Course) ISEmpty() bool {
	return c.CourseId == "" && c.CourseName == ""
}

func main() {
	fmt.Print("We do mod gorilla stuff")
}
