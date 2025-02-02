package _390

import "fmt"

type Human interface {
	Say()
}

type Person struct {
	Name string
}

func (p Person) Say() {
	fmt.Println("Hello")
}

func (p Person) Action(act string) {
	fmt.Println(act)
}

func C() {
	var mike Human = Person{Name: "Mike"}
	mike.Say()
}
