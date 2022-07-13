package main

import (
	"fmt"
	"io"
	"net/http"
)

type writeWeb struct{}

func (writeWeb) Write(p []byte) (int, error) {
	fmt.Println(string(p))
	return len(p), nil
}

func main() {
	response, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println(err)
	}
	w := writeWeb{}
	io.Copy(w, response.Body)
}
