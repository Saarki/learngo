package main

import (
	"fmt"
	"math"
)

func main() {
	name, age, company := "Saarki", 30, "Gulp water"
	a, b := 250.0, 251.0
	c := math.Min(a, b)
	d := math.Max(a, b)
	fmt.Println("My name is", name, "and I am", age, "and my company is", company)
	fmt.Println(company, "produced", c, "litres of water on Monday and", d, "litres of water on Tuesday")
}
