package service

import (
	"FILEMANAGESYS/entity"
	"bufio"
	"os"
)

// 读取Txt文件内容并输出
func ReadTxtFile(filename string) ([]entity.Article, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var result []entity.Article
	var article entity.Article
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			result = append(result, article)
			article = entity.Article{}
		} else if article.Title == "" {
			article.Title = line
		} else {
			article.Content += line + "\n"
		}
	}
	if article.Title != "" {
		result = append(result, article)
	}
	return result, nil
}
