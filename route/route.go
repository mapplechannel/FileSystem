package route

import (
	"FILEMANAGESYS/entity"
	"FILEMANAGESYS/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"

	_ "FILEMANAGESYS/docs"

	swaggerFiles "github.com/swaggo/files"
	gs "github.com/swaggo/gin-swagger"
)

// Setup 配置路由
func Setup(r *gin.Engine) {
	r.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))

	r.GET("/csv", GetCsvData)
	r.GET("/excel/:filename", ExcelHandler) //over
	r.GET("/xml", GetXmlData)
	r.GET("/json", GetJsonData)
	r.GET("/txt", GetTxtData)

	r.GET("/ftp/:path", FileInfoHandler)
	r.GET("/catalog", GetFileCatalogHandler)
}

// GetCsvData @Tags 获取CSV数据
// @Summary 获取CSV数据
// @Description 指定CSV文件获取数据
// @Produce json
// @Success 200 {string} string "成功"
// @Failure 401 {string} string "失败"
// @Router /csv [get]
func GetCsvData(c *gin.Context) {
	data, err := service.ReadCsvFile("data.csv")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, data)
	}
}

// ExcelHandler @Tags 获取excel数据
// @Summary 获取excel数据
// @Description excel文件获取数据
// @Produce json
// @Success 200 {string} string "成功"
// @Failure 401 {string} string "失败"
// @Router /excel/ [get]
func ExcelHandler(c *gin.Context) {
	fileName := c.Param("filename")

	//创建文件对象
	excelService := service.NewExcelService(fileName)

	data, err := excelService.ReadExcel()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    data,
	})
}

// GetXmlData @Tags 获取xml数据
// @Summary 获取xml数据
// @Description 指定xml文件获取数据
// @Produce json
// @Success 200 {string} string "成功"
// @Failure 401 {string} string "失败"
// @Router /xml [get]
func GetXmlData(c *gin.Context) {
	data, err := service.ReadXmlFile("data.xml")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, data)
	}
}

// GetJsonData @Tags 获取json数据
// @Summary 获取json数据
// @Description 指定json文件获取数据
// @Produce json
// @Success 200 {string} string "成功"
// @Failure 401 {string} string "失败"
// @Router /json [get]
func GetJsonData(c *gin.Context) {
	data, err := service.ReadJsonFile("data.json")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, data)
	}
}

// GetTxtData @Tags 获取txt数据
// @Summary 获取txt数据
// @Description 指定txt文件获取数据
// @Produce json
// @Success 200 {string} string "成功"
// @Failure 401 {string} string "失败"
// @Router /txt [get]
func GetTxtData(c *gin.Context) {
	data, err := service.ReadTxtFile("data.txt")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, data)
	}
}

// FileInfoHandler @Tags 获取FTP服务器文件信息
// @Summary 获取FTP服务器文件信息
// @Description 获取FTP服务器文件信息
// @Produce json
// @Success 200 {string} string "成功"
// @Failure 401 {string} string "失败"
// @Router /ftp/ [get]
func FileInfoHandler(c *gin.Context) {
	// 解析路径
	path := c.Param("path")

	// 定义FTP
	ftpConfig := entity.FTPConfig{
		Host: "127.0.0.1",
		Port: "21",
		User: "test",
		Pass: "123456",
	}

	// 连接FTP服务器
	client, err := service.ConnectFTP(ftpConfig)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer client.Quit()

	// 获取文件信息
	fileInfo, err := service.GetFileInfo(client, path)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 返回文件信息
	c.JSON(http.StatusOK, fileInfo)
}

// GetFileCatalogHandler @Tags 获取指定文件夹下的所有文件目录
// @Summary 获取指定文件夹下的所有文件目录
// @Description 获取指定文件夹下的所有文件目录
// @Produce json
// @Success 200 {string} string "成功"
// @Failure 401 {string} string "失败"
// @Router /catalog [get]
func GetFileCatalogHandler(c *gin.Context) {
	dir := c.Query("dir")
	// 如果dir为空，则默认获取当前项目下的文件目录
	if dir == "" {
		dir = "."
	}

	names, err := service.ListDir(dir)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"dir": dir, "names": names}) // Return the result as a JSON object
}
