package main

import (
	"springct.com/courses"
	"springct.com/databases"
	"springct.com/mapstudentscourses"
	"springct.com/students"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	//Here can be grouped in individual model
	router.GET("/", courses.GetTest)
	router.GET("/test", courses.GetTest)
	router.GET("/courses", courses.GetCourses)
	router.POST("/courses", courses.GetCourses)
	router.GET("/courselist", courses.GetCourseList)
	router.POST("/addcourse", courses.AddCourse)
	router.POST("/newcourse", courses.NewCourse)
	router.DELETE("/deletecourse", courses.DeleteCourse)

	router.POST("/addstudent", students.AddStudent)
	router.GET("/studentlist", students.GetStudentList)
	router.GET("/conn", databases.GetDbConnect)
	router.GET("/viewstudentlist", students.GetStudentViewList)

	router.POST("/enroll", mapstudentscourses.AddMapStudentCourse)

	router.Run("localhost:8088")
}
