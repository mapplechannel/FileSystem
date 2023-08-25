package entity

type Person struct {
	Name   string `csv:"name" xlsx:"name"`
	Age    int    `csv:"age" xlsx:"age"`
	Gender string `csv:"gender" xlsx:"gender"`
}
