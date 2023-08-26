package service

import (
	"FILEMANAGESYS/entity"
	"encoding/xml"
	"io/ioutil"
)

// 用于读取xml文件内容并输出
func ReadXmlFile(filename string) ([]entity.Book, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	type Catalog struct { // 用于存储xml文件中的根元素和子元素
		XMLName  xml.Name      `xml:"catalog"`
		Date     string        `xml:"date"`
		Cost     float64       `xml:"cost"`
		Currency string        `xml:"currency"`
		Books    []entity.Book `xml:"book"`
	}

	var catalog Catalog

	err = xml.Unmarshal(data, &catalog)
	if err != nil {
		return nil, err
	}

	return catalog.Books, nil
}
