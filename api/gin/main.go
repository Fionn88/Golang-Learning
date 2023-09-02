package main

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.RedirectFixedPath = true

	router.GET("/json", returnJson)
	router.GET("/json2", returnJson2)
	router.GET("/para1", para1)
	router.GET("/para2/:input", para2)
	router.POST("/post", post)
	router.Any("/any", any)

	router.Run(":80")
}

func returnJson(c *gin.Context) {
	m := map[string]string{"status": "ok"}
	j, _ := json.Marshal(m)
	c.Data(http.StatusOK, "application/json", j)
}

func returnJson2(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"狀態": "ok",
	})
}

func para1(c *gin.Context) {
	// input := c.DefaultQuery("input", "使用者沒有任何輸入。")  // 使用者沒有輸入參數時 可設定預設值
	input := c.Query("input")
	msg := []byte("您輸入的文字為: \n" + input)                     // 純文字(text/plain)中的換行是\n，網頁格式(html)中的換行才是<br />
	c.Data(http.StatusOK, "text/plain; charset=utf-8;", msg) // 如果沒有指定文字編碼、拿掉`charset=utf-8;`的話，中文會變亂碼。

}

func para2(c *gin.Context) {
	msg := c.Param("input")
	c.String(http.StatusOK, "您輸入的文字為: \n%s", msg) // 也可使用 `c.String`返回。第二個參數為組合樣式format
	// c.String(http.StatusOK, msg)                  // 如果沒有組合樣式，可直接輸入字串
}

func post(c *gin.Context) {
	//msg := c.PostForm("input")
	msg := c.DefaultPostForm("input", "表單沒有input。") // 沒有輸入參數時 可設定預設值

	c.String(http.StatusOK, "您輸入的文字為: \n%s", msg)
}

func any(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}
