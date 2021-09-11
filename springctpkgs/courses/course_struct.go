package courses

import "time"

var _tb_ps = "courses"

type CourseListStruct struct {
	CourseList Course
}

type CourseNew struct {
	Id   int    `form:"id"`
	Name string `form:"name"`
}

type Course struct {
	Id   int
	Name string
}

type CourseStruct struct {
	ID             int       `gorm:"column:id;primary_key"`
	CourseName     string    `gorm:"column:course_name"`
	Description    string    `gorm:"column:description"`
	CourseProfName string    `gorm:"column:course_prof_name"`
	CreatedDt      time.Time `gorm:"column:created_dt"`

	//id, course_name, course_prof_name, description, status, created_dt
}

type ViewCourseStruct struct {
	ID             int       `json:"id"`
	CourseName     string    `json:"course_name"`
	Description    string    `json:"description"`
	CourseProfName string    `json:"course_prof_name"`
	CreatedDt      time.Time `json:"created_dt"`

	//id, course_name, course_prof_name, description, status, created_dt
}

// Course validation information
type ValidatorCourseRegistration struct {
	CourseName     string `validate:"required,min=3,max=65"`
	Description    string `validate:"required,gte=3,lte=465"`
	CourseProfName string `validate:"required,gte=3,lte=25"`
}

type CourseRegistrationError struct {
	ErrData []ErrorCourseRegistrationMessage
}

type ErrorCourseRegistrationMessage struct {
	Name  string `json:"name"`
	Desc  string `json:"desc"`
	Error string `json:"error"`
}

func (t *CourseStruct) TableName() string {
	return _tb_ps
}

func (t *ViewCourseStruct) TableName() string {
	return _tb_ps
}
