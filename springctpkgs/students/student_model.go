package students

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"springct.com/databases"

	"github.com/gin-gonic/gin"
)

func GetTest(c *gin.Context) {
	var dbc = databases.InitDB()
	_ = dbc

	fmt.Println("In GetTest Students")
	c.JSON(200, gin.H{
		"message": "In GetTest Students",
	})

}

func AddStudent(c *gin.Context) {
	var dbc = databases.InitDB()
	defer dbc.DB().Close()

	fmt.Println("In AddStudent student")
	var addstudent StudentStruct
	var defValidateStruct ValidatorStudentRegistration
	v1 := strings.TrimSpace(c.PostForm("name"))
	v2 := strings.TrimSpace(c.PostForm("email"))
	v3 := strings.TrimSpace(c.PostForm("phone"))

	// Add
	addstudent.StName = v1
	addstudent.StEmail = v2
	addstudent.StPhone = v3
	addstudent.CreatedDt = time.Now()

	// Validator
	defValidateStruct.StName = v1
	defValidateStruct.StEmail = v2
	defValidateStruct.StPhone = v3

	validationErrStatus := ValidateStudent(defValidateStruct, c)
	fmt.Println("validationErrStatus : ", validationErrStatus)
	if validationErrStatus == true { // validation false
		if err := dbc.Debug().Save(&addstudent).Error; err != nil {
			_errData := err.Error()
			c.JSON(200, gin.H{
				"message": "Register error",
				"data":    _errData,
			})
			return
		}

		c.JSON(200, gin.H{
			"message": "Register",
			"data":    addstudent,
		})
	}
}

//get GetStudentList
func getStudentList() []byte {
	var dbc = databases.InitDB()
	defer dbc.DB().Close()

	var tmpData []ViewStudentStruct
	var selectFields = "`st_id`, `st_name`, `st_email`, `st_phone`, `status`, `created_dt`"
	dbc.Debug().Select(selectFields).Find(&tmpData)

	_tmpData, _ := json.Marshal(tmpData)
	return _tmpData
}

func GetStudentList(c *gin.Context) {
	fmt.Println("In GetStudentList Students")
	var databyte = getStudentList()
	var Studentlist []ViewStudentStruct
	json.Unmarshal(databyte, &Studentlist)
	c.JSON(200, gin.H{
		"message": "In GetStudentList Students",
		"data":    Studentlist,
	})

}
