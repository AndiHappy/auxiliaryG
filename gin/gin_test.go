package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"runtime"
	"testing"

	"github.com/gin-gonic/gin"
)

// 模拟的是数据库
var db = make(map[string]string)

// 路由器
func setupRouter() *gin.Engine {
	// Disable Console Color
	r := gin.Default()

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// Get user value
	r.GET("/user/:name", func(c *gin.Context) {
		user := c.Params.ByName("name")
		value, ok := db[user]
		if ok {
			c.JSON(http.StatusOK, gin.H{"user": user, "value": value})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
		}
	})

	// Authorized group (uses gin.BasicAuth() middleware)
	// Same than:
	// authorized := r.Group("/")
	// authorized.Use(gin.BasicAuth(gin.Credentials{
	//	  "foo":  "bar",
	//	  "manu": "123",
	//}))
	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		"foo":  "bar", // user:foo password:bar
		"manu": "123", // user:manu password:123
	}))

	/* example curl for /admin with basicauth header
	   Zm9vOmJhcg== is base64("foo:bar")

		curl -X POST \
	  	http://localhost:8080/admin \
	  	-H 'authorization: Basic Zm9vOmJhcg==' \
	  	-H 'content-type: application/json' \
	  	-d '{"value":"bar"}'
	*/
	authorized.POST("admin", func(c *gin.Context) {
		user := c.MustGet(gin.AuthUserKey).(string)

		// Parse JSON
		var json struct {
			Value string `json:"value" binding:"required"`
		}

		if c.Bind(&json) == nil {
			db[user] = json.Value
			c.JSON(http.StatusOK, gin.H{"status": "ok"})
		}
	})

	return r
}

func TestGin(t *testing.T) {
	runtime.SetBlockProfileRate(1)     // 开启对阻塞操作的跟踪，block
	runtime.SetMutexProfileFraction(1) // 开启对锁调用的跟踪，mutex
	r := setupRouter()
	//go func() {
	//	err := http.ListenAndServe("0.0.0.0:6060", nil)
	//	if err != nil {
	//		_ = fmt.Errorf("ListenAndServe err:%+v \n", err)
	//	}
	//}()
	err := r.Run(":8080")
	if err != nil {
		_ = fmt.Errorf("start err:%+v \n", err)
		return
	}
}
