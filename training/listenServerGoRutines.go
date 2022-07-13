package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	init := time.Now()
	channel := make(chan string)
	servers := []string{
		"https://google.com",
		"https://platzi.com",
		"https://github.com",
	}
	for {
		for _, value := range servers {
			go checkServer(value, channel)
		}
		time.Sleep(1 * time.Second)
		fmt.Println(<-channel)
	}

	final := time.Since(init)
	fmt.Println(final)
}

func checkServer(server string, channel chan string) {
	_, err := http.Get(server)
	if err != nil {
		channel <- server + "disabled"
	}
	channel <- server + "available"
}
