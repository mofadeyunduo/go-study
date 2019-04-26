package ormbenchmark

import "database/sql"

type Student struct {
	Id      int    `json:"id" gorm:"primary_key"`
	Name    string `json:"name"`
	Sex     int    `json:"sex"`
	Address string `json:"address"`
}

type StudentGrade struct {
	Student
	StudentId sql.NullInt64  `gorm:"column:student_id" db:"student_id"`
	GradeName sql.NullString `gorm:"column:grade_name" db:"grade_name"`
}

func (Student) TableName() string {
	return "student"
}
