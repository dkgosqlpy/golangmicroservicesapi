module springct.com/springctapp

go 1.16

replace springct.com/mapstudentscourses => ../springctpkgs/mapstudentscourses

replace springct.com/databases => ../springctpkgs/databases

replace springct.com/courses => ../springctpkgs/courses

replace springct.com/students => ../springctpkgs/students

require (
	github.com/gin-gonic/gin v1.7.4
	github.com/jinzhu/gorm v1.9.16 // indirect
	springct.com/courses v0.0.0
	springct.com/databases v0.0.0
	springct.com/mapstudentscourses v0.0.0
	springct.com/students v0.0.0
)
