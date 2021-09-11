package courses

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// use a single instance of Validate, it caches struct info
var validate *validator.Validate

/*Form validation Type */
func ValidateCourse(validateCourse ValidatorCourseRegistration, c *gin.Context) bool {
	fmt.Println("ValidatorCourse: ", validateCourse)
	validate = validator.New()
	err := validate.Struct(validateCourse)
	fmt.Println("err: ", err)

	if err != nil {
		errorMessage := []ErrorCourseRegistrationMessage{}
		CourseRegistrationErrors := CourseRegistrationError{errorMessage}
		for _, err := range err.(validator.ValidationErrors) {
			_keyName := err.Field()
			_errKey := fmt.Sprintf("Course Validation Error")

			_errVal := fmt.Sprintf("%s, %s, %s, %s", err.Field(), err.ActualTag(), err.Param(), err.Value())
			_error := ErrorCourseRegistrationMessage{Name: _keyName, Desc: _errKey, Error: _errVal}
			CourseRegistrationErrors.AddError(_error)
		}

		c.JSON(http.StatusOK, gin.H{"message": "Error found", "content": CourseRegistrationErrors, "errormsg": "validation error"})
		return false
	}
	fmt.Println("ValidateCourse Out : ")
	return true
}

func (distregi *CourseRegistrationError) AddError(errormsg ErrorCourseRegistrationMessage) []ErrorCourseRegistrationMessage {
	distregi.ErrData = append(distregi.ErrData, errormsg)
	return distregi.ErrData
}
