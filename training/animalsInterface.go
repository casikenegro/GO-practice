package main

import "fmt"

type animal interface {
	move() string
}

type dog struct{}

type fish struct{}

type bird struct{}

func (dog) move() string {
	return "i am dog and walk"
}

func (bird) move() string {
	return "i am bird and fly"
}

func (fish) move() string {
	return "i am fish and glu glu glu"
}

func moveAnimal(a animal) {
	fmt.Println(a.move())
}

func main() {
	doggy := dog{}
	moveAnimal(doggy)
	fishi := fish{}
	moveAnimal(fishi)
	birdy := bird{}
	moveAnimal(birdy)

}
