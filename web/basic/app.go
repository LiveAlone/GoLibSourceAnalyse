package basic

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LocalServerHandler struct {
}

func (engine *LocalServerHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Println(req.Header)
	w.Write([]byte{'h', 'e', 'l', 'l', '0'})
}

func SelfDefHandler() {
	// 自定义response 定义方式
	http.ListenAndServe(":8080", &LocalServerHandler{})
}

func AppRequestDemo() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World!")
	})

	r.GET("/error", func(context *gin.Context) {
		zero := 0
		fmt.Println(1 / zero)
		context.String(http.StatusOK, "pong")
	})

	r.GET("user/:name", func(context *gin.Context) {
		name, ok := context.Params.Get("name")
		if ok {
			context.String(http.StatusOK, fmt.Sprintf("hello %s", name))
			return
		}
		context.String(http.StatusOK, "none one")
	})
	//
	//authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
	//	"yao": "yao",
	//}))
	//
	//authorized.POST("admin", func(context *gin.Context) {
	//	user := context.MustGet(gin.AuthUserKey).(string)
	//
	//	// Parse JSON
	//	var json struct {
	//		// 绑定对应的json Value config
	//		Value string `json:"value" binding:"required"`
	//	}
	//
	//	if context.Bind(&json) == nil {
	//		fmt.Println(json.Value)
	//		db[user] = json.Value
	//		context.JSON(http.StatusOK, gin.H{"status": "ok"})
	//	}
	//})

	r.Run(":8080")
	http.Handle("/", r)
}

func StartAppFunc() {
	r := gin.New()
	//r.GET("/", func(c *gin.Context) {
	//	c.String(http.StatusOK, "Hello World!")
	//})
	//r.GET("/ping", func(c *gin.Context) {
	//	c.String(http.StatusOK, "pong")
	//})

	rt := r.Group("/root")
	rt.Use(None)
	{
		rt.GET("/pong", func(c *gin.Context) {
			c.String(http.StatusOK, "Hello World!")
		})
		rt.GET("/ping", func(c *gin.Context) {
			c.String(http.StatusOK, "pong")
		})
	}

	r.Run(":8080")
	http.Handle("/", r)
}

func None(c *gin.Context) {
	return
}
