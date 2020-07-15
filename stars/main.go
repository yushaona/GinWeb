package main

import (
	"fmt"
	"net/http"

	"GinWeb/stars/module"

	"github.com/gin-gonic/gin"
)

type ReadInter interface {
	Read()
}

type WriteInter interface {
	Write()
}

type Combine interface {
	ReadInter
	WriteInter
}

type Man struct {
	inter Combine
}

type TheStruct struct {
	Combine
}

type CombinStruct struct {
}

func (this *CombinStruct) Read() {
	fmt.Println("Read")
}

func (this *CombinStruct) Write() {
	fmt.Println("write")
}

type s1 struct {
	a string
	s2
}

func (t *s1) hello1() {
	fmt.Println("s1" + t.a)
}

type s2 struct {
	a string
}

func (t s2) hello() {
	fmt.Println("s2" + t.a)
}

// func (t s2) hello(a string) {

// }

func main() {

	// s := s1{a: "10"}
	// s.hello()
	// s.hello1()

	// s1 := &s1{a: "100"}
	// s1.hello()
	// s1.hello1()

	// return
	// man := &Man{
	// 	inter: &CombinStruct{},
	// }
	// man.inter.Read()

	// man1 := &TheStruct{
	// 	&CombinStruct{},
	// }
	// man1.Read()
	// return

	r := gin.Default()

	r.Use(gin.BasicAuth(gin.Accounts{
		"admin": "123456",
	}))

	r.GET("/jsonp", func(c *gin.Context) {
		c.JSONP(200, gin.H{"wechat": "flysnow_org"})
	})
	//查询参数

	r.GET("/func", func(c *gin.Context) {

		_ = c.PostForm("param")
		q := c.Query("param")
		c.JSON(200, map[string]interface{}{
			"funcid": q,
		})
	})
	//返回重定向
	r.NoRoute(func(c *gin.Context) {
		//id := c.Param("id")
		//c.String(404, "f u")
		c.Redirect(302, "/users/3333")
	})

	//路由参数
	r.GET("/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.String(200, "The user id is  %s", id)
	})

	//自定的方式添加路由
	r.Handle("GET", "/users", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"Blog":   "www.flysnow.org",
			"wechat": "flysnow_org",
		})

	})

	//内部调用的也是Handle
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"Blog":   "www.flysnow.org",
			"wechat": "flysnow_org",
		})
		//c.XML()
	})

	r.GET("/xml", func(c *gin.Context) {
		allUsers := []User{{ID: 123, Name: "张三", Age: 20}, {ID: 456, Name: "李四", Age: 25}}
		//c.XML(200, gin.H{"users": allUsers})
		c.XML(200, Root{Data: allUsers})
	})

	r.LoadHTMLGlob("html/*/*")
	r.GET("/html", func(c *gin.Context) {
		c.HTML(200, "index.html", "flysnow_org")
	})

	r.GET("/protobuf", func(c *gin.Context) {
		data := &module.User{
			Name: "张三",
			Age:  20,
		}
		c.ProtoBuf(http.StatusOK, data)
	})

	r.GET("/protobuf_map", func(c *gin.Context) {
		pro := make(map[string]interface{})
		pro["nash"] = 1
		pro["ss"] = "sdf"

		c.ProtoBuf(http.StatusOK, pro)
	})

	r.Run(":8080")

}

type Root struct {
	Data interface{} `xml:data`
}

type User struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
	Age  int    `xml:"age"`
}
