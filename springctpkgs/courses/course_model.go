package courses

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"springct.com/databases"
	"springct.com/mapstudentscourses"

	"github.com/gin-gonic/gin"
)

func GetTest(c *gin.Context) {
	var dbc = databases.InitDB()
	defer dbc.DB().Close()

	fmt.Println("In GetTest courses")
	c.JSON(200, gin.H{
		"message": "In GetTest courses",
	})

}
func NewCourse(c *gin.Context) {
	var dbc = databases.InitDB()
	defer dbc.DB().Close()
	fmt.Println("In newcourse courses")
	v1 := strings.TrimSpace(c.PostForm("getkey"))
	q1, _ := c.GetQuery("q")
	q2, _ := c.GetQuery("table")
	fmt.Println("V1 data:", v1)

	fmt.Println("q1 data:", q1)

	fmt.Println("q2 data:", q2)

	var newcourse []Course
	text := "[{\"Id\": 100, \"Name\": \"Go " + v1 + ", q1:" + q1 + ", q2:" + q2 + "\"}]" //, {\"Id\": 200, \"Name\": \"Java " + v1 + "\"}
	bytes := []byte(text)
	json.Unmarshal(bytes, &newcourse)

	c.JSON(200, gin.H{
		"message": "NewCourse",
		"data":    newcourse,
	})
}

func AddCourse(c *gin.Context) {
	var dbc = databases.InitDB()
	defer dbc.DB().Close()
	fmt.Println("In AddCourse courses")
	var addcourse CourseStruct
	var defValidateStruct ValidatorCourseRegistration
	v1 := strings.TrimSpace(c.PostForm("name"))
	v2 := strings.TrimSpace(c.PostForm("desc"))
	v3 := strings.TrimSpace(c.PostForm("profname"))

	// Add
	addcourse.CourseName = v1
	addcourse.Description = v2
	addcourse.CourseProfName = v3
	addcourse.CreatedDt = time.Now()

	// Validator
	defValidateStruct.CourseName = v1
	defValidateStruct.Description = v2
	defValidateStruct.CourseProfName = v3

	validationErrStatus := ValidateCourse(defValidateStruct, c)
	fmt.Println("validationErrStatus : ", validationErrStatus)
	if validationErrStatus == true { // validation false
		if err := dbc.Debug().Save(&addcourse).Error; err != nil {
			_errData := err.Error()
			c.JSON(200, gin.H{
				"message": "Register error",
				"data":    _errData,
			})
			return
		}

		c.JSON(200, gin.H{
			"message": "Register",
			"data":    addcourse,
		})
	}
}

//get courselist
func getCourseList() []byte {
	var dbc = databases.InitDB()
	defer dbc.DB().Close()

	var tmpData []ViewCourseStruct
	var selectFields = "`id`, `course_name`, `course_prof_name`, `description`, `status`, `created_dt`"
	dbc.Debug().Select(selectFields).Find(&tmpData)

	_tmpData, _ := json.Marshal(tmpData)
	return _tmpData
}

func GetCourses(c *gin.Context) {
	fmt.Println("In GetCourses courses")
	var coursetest CourseNew

	if c.BindQuery(&coursetest) == nil {
		log.Println("====== Only Bind Query String ======")
		log.Println(coursetest.Id)
		log.Println(coursetest.Name)
		fmt.Println("====== Only Bind Query String ======")
		fmt.Println(coursetest.Name)
	}

	//# only bind query
	// '''
	// QUERY:
	// ------
	// http://localhost:8088/courses?timestamp=124245&number=145

	// curl -X GET "localhost:8088/courses?id=9000&name=cic9000"

	// # only bind query string, ignore form data
	// curl -X POST "localhost:8088/courses?id=9000&name=cic9000" --data 'id=ignore&name=ignore' -H "Content-Type:application/x-www-form-urlencoded"
	// '''

	timestamp, _ := c.GetQuery("timestamp")
	number, _ := c.GetQuery("number")
	byt := []byte(`{"time":` + timestamp + `,"num":` + number + `,"timev2":` + timestamp + `,"numv2":` + number + `}`)
	var dat map[string]interface{}
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	//dat := "Params :" + "time:" + timestamp + ", num:" + number

	var databyte = getCourseList()
	var courselist []ViewCourseStruct
	json.Unmarshal(databyte, &courselist)
	c.JSON(200, gin.H{
		"allparams": dat,
		"message":   "In GetCourseList courses",
		"data":      courselist,
	})
}

func GetCourseList(c *gin.Context) {
	fmt.Println("In GetCourseList courses")
	var databyte = getCourseList()
	var courselist []ViewCourseStruct
	json.Unmarshal(databyte, &courselist)
	c.JSON(200, gin.H{
		"message": "In GetCourseList courses",
		"data":    courselist,
	})
}

//get course by id
func (c CourseStruct) getCourseById(csid int) CourseStruct {
	var dbc = databases.InitDB()
	defer dbc.DB().Close()

	var tmpData CourseStruct
	var selectFields = "`id`, `course_name`, `course_prof_name`, `description`, `status`, `created_dt`"
	dbc.Debug().Select(selectFields).Where("id = ?", csid).First(&tmpData)

	fmt.Println("HERE getCourseById :", tmpData)
	//_tmpData, _ := json.Marshal(tmpData)
	return tmpData
}

//Delete course by id
func (c CourseStruct) deleteCourseById(coursename CourseStruct) string {
	var dbc = databases.InitDB()
	defer dbc.DB().Close()

	if dbc.Debug().Delete(coursename).Error != nil {
		fmt.Println("Error:")
		return "Failed"
	}

	return "Success"
}

func DeleteCourse(c *gin.Context) {
	var dbc = databases.InitDB()
	defer dbc.DB().Close()
	fmt.Println("In AddMapStudentCourse MapStudentCourses")

	v1, _ := strconv.Atoi(c.PostForm("csid"))
	if v1 == 0 {
		c.JSON(200, gin.H{
			"message": "In DeleteCourse courses",
			"data":    "Invalid course id",
		})
		return
	}
	var mapStudentCourseModel mapstudentscourses.MapStudentCourseStruct

	var mapdatabyte = mapStudentCourseModel.GetMappedCourseById(v1)
	var mappedcourse []mapstudentscourses.MapStudentCourseStruct
	json.Unmarshal(mapdatabyte, &mappedcourse)

	mappedCourseLen := len(mappedcourse)
	if mappedCourseLen == 0 {
		var courseModel CourseStruct
		courseModel = courseModel.getCourseById(v1)
		output := ""
		if courseModel.ID > 0 {
			output = courseModel.deleteCourseById(courseModel)
		}

		c.JSON(200, gin.H{
			"message": "In GetCourseList courses",
			"data":    output,
		})
	} else {
		var courseModel CourseStruct
		c.JSON(200, gin.H{
			"message": "You Can Not delete Allocated course",
			"data":    courseModel,
		})
	}
}
