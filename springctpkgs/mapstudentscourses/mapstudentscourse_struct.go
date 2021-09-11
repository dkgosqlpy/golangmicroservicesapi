package mapstudentscourses

import "time"

var _tb_msc = "map_students_courses"

type MapStudentCourseStruct struct {
	MscId     int       `gorm:"column:msc_id;primary_key"`
	StId      int       `gorm:"column:st_id"`
	CsId      int       `gorm:"column:cs_id"`
	CreatedDt time.Time `gorm:"column:created_dt"`
}

// MapStudentCourse validation information
type ValidatorMapStudentCourseRegistration struct {
	StId int `validate:"required,gt=0"`
	CsId int `validate:"required,gt=0"`
}

type MapStudentCourseRegistrationError struct {
	ErrData []ErrorMapStudentCourseRegistrationMessage
}

type ErrorMapStudentCourseRegistrationMessage struct {
	Name  string `json:"name"`
	Desc  string `json:"desc"`
	Error string `json:"error"`
}

func (t *MapStudentCourseStruct) TableName() string {
	return _tb_msc
}
