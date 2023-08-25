package entity

// Student 结构体，用于存储JSON文件中的数据
type Student struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}
