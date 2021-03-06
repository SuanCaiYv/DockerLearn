package main

import (
	"docker_learn/nosql"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	redisOps := nosql.NewRedisClient()
	fmt.Println("redis ops get")
	server.POST("/t1", func(context *gin.Context) {
		fmt.Println("api1")
		v := context.Query("name")
		fmt.Println(v)
		err := redisOps.Set("name", v)
		fmt.Println(err)
		context.JSON(200, "ok")
	})
	server.GET("/t2/:name", func(context *gin.Context) {
		name := context.Param("name")
		v := ""
		redisOps.Get(name, &v)
		context.JSONP(200, v)
	})
	server.Run(":8190")
}
