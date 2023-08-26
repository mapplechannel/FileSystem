package service

import (
	"io/ioutil"
	"path/filepath"
)

// ListDir 返回当前目录下的所有文件目录
func ListDir(dir string) ([]string, error) {
	// 获取文件对象
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	// 定义切片去存储路径
	var names []string
	for _, file := range files {
		// 获取文件名称
		name := file.Name()
		// 判断该文件是否是一个路径，即文件夹
		if file.IsDir() {
			// 是路径，在名称后面+/，进入下一级目录
			name += "/"
			// Join的功能是将两个字符串使用“/”拼接起来
			// 开启递归，获取下一级的目录
			surnames, err := ListDir(filepath.Join(dir, name))
			if err != nil {
				return nil, err
			}
			// 递归到最后一层时，向切片中添加路径，然后逐级向上回溯
			for _, surname := range surnames {
				names = append(names, name+surname)
			}
		} else {
			// 如果不是文件夹，直接添加目录
			names = append(names, name)
		}
	}
	return names, nil
}
