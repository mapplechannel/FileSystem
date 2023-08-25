package test

import (
	"FILEMANAGESYS/entity"
	"FILEMANAGESYS/route"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestJsonHandler(t *testing.T) {
	// 创建一个Gin引擎
	router := gin.Default()

	// 注册路由
	router.POST("/getJson", route.JsonHandler)

	// 创建一个Student实例
	stu := entity.Student{
		Name:  "Alice",
		Age:   18,
		Email: "alice@example.com",
	}

	// 将Student实例转换为JSON字节
	stuBytes, _ := json.Marshal(stu)

	// 创建一个POST请求，请求体为JSON字节
	req, _ := http.NewRequest("POST", "/getJson", bytes.NewReader(stuBytes))

	// 使用httptest包创建一个响应记录器
	w := httptest.NewRecorder()

	// 将请求和响应记录器传递给Gin引擎
	router.ServeHTTP(w, req)

	// 检查响应状态码是否为200
	assert.Equal(t, 200, w.Code)

	// 检查响应体是否与请求体相同
	assert.Equal(t, string(stuBytes), w.Body.String())
}
