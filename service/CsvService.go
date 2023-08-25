package service

import (
	"FILEMANAGESYS/entity"
	"encoding/csv"
	"os"
	"strconv"
)

// 用于读取csv文件内容并输出
func ReadCsvFile(filename string) ([]entity.Person, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.LazyQuotes = true
	reader.FieldsPerRecord = -1
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var result []entity.Person
	for _, record := range records[1:] {
		age, err := strconv.Atoi(record[1])
		if err != nil {
			return nil, err
		}
		person := entity.Person{
			Name:   record[0],
			Age:    age,
			Gender: record[2],
		}
		result = append(result, person)
	}
	return result, nil
}
