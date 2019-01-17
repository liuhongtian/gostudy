package main

import (
	"fmt"
	"github.com/liuhongtian/golibs/redisutil"
	"github.com/liuhongtian/golibs/stringutil"
)

func main() {
	fmt.Println("Hello, GO!")
	fmt.Println(stringutil.Revers("Hello, GO!"))

	fmt.Println("redis:")

	value := redisutil.GetFromRedis("liu")
	fmt.Println(value)

}
