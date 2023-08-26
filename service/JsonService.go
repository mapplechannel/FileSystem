package service

import (
	"FILEMANAGESYS/entity"
	"encoding/json"
	"io/ioutil"
)

// 读取Json文件内容并输出
func ReadJsonFile(filename string) ([]entity.Student, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var result []entity.Student
	err = json.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
