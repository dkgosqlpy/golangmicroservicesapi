package students

import (
	"encoding/json"
	"fmt"

	"springct.com/databases"

	"github.com/gin-gonic/gin"
)

//get getStudentViewList
func getStudentViewList() []byte {
	var dbc = databases.InitDB()
	var tmpData []StudentViewStruct
	var selectFields = "`st_id`, `st_name`, `st_email`, `st_phone`, `enrolled_courses`"
	dbc.Debug().Select(selectFields).Find(&tmpData)

	_tmpData, _ := json.Marshal(tmpData)
	return _tmpData
}

func GetStudentViewList(c *gin.Context) {
	fmt.Println("In GetStudentViewList Students")
	var databyte = getStudentViewList()
	var studentlist []StudentViewStruct
	json.Unmarshal(databyte, &studentlist)
	c.JSON(200, gin.H{
		"message": "In GetStudentViewList Students",
		"data":    studentlist,
	})

}
