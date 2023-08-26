package main

import (
	"FILEMANAGESYS/route"
	"github.com/gin-gonic/gin"
)

// @title FILESYSTEMS
// @version 1.0
// @description 该项目用于读取Excel、Txt、Json、Xml、Csv类型的文件
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

func main() {
	r := gin.Default()

	route.Setup(r)

	r.Run()
}
