package main

import (
	"gee/gee"
	"log"
	"net/http"
	"time"
)

func main() {
	r := gee.Default()
	r.Static("/assets", "./static")
	r.GET("/", func(c *gee.Context) {
		c.String(http.StatusOK, "gee ok")
	})
	r.GET("/panic", func(c *gee.Context) {
		var ls []string
		log.Println(ls[1])
	})

	v2 := r.Group("/v2")
	v2.Use(onlyForV2()) // v2 group middleware
	{
		v2.GET("/hello/:name", func(c *gee.Context) {
			// expect /hello/geektutu
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
	}

	r.Run(":9000")
}

func onlyForV2() gee.HandlerFunc {
	return func(c *gee.Context) {
		// Start timer
		t := time.Now()
		// if a server error occurred
		c.String(500, "Internal Server Error")
		// Calculate resolution time
		log.Printf("[%d] %s in %v for group v2", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}