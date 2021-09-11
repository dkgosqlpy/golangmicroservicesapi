package students

var _tb_view = "view_students"

type StudentViewStruct struct {
	StId            int    `json:"st_id"`
	StName          string `json:"st_name"`
	StEmail         string `json:"st_email"`
	StPhone         string `json:"st_phone"`
	EnrolledCourses string `json:"enrolled_courses"`
}

func (t *StudentViewStruct) TableName() string {
	return _tb_view
}
