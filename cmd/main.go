package main

import (
	"docker_learn/nosql"
	"fmt"
)

func main() {
	redisOps := nosql.NewRedisClient()
	redisOps.Set("name", "cwb")
	str := ""
	redisOps.Get("name", &str)
	fmt.Println("value: " + str)
}
