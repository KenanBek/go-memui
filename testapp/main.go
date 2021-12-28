package main

import (
	"fmt"

	"github.com/KenanBek/go-memui"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	fmt.Println("Hello, World!")

	p1 := Person{"John", 30}
	p2 := Person{"Mary", 25}

	mui := memui.New()
	mui.Register(p1, p2)

	fmt.Println(mui.ListValues("main.Person"))
}
