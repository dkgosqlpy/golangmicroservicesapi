package mapstudentscourses

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// use a single instance of Validate, it caches struct info
var validate *validator.Validate

/*Form validation Type */
func ValidateMapStudentCourse(validateMapStudentCourse ValidatorMapStudentCourseRegistration, c *gin.Context) bool {
	fmt.Println("ValidatorMapStudentCourse: ", validateMapStudentCourse)
	validate = validator.New()
	err := validate.Struct(validateMapStudentCourse)
	fmt.Println("err: ", err)

	if err != nil {
		errorMessage := []ErrorMapStudentCourseRegistrationMessage{}
		MapStudentCourseRegistrationErrors := MapStudentCourseRegistrationError{errorMessage}
		for _, err := range err.(validator.ValidationErrors) {
			_keyName := err.Field()
			_errKey := fmt.Sprintf("MapStudentCourse Validation Error")

			_errVal := fmt.Sprintf("%s, %s, %s, %s", err.Field(), err.ActualTag(), err.Param(), err.Value())
			_error := ErrorMapStudentCourseRegistrationMessage{Name: _keyName, Desc: _errKey, Error: _errVal}
			MapStudentCourseRegistrationErrors.AddError(_error)
		}

		c.JSON(http.StatusOK, gin.H{"message": "Error found", "content": MapStudentCourseRegistrationErrors, "errormsg": "validation error"})
		return false
	}
	fmt.Println("ValidateMapStudentCourse Out : ")
	return true
}

func (distregi *MapStudentCourseRegistrationError) AddError(errormsg ErrorMapStudentCourseRegistrationMessage) []ErrorMapStudentCourseRegistrationMessage {
	distregi.ErrData = append(distregi.ErrData, errormsg)
	return distregi.ErrData
}
