package students

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// use a single instance of Validate, it caches struct info
var validate *validator.Validate

/*Form validation Type */
func ValidateStudent(validateStudent ValidatorStudentRegistration, c *gin.Context) bool {
	fmt.Println("ValidatorStudent: ", validateStudent)
	validate = validator.New()
	err := validate.Struct(validateStudent)
	fmt.Println("err: ", err)

	if err != nil {
		errorMessage := []ErrorStudentRegistrationMessage{}
		StudentRegistrationErrors := StudentRegistrationError{errorMessage}
		for _, err := range err.(validator.ValidationErrors) {
			_keyName := err.Field()
			_errKey := fmt.Sprintf("Student Validation Error")

			_errVal := fmt.Sprintf("%s, %s, %s, %s", err.Field(), err.ActualTag(), err.Param(), err.Value())
			_error := ErrorStudentRegistrationMessage{Name: _keyName, Desc: _errKey, Error: _errVal}
			StudentRegistrationErrors.AddError(_error)
		}

		c.JSON(http.StatusOK, gin.H{"message": "Error found", "content": StudentRegistrationErrors, "errormsg": "validation error"})
		return false
	}
	fmt.Println("ValidateStudent Out : ")
	return true
}

func (distregi *StudentRegistrationError) AddError(errormsg ErrorStudentRegistrationMessage) []ErrorStudentRegistrationMessage {
	distregi.ErrData = append(distregi.ErrData, errormsg)
	return distregi.ErrData
}
