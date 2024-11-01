package model

type StudentInfo struct {
	Id   uint   `json:"id"`   // Student ID
	Name string `json:"name"` // Student Name
	Age  uint   `json:"age"`  // Student Age
}
