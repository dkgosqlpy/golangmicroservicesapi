package mapstudentscourses

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"springct.com/databases"

	"github.com/gin-gonic/gin"
)

func GetTest(c *gin.Context) {
	var dbc = databases.InitDB()
	defer dbc.DB().Close()

	fmt.Println("In GetTest MapStudentCourses")
	c.JSON(200, gin.H{
		"message": "In GetTest MapStudentCourses",
	})
}

//get mapped course by id
func (c MapStudentCourseStruct) GetMappedCourseById(csid int) []byte {
	var dbc = databases.InitDB()
	defer dbc.DB().Close()

	var tmpData []MapStudentCourseStruct
	var selectFields = " `msc_id`, `st_id`, `cs_id`, `enrolled_dt`, `status`, `created_dt`"
	dbc.Debug().Select(selectFields).Where("cs_id = ?", csid).Find(&tmpData)

	_tmpData, _ := json.Marshal(tmpData)
	return _tmpData
}

func AddMapStudentCourse(c *gin.Context) {
	var dbc = databases.InitDB()
	defer dbc.DB().Close()
	fmt.Println("In AddMapStudentCourse MapStudentCourses")
	var addMapStudentCourse MapStudentCourseStruct
	var defValidateStruct ValidatorMapStudentCourseRegistration
	v1, _ := strconv.Atoi(c.PostForm("stid"))
	v2, _ := strconv.Atoi(c.PostForm("csid"))

	// Add
	addMapStudentCourse.StId = v1
	addMapStudentCourse.CsId = v2
	addMapStudentCourse.CreatedDt = time.Now()

	// Validator
	defValidateStruct.StId = v1
	defValidateStruct.CsId = v2

	validationErrStatus := ValidateMapStudentCourse(defValidateStruct, c)
	fmt.Println("validationErrStatus : ", validationErrStatus)
	if validationErrStatus == true { // validation false
		if err := dbc.Debug().Save(&addMapStudentCourse).Error; err != nil {
			_errData := err.Error()
			c.JSON(200, gin.H{
				"message": "Register error",
				"data":    _errData,
			})
			return
		}

		c.JSON(200, gin.H{
			"message": "Register",
			"data":    addMapStudentCourse,
		})
	}
}
