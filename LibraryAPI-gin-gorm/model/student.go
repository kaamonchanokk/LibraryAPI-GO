package model

type Student struct {
	STUDENT_ID int `gorm:"primaryKey"  form:"STUDENT_ID" json:"studentId"`
	STUDENT_CODE *string `form:"STUDENT_CODE" json:"studentCode"`
	STUDENT_NAME *string `form:"STUDENT_NAME" json:"studentName"`
	STUDENT_YEAR *int `form:"STUDENT_YEAR" json:"studentYear"`
}

func (student *Student) TableName() string { 
	return "student"
}

type StudentResponse struct {
	Status int `json:"status"`
	Message string `json:"message"`
	Data    []Student
}