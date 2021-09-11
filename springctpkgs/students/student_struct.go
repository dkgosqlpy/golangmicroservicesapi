package students

import "time"

var _tb_ps = "students"

type StudentStruct struct {
	StId      int       `gorm:"column:st_id;primary_key"`
	StName    string    `gorm:"column:st_name"`
	StEmail   string    `gorm:"column:st_email"`
	StPhone   string    `gorm:"column:st_phone"`
	CreatedDt time.Time `gorm:"column:created_dt"`
}

type ViewStudentStruct struct {
	StId      int       `json:"st_id"`
	StName    string    `json:"st_name"`
	StEmail   string    `json:"st_email"`
	StPhone   string    `json:"st_phone"`
	CreatedDt time.Time `json:"created_dt"`
}

// Student validation information
type ValidatorStudentRegistration struct {
	StName  string `validate:"required,min=3,max=25"`
	StEmail string `validate:"required,email,gte=3,lte=465"`
	StPhone string `validate:"required,len=10,numeric"`
}

type StudentRegistrationError struct {
	ErrData []ErrorStudentRegistrationMessage
}

type ErrorStudentRegistrationMessage struct {
	Name  string `json:"name"`
	Desc  string `json:"desc"`
	Error string `json:"error"`
}

func (t *StudentStruct) TableName() string {
	return _tb_ps
}

func (t *ViewStudentStruct) TableName() string {
	return _tb_ps
}
