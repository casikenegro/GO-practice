package main

import (
	"fmt"
	"sync"
	"time"
)

/*
traffic light.

this uses channels and waint groups to 1. execute only 2 doSmth() func
at a time and 2. be able to wait for all of them.

in order of execution it'll:
c := [][] -- two free spaces
c := [routine][] -- one free space
c := [rountine][routine] -- all occupied
c := [][routine] -- one free space
*/

func doSmth(i int, wg *sync.WaitGroup, c chan int) {
	defer wg.Done()
	fmt.Printf("Id: %d -> started...\n", i)
	time.Sleep(time.Second * 4)
	fmt.Printf("Id: %d -> finished...\n", i)

	<- c // frees the space for new routines
}

func main() {
	c := make(chan int, 2) // creates a buffered channel with a capacity of two
	var wg sync.WaitGroup // creates wait group

	for i := 0; i < 10; i++ {
		c <- 1 // alocate a new "instance" in the free space
		wg.Add(1) // adds to the wait group
		go doSmth(i, &wg, c)
	}

	wg.Wait()
}
