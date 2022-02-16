package models

type Student struct {
	Id  string `json:"id";gorm:"primaryKey"`
	Number string `json:"number"`
}
//设置查询表名为student，否则默认为students
func (Student) TableName() string {
	return "student"
}
