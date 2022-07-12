package main

import (
	"fmt"
	"time"
)

//func double

func main() {
	x := 5
	y := func(n int) int {
		return n * 2
	}(x)
	fmt.Println(y)
	c := make(chan int)
	go func() {
		fmt.Println("Starting Function")
		time.Sleep(5 * time.Second)
		fmt.Println("End")
		c <- 1
	}()
	<-c
}
