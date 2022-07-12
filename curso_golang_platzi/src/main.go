package main

import (
	"fmt"
	"time"
)

func main() {
	// Map
	var m = make(map[string]int)
	m["key"] = 1
	m["key2"] = 2
	fmt.Println("Map")
	fmt.Println(m)
	for index, value := range m {
		fmt.Println(`el index del map es`, index)
		fmt.Println(`el valor del map es `, value)
	}
	var s = []int{1, 2, 3}
	fmt.Println("Slice")
	fmt.Println(s)
	for index, value := range s {
		fmt.Println(`el index del slice es`, index)
		fmt.Println(`el valor del slice es `, value)
	}
	s = append(s, 4)
	fmt.Println("New Slice")
	fmt.Println(s)
	var c = make(chan int)
	go doSomthing(c)
	<-c
	fmt.Println(c)
	g := 25
	h := &g
	fmt.Println(g)
	fmt.Println(h)
	fmt.Println(*h)

}

func doSomthing(c chan int) {
	time.Sleep(3 * time.Second)
	fmt.Println("done")
	c <- 1
}
