package main

import (
	"fmt"
	"log"

	"github.com/KenanBek/go-memui"
)

type Department struct {
	Name string
}

type Person struct {
	Name       string
	Age        int
	Department *Department
}

func main() {
	fmt.Println("Hello, World!")

	d1 := Department{Name: "IT"}
	d2 := Department{Name: "HR"}

	p1 := Person{"John", 30, &d1}
	p2 := Person{"Mary", 25, &d2}

	mui := memui.New()

	err := mui.Register(&p1, &p2, &d1, &d2)
	if err != nil {
		fmt.Println(err)
	}

	log.Fatalln(mui.ServerHTTP())
}
