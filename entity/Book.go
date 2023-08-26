package entity

type Book struct {
	Title  string  `xml:"title"`
	Author string  `xml:"author"`
	Price  float64 `xml:"price"`
}
