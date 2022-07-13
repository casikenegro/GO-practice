package main

import "fmt"

func main() {
	kids := make(map[string]int) //map of children and their ages
	kids["juan"] = 8
	kids["michael"] = 9
	kids["caroline"] = 7
	fmt.Println(kids)
}
