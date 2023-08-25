package service

import (
	"FILEMANAGESYS/util"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"golang.org/x/text/encoding/simplifiedchinese"
	"io/ioutil"
	"log"

	"golang.org/x/text/encoding/unicode"
)

type ExcelService struct {
	FileName string
}

// NewExcelService 返回要读取Excel文件的名称
func NewExcelService(fileName string) *ExcelService {
	return &ExcelService{FileName: fileName}
}

// ReadExcel 返回一个空接口，可以存储任何值
func (e *ExcelService) ReadExcel() (interface{}, error) {
	// 打开文件
	f, err := excelize.OpenFile(e.FileName)
	if err != nil {
		return nil, err
	}

	file, err := ioutil.ReadFile(e.FileName)

	fmt.Println(file)
	fmt.Println(util.IsUtf8(file))

	//获取所有的Sheet
	sheetNames := f.GetSheetMap()

	//创建一个map，value为空接口，可以接收任何类型值
	data := make(map[string]interface{})

	for _, sheetName := range sheetNames {
		//获取所有的行，返回的是一个二维数组
		rows := f.GetRows(sheetName)
		//for i, row := range rows {
		//	for j, _ := range row {
		//fmt.Println(rows[i][j])
		//fmt.Println(reflect.TypeOf(rows[i][j]))
		//b := []byte(rows[i][j])
		//isGBK := util.IsUtf8(b)
		//fmt.Println(isGBK)
		//	}
		//}
		//将Sheet作为key，每一行的值为value
		data[sheetName] = rows
	}

	return data, nil
}

// Test1 测试文件使用哪种字符集编码
func Test1() {
	// 读取Excel文件
	data, err := ioutil.ReadFile("data.csv")
	if err != nil {
		log.Fatal(err)
	}
	// 尝试用GBK解码
	gbkDecoder := simplifiedchinese.GBK.NewDecoder()
	_, err = gbkDecoder.Bytes(data)
	if err == nil {
		fmt.Println("The encoding is GBK")
		return
	}
	// 尝试用UTF-8解码
	utf8Decoder := unicode.UTF8.NewDecoder()
	_, err = utf8Decoder.Bytes(data)
	if err == nil {
		fmt.Println("The encoding is UTF-8")
		return
	}
	// 无法确定编码
	fmt.Println("The encoding is unknown")
}
