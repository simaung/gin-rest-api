package main

import (
	"fmt"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello World",
	})
}

func PostHomePage(c *gin.Context) {
	body := c.Request.Body
	value, err := ioutil.ReadAll(body)
	if err != nil {
		fmt.Println(err.Error())
	}
	c.JSON(200, gin.H{
		"message": string(value),
	})
}

func QueryStrings(c *gin.Context) {
	name := c.Query("name")
	age := c.Query("age")

	c.JSON(200, gin.H{
		"name": name,
		"age":  age,
	})
}

func PathParameters(c *gin.Context) {
	name := c.Param("name")
	age := c.Param("age")

	c.JSON(200, gin.H{
		"name": name,
		"age":  age,
	})
}

func main() {
	r := gin.Default()
	r.GET("/", Home)
	r.POST("/", PostHomePage)
	r.GET("/query", QueryStrings)             // /query?name=dede&age=35
	r.GET("/path/:name/:age", PathParameters) // /path/dede/35

	r.Run() // listen and serve on 0.0.0.0:8080
}
